/**
 * Autogenerated by Frugal Compiler (1.0.5)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package foo;

import com.workiva.frugal.protocol.*;
import com.workiva.frugal.provider.FScopeProvider;
import com.workiva.frugal.transport.FScopeTransport;
import com.workiva.frugal.transport.FSubscription;
import org.apache.thrift.TException;
import org.apache.thrift.TApplicationException;
import org.apache.thrift.transport.TTransportException;
import org.apache.thrift.protocol.*;

import javax.annotation.Generated;
import java.util.logging.Logger;




@Generated(value = "Autogenerated by Frugal Compiler (1.0.5)", date = "2015-11-24")
public class BlahSubscriber {

	private static final String DELIMITER = ".";
	private static Logger LOGGER = Logger.getLogger(BlahSubscriber.class.getName());

	private final FScopeProvider provider;

	public BlahSubscriber(FScopeProvider provider) {
		this.provider = provider;
	}

	public interface DoStuffHandler {
		void onDoStuff(FContext ctx, Thing req);
	}

	public FSubscription subscribeDoStuff(final DoStuffHandler handler) throws TException {
		final String op = "DoStuff";
		String prefix = "";
		String topic = String.format("%sBlah%s%s", prefix, DELIMITER, op);
		final FScopeProvider.Client client = provider.build();
		FScopeTransport transport = client.getTransport();
		transport.subscribe(topic);

		final FSubscription sub = new FSubscription(topic, transport);
		new Thread(new Runnable() {
			public void run() {
				while (true) {
					try {
						FContext ctx = client.getProtocol().readRequestHeader();
						Thing received = recvDoStuff(op, client.getProtocol());
						handler.onDoStuff(ctx, received);
					} catch (TException e) {
						if (e instanceof TTransportException) {
							TTransportException transportException = (TTransportException) e;
							if (transportException.getType() == TTransportException.END_OF_FILE) {
								return;
							}
						}
						LOGGER.severe("Subscriber recvDoStuff error " + e.getMessage());
						sub.signal(e);
						sub.unsubscribe();
						return;
					}
				}
			}
		}, "subscription").start();

		return sub;
	}

	private Thing recvDoStuff(String op, FProtocol iprot) throws TException {
		TMessage msg = iprot.readMessageBegin();
		if (!msg.name.equals(op)) {
			TProtocolUtil.skip(iprot, TType.STRUCT);
			iprot.readMessageEnd();
			throw new TApplicationException(TApplicationException.UNKNOWN_METHOD);
		}
		Thing req = new Thing();
		req.read(iprot);
		iprot.readMessageEnd();
		return req;
	}


}
