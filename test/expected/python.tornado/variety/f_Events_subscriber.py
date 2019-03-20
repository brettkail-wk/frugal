#
# Autogenerated by Frugal Compiler (3.1.0)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#



import sys
import traceback

from thrift.Thrift import TApplicationException
from thrift.Thrift import TMessageType
from thrift.Thrift import TType
from tornado import gen
from frugal.exceptions import TApplicationExceptionType
from frugal.middleware import Method
from frugal.subscription import FSubscription
from frugal.transport import TMemoryOutputBuffer

from .ttypes import *




class EventsSubscriber(object):
    """
    This docstring gets added to the generated code because it has
    the @ sign. Prefix specifies topic prefix tokens, which can be static or
    variable.
    """

    _DELIMITER = '.'

    def __init__(self, provider, middleware=None):
        """
        Create a new EventsSubscriber.

        Args:
            provider: FScopeProvider
            middleware: ServiceMiddleware or list of ServiceMiddleware
        """

        middleware = middleware or []
        if middleware and not isinstance(middleware, list):
            middleware = [middleware]
        middleware += provider.get_middleware()
        self._middleware = middleware
        self._provider = provider

    @gen.coroutine
    def subscribe_EventCreated(self, user, EventCreated_handler):
        """
        This is a docstring.
        
        Args:
            user: string
            EventCreated_handler: function which takes FContext and Event
        """

        op = 'EventCreated'
        prefix = 'foo.{}.'.format(user)
        topic = '{}Events{}{}'.format(prefix, self._DELIMITER, op)

        transport, protocol_factory = self._provider.new_subscriber()
        yield transport.subscribe(topic, self._recv_EventCreated(protocol_factory, op, EventCreated_handler))
        raise gen.Return(FSubscription(topic, transport))

    def _recv_EventCreated(self, protocol_factory, op, handler):
        method = Method(handler, self._middleware)

        def callback(transport):
            iprot = protocol_factory.get_protocol(transport)
            ctx = iprot.read_request_headers()
            mname, _, _ = iprot.readMessageBegin()
            if mname != op:
                iprot.skip(TType.STRUCT)
                iprot.readMessageEnd()
                raise TApplicationException(TApplicationExceptionType.UNKNOWN_METHOD)
            req = Event()
            req.read(iprot)
            iprot.readMessageEnd()
            try:
                method([ctx, req])
            except:
                traceback.print_exc()
                sys.exit(1)

        return callback



    @gen.coroutine
    def subscribe_SomeInt(self, user, SomeInt_handler):
        """
        Args:
            user: string
            SomeInt_handler: function which takes FContext and i64
        """

        op = 'SomeInt'
        prefix = 'foo.{}.'.format(user)
        topic = '{}Events{}{}'.format(prefix, self._DELIMITER, op)

        transport, protocol_factory = self._provider.new_subscriber()
        yield transport.subscribe(topic, self._recv_SomeInt(protocol_factory, op, SomeInt_handler))
        raise gen.Return(FSubscription(topic, transport))

    def _recv_SomeInt(self, protocol_factory, op, handler):
        method = Method(handler, self._middleware)

        def callback(transport):
            iprot = protocol_factory.get_protocol(transport)
            ctx = iprot.read_request_headers()
            mname, _, _ = iprot.readMessageBegin()
            if mname != op:
                iprot.skip(TType.STRUCT)
                iprot.readMessageEnd()
                raise TApplicationException(TApplicationExceptionType.UNKNOWN_METHOD)
            req = iprot.readI64()
            iprot.readMessageEnd()
            try:
                method([ctx, req])
            except:
                traceback.print_exc()
                sys.exit(1)

        return callback



    @gen.coroutine
    def subscribe_SomeStr(self, user, SomeStr_handler):
        """
        Args:
            user: string
            SomeStr_handler: function which takes FContext and string
        """

        op = 'SomeStr'
        prefix = 'foo.{}.'.format(user)
        topic = '{}Events{}{}'.format(prefix, self._DELIMITER, op)

        transport, protocol_factory = self._provider.new_subscriber()
        yield transport.subscribe(topic, self._recv_SomeStr(protocol_factory, op, SomeStr_handler))
        raise gen.Return(FSubscription(topic, transport))

    def _recv_SomeStr(self, protocol_factory, op, handler):
        method = Method(handler, self._middleware)

        def callback(transport):
            iprot = protocol_factory.get_protocol(transport)
            ctx = iprot.read_request_headers()
            mname, _, _ = iprot.readMessageBegin()
            if mname != op:
                iprot.skip(TType.STRUCT)
                iprot.readMessageEnd()
                raise TApplicationException(TApplicationExceptionType.UNKNOWN_METHOD)
            req = iprot.readString()
            iprot.readMessageEnd()
            try:
                method([ctx, req])
            except:
                traceback.print_exc()
                sys.exit(1)

        return callback



    @gen.coroutine
    def subscribe_SomeList(self, user, SomeList_handler):
        """
        Args:
            user: string
            SomeList_handler: function which takes FContext and list<map<id,Event>>
        """

        op = 'SomeList'
        prefix = 'foo.{}.'.format(user)
        topic = '{}Events{}{}'.format(prefix, self._DELIMITER, op)

        transport, protocol_factory = self._provider.new_subscriber()
        yield transport.subscribe(topic, self._recv_SomeList(protocol_factory, op, SomeList_handler))
        raise gen.Return(FSubscription(topic, transport))

    def _recv_SomeList(self, protocol_factory, op, handler):
        method = Method(handler, self._middleware)

        def callback(transport):
            iprot = protocol_factory.get_protocol(transport)
            ctx = iprot.read_request_headers()
            mname, _, _ = iprot.readMessageBegin()
            if mname != op:
                iprot.skip(TType.STRUCT)
                iprot.readMessageEnd()
                raise TApplicationException(TApplicationExceptionType.UNKNOWN_METHOD)
            req = []
            (_, elem62) = iprot.readListBegin()
            for _ in range(elem62):
                elem63 = {}
                (_, _, elem64) = iprot.readMapBegin()
                for _ in range(elem64):
                    elem66 = iprot.readI64()
                    elem65 = Event()
                    elem65.read(iprot)
                    elem63[elem66] = elem65
                iprot.readMapEnd()
                req.append(elem63)
            iprot.readListEnd()
            iprot.readMessageEnd()
            try:
                method([ctx, req])
            except:
                traceback.print_exc()
                sys.exit(1)

        return callback




