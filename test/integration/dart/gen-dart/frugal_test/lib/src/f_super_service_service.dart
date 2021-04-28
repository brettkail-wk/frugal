// Autogenerated by Frugal Compiler (3.14.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING



// ignore_for_file: unused_import
// ignore_for_file: unused_field
import 'dart:async';
import 'dart:typed_data' show Uint8List;

import 'package:collection/collection.dart';
import 'package:logging/logging.dart' as logging;
import 'package:thrift/thrift.dart' as thrift;
import 'package:frugal/frugal.dart' as frugal;
import 'package:w_common/disposable.dart' as disposable;

import 'package:frugal_test/frugal_test.dart' as t_frugal_test;


abstract class FSuperService {
  Future testSuperClass(frugal.FContext ctx);
}

FSuperServiceClient fSuperServiceClientFactory(frugal.FServiceProvider provider, {List<frugal.Middleware> middleware}) =>
    FSuperServiceClient(provider, middleware);

class FSuperServiceClient extends disposable.Disposable implements FSuperService {
  static final logging.Logger _frugalLog = logging.Logger('SuperService');
  Map<String, frugal.FMethod> _methods;

  FSuperServiceClient(frugal.FServiceProvider provider, [List<frugal.Middleware> middleware])
      : this._provider = provider {
    _transport = provider.transport;
    _protocolFactory = provider.protocolFactory;
    var combined = middleware ?? [];
    combined.addAll(provider.middleware);
    this._methods = {};
    this._methods['testSuperClass'] = frugal.FMethod(this._testSuperClass, 'SuperService', 'testSuperClass', combined);
  }

  frugal.FServiceProvider _provider;
  frugal.FTransport _transport;
  frugal.FProtocolFactory _protocolFactory;

  @override
  Future<Null> onDispose() async {
    if (_provider is disposable.Disposable && !_provider.isOrWillBeDisposed)  {
      return _provider.dispose();
    }
    return null;
  }

  @override
  Future testSuperClass(frugal.FContext ctx) {
    return this._methods['testSuperClass']([ctx]);
  }

  Future _testSuperClass(frugal.FContext ctx) async {
    final args = testSuperClass_args();
    final message = frugal.prepareMessage(ctx, 'testSuperClass', args, thrift.TMessageType.CALL, _protocolFactory, _transport.requestSizeLimit);
    var response = await _transport.request(ctx, message);

    final result = testSuperClass_result();
    frugal.processReply(ctx, result, response, _protocolFactory);
  }
}

// ignore: camel_case_types
class testSuperClass_args implements thrift.TBase {
  static final thrift.TStruct _STRUCT_DESC = thrift.TStruct('testSuperClass_args');



  @override
  getFieldValue(int fieldID) {
    switch (fieldID) {
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  setFieldValue(int fieldID, Object value) {
    switch (fieldID) {
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  // Returns true if the field corresponding to fieldID is set (has been assigned a value) and false otherwise
  @override
  bool isSet(int fieldID) {
    switch (fieldID) {
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  read(thrift.TProtocol iprot) {
    iprot.readStructBegin();
    for (thrift.TField field = iprot.readFieldBegin();
        field.type != thrift.TType.STOP;
        field = iprot.readFieldBegin()) {
      switch (field.id) {
        default:
          thrift.TProtocolUtil.skip(iprot, field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();

    validate();
  }

  @override
  write(thrift.TProtocol oprot) {
    validate();

    oprot.writeStructBegin(_STRUCT_DESC);
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @override
  String toString() {
    StringBuffer ret = StringBuffer('testSuperClass_args(');

    ret.write(')');

    return ret.toString();
  }

  @override
  bool operator ==(Object o) {
    return o is testSuperClass_args;
  }

  @override
  int get hashCode {
    var value = 17;
    return value;
  }

  testSuperClass_args clone() {
    return testSuperClass_args();
  }

  validate() {
  }
}
// ignore: camel_case_types
class testSuperClass_result implements thrift.TBase {
  static final thrift.TStruct _STRUCT_DESC = thrift.TStruct('testSuperClass_result');



  @override
  getFieldValue(int fieldID) {
    switch (fieldID) {
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  setFieldValue(int fieldID, Object value) {
    switch (fieldID) {
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  // Returns true if the field corresponding to fieldID is set (has been assigned a value) and false otherwise
  @override
  bool isSet(int fieldID) {
    switch (fieldID) {
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  read(thrift.TProtocol iprot) {
    iprot.readStructBegin();
    for (thrift.TField field = iprot.readFieldBegin();
        field.type != thrift.TType.STOP;
        field = iprot.readFieldBegin()) {
      switch (field.id) {
        default:
          thrift.TProtocolUtil.skip(iprot, field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();

    validate();
  }

  @override
  write(thrift.TProtocol oprot) {
    validate();

    oprot.writeStructBegin(_STRUCT_DESC);
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @override
  String toString() {
    StringBuffer ret = StringBuffer('testSuperClass_result(');

    ret.write(')');

    return ret.toString();
  }

  @override
  bool operator ==(Object o) {
    return o is testSuperClass_result;
  }

  @override
  int get hashCode {
    var value = 17;
    return value;
  }

  testSuperClass_result clone() {
    return testSuperClass_result();
  }

  validate() {
  }
}
