/**
 * Autogenerated by Frugal Compiler (3.5.0)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */
package variety.java;

import org.apache.thrift.scheme.IScheme;
import org.apache.thrift.scheme.SchemeFactory;
import org.apache.thrift.scheme.StandardScheme;

import org.apache.thrift.scheme.TupleScheme;
import org.apache.thrift.protocol.TTupleProtocol;
import org.apache.thrift.protocol.TProtocolException;
import org.apache.thrift.EncodingUtils;
import org.apache.thrift.TException;
import org.apache.thrift.async.AsyncMethodCallback;
import org.apache.thrift.server.AbstractNonblockingServer.*;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.EnumMap;
import java.util.Set;
import java.util.HashSet;
import java.util.EnumSet;
import java.util.Collections;
import java.util.BitSet;
import java.nio.ByteBuffer;
import java.util.Arrays;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class varietyConstants {
	public static final int redef_const = actual_base.java.baseConstants.const_i32_from_base;

	public static final actual_base.java.thing const_thing = new actual_base.java.thing();
	static {
		const_thing.setAn_id(1);
		const_thing.setA_string("some string");
	}

	public static final long DEFAULT_ID = -1L;

	public static final long other_default = varietyConstants.DEFAULT_ID;

	public static final byte thirtyfour = (byte)34;

	public static final java.util.Map<String, String> MAPCONSTANT = new HashMap<String,String>();
	static {
		MAPCONSTANT.put("hello", "world");
		MAPCONSTANT.put("goodnight", "moon");
	}

	public static final java.util.Set<String> SETCONSTANT = new HashSet<String>();
	static {
		SETCONSTANT.add("hello");
		SETCONSTANT.add("world");
	}

	public static final Event ConstEvent1 = new Event();
	static {
		ConstEvent1.setID(-2L);
		ConstEvent1.setMessage("first one");
	}

	public static final Event ConstEvent2 = new Event();
	static {
		ConstEvent2.setID(-7L);
		ConstEvent2.setMessage("second one");
	}

	public static final java.util.List<Integer> NumsList = new ArrayList<Integer>();
	static {
		NumsList.add(2);
		NumsList.add(4);
		NumsList.add(7);
		NumsList.add(1);
	}

	public static final java.util.Set<Integer> NumsSet = new HashSet<Integer>();
	static {
		NumsSet.add(1);
		NumsSet.add(3);
		NumsSet.add(8);
		NumsSet.add(0);
	}

	public static final java.util.Map<String, Event> MAPCONSTANT2 = new HashMap<String,Event>();
	static {
		Event elem0 = new Event();
		elem0.setID(-2L);
		elem0.setMessage("first here");
		MAPCONSTANT2.put("hello", elem0);
	}

	public static final java.nio.ByteBuffer bin_const = java.nio.ByteBuffer.wrap("hello".getBytes());

	public static final boolean true_constant = true;

	public static final boolean false_constant = false;

	public static final HealthCondition const_hc = HealthCondition.WARN;

	public static final String evil_string = "thin'g\" \"";

	public static final String evil_string2 = "th'ing\"ad\"f";

	public static final TestLowercase const_lower = new TestLowercase();
	static {
		const_lower.setLowercaseInt(2);
	}

}
