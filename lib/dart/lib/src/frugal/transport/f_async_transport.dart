/*
 * Copyright 2017 Workiva
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

part of frugal.src.frugal;

/// FAsyncTransport is an extension of FTransport that asynchronous frameworks
/// can implement. Implementations need only implement [flush] to send request
/// data and call [handleResponse] when asynchronous responses are received.
abstract class FAsyncTransport extends FTransport {
  final Logger _log = new Logger('FAsyncTransport');
  /// Completed with null if service not available.
  Map<int, Completer<Uint8List>> _handlers = {};

  /// Instantiate an [FAsyncTransport].
  FAsyncTransport({int requestSizeLimit})
      : super(requestSizeLimit: requestSizeLimit);

  /// Flush the payload to the server.
  /// Implementations must be threadsafe.
  /// This method is not called for requests if [flushOp] is overridden.
  Future<Null> flush(Uint8List payload);

  /// Flush a request payload to the server.
  /// By default, this method calls [flush].
  Future<Null> flushOp(int opId, Uint8List payload) {
    flush(payload);
  }

  @override
  Future<Null> oneway(FContext ctx, Uint8List payload) async {
    _preflightRequestCheck(payload);
    await flush(payload).timeout(ctx.timeout, onTimeout: () {
      throw new TTransportError(FrugalTTransportErrorType.TIMED_OUT,
          'request timed out after ${ctx.timeout}');
    });
  }

  @override
  Future<TTransport> request(FContext ctx, Uint8List payload) async {
    _preflightRequestCheck(payload);

    Completer<Uint8List> resultCompleter = new Completer();

    if (_handlers.containsKey(ctx._opId)) {
      throw new StateError("frugal: context already registered");
    }
    _handlers[ctx._opId] = resultCompleter;
    Completer<Uint8List> closedCompleter = new Completer();
    StreamSubscription<Object> closedSub = onClose.listen((_) {
      closedCompleter.completeError(
          new TTransportError(FrugalTTransportErrorType.NOT_OPEN));
    });

    try {
      await flushOp(ctx._opId, payload);
      Future<Uint8List> resultFuture =
          resultCompleter.future.timeout(ctx.timeout);

      // Bail early if the transport is closed
      Uint8List response =
          await Future.any([resultFuture, closedCompleter.future]);
      if (response == null) {
        throw new TTransportError(FrugalTTransportErrorType.SERVICE_NOT_AVAILABLE,
            "service not available");
      }
      return new TMemoryTransport.fromUint8List(response);
    } on TimeoutException catch (_) {
      throw new TTransportError(FrugalTTransportErrorType.TIMED_OUT,
          "request timed out after ${ctx.timeout}");
    } finally {
      _handlers.remove(ctx._opId);

      // don't wait until this is disposed to cancel these
      await closedSub.cancel();
      if (!closedCompleter.isCompleted) {
        closedCompleter.complete();
      }
      if (!resultCompleter.isCompleted) {
        resultCompleter.complete();
      }
    }
  }

  /// Handles a frugal frame response. NOTE: this frame must NOT include the
  /// frame size. Implementations should call this when asynchronous responses
  /// are received from the server.
  void handleResponse(Uint8List frame) {
    var headers = Headers.decodeFromFrame(frame);
    var opId;
    try {
      opId = int.parse(headers[_opidHeader]);
    } catch (e) {
      _log.severe("frugal: invalid protocol frame: op id not a uint64", e);
      return;
    }

    _handleOpResponse(opId, frame);
  }

  /// Handles a frugal response for SERVICE_NOT_AVAILABLE.
  void handleServiceNotAvailable(int opId) {
    _handleOpResponse(opId, null);
  }

  void _handleOpResponse(int opId, Uint8List frame) {
    Completer<Uint8List> handler = _handlers[opId];
    if (handler == null) {
      _log.severe("frugal: no handler found for message, dropping message");
      return;
    }

    if (handler.isCompleted) {
      _log.severe(
          "frugal: handler already called for message, dropping message");
      return;
    }
    handler.complete(frame);
  }
}
