#
# Autogenerated by Frugal Compiler (2.2.0)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#



from thrift.Thrift import TMessageType
from frugal.middleware import Method
from frugal.transport import TMemoryOutputBuffer




class EventsPublisher(object):
    """
    This docstring gets added to the generated code because it has
    the @ sign. Prefix specifies topic prefix tokens, which can be static or
    variable.
    """

    _DELIMITER = '.'

    def __init__(self, provider, middleware=None):
        """
        Create a new EventsPublisher.

        Args:
            provider: FScopeProvider
            middleware: ServiceMiddleware or list of ServiceMiddleware
        """

        middleware = middleware or []
        if middleware and not isinstance(middleware, list):
            middleware = [middleware]
        middleware += provider.get_middleware()
        self._transport, self._protocol_factory = provider.new_publisher()
        self._methods = {
            'publish_EventCreated': Method(self._publish_EventCreated, middleware),
            'publish_SomeInt': Method(self._publish_SomeInt, middleware),
            'publish_SomeStr': Method(self._publish_SomeStr, middleware),
            'publish_SomeList': Method(self._publish_SomeList, middleware),
        }

    def open(self):
        self._transport.open()

    def close(self):
        self._transport.close()

    def publish_EventCreated(self, ctx, user, req):
        """
        This is a docstring.
        
        Args:
            ctx: FContext
            user: string
            req: Event
        """
        self._methods['publish_EventCreated']([ctx, user, req])

    def _publish_EventCreated(self, ctx, user, req):
        ctx.set_request_header('_topic_user', user)
        op = 'EventCreated'
        prefix = 'foo.{}.'.format(user)
        topic = '{}Events{}{}'.format(prefix, self._DELIMITER, op)
        buffer = TMemoryOutputBuffer(self._transport.get_publish_size_limit())
        oprot = self._protocol_factory.get_protocol(buffer)
        oprot.write_request_headers(ctx)
        oprot.writeMessageBegin(op, TMessageType.CALL, 0)
        req.write(oprot)
        oprot.writeMessageEnd()
        self._transport.publish(topic, buffer.getvalue())


    def publish_SomeInt(self, ctx, user, req):
        """
        Args:
            ctx: FContext
            user: string
            req: i64
        """
        self._methods['publish_SomeInt']([ctx, user, req])

    def _publish_SomeInt(self, ctx, user, req):
        ctx.set_request_header('_topic_user', user)
        op = 'SomeInt'
        prefix = 'foo.{}.'.format(user)
        topic = '{}Events{}{}'.format(prefix, self._DELIMITER, op)
        buffer = TMemoryOutputBuffer(self._transport.get_publish_size_limit())
        oprot = self._protocol_factory.get_protocol(buffer)
        oprot.write_request_headers(ctx)
        oprot.writeMessageBegin(op, TMessageType.CALL, 0)
        oprot.writeI64(req)
        oprot.writeMessageEnd()
        self._transport.publish(topic, buffer.getvalue())


    def publish_SomeStr(self, ctx, user, req):
        """
        Args:
            ctx: FContext
            user: string
            req: string
        """
        self._methods['publish_SomeStr']([ctx, user, req])

    def _publish_SomeStr(self, ctx, user, req):
        ctx.set_request_header('_topic_user', user)
        op = 'SomeStr'
        prefix = 'foo.{}.'.format(user)
        topic = '{}Events{}{}'.format(prefix, self._DELIMITER, op)
        buffer = TMemoryOutputBuffer(self._transport.get_publish_size_limit())
        oprot = self._protocol_factory.get_protocol(buffer)
        oprot.write_request_headers(ctx)
        oprot.writeMessageBegin(op, TMessageType.CALL, 0)
        oprot.writeString(req)
        oprot.writeMessageEnd()
        self._transport.publish(topic, buffer.getvalue())


    def publish_SomeList(self, ctx, user, req):
        """
        Args:
            ctx: FContext
            user: string
            req: list
        """
        self._methods['publish_SomeList']([ctx, user, req])

    def _publish_SomeList(self, ctx, user, req):
        ctx.set_request_header('_topic_user', user)
        op = 'SomeList'
        prefix = 'foo.{}.'.format(user)
        topic = '{}Events{}{}'.format(prefix, self._DELIMITER, op)
        buffer = TMemoryOutputBuffer(self._transport.get_publish_size_limit())
        oprot = self._protocol_factory.get_protocol(buffer)
        oprot.write_request_headers(ctx)
        oprot.writeMessageBegin(op, TMessageType.CALL, 0)
        oprot.writeListBegin(TType.MAP, len(req))
        for elem56 in req:
            oprot.writeMapBegin(TType.I64, TType.STRUCT, len(elem56))
            for elem58, elem57 in elem56.items():
                oprot.writeI64(elem58)
                elem57.write(oprot)
            oprot.writeMapEnd()
        oprot.writeListEnd()
        oprot.writeMessageEnd()
        self._transport.publish(topic, buffer.getvalue())

