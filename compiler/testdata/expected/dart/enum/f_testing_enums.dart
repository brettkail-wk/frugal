// Autogenerated by Frugal Compiler (3.14.4)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

enum testing_enums {
  /// This docstring gets added to the generated code because it
  /// has the @ sign.
  one,
  /// Deprecated: use something else
  two,
  /// This is a docstring comment for a deprecated enum value that has been
  /// spread across two lines.
  /// Deprecated: don't use this; use "something else"
  Three,
}

int serializetesting_enums(testing_enums variant) {
  switch (variant) {
    case testing_enums.one:
      return 45;
    case testing_enums.two:
      return 3;
    case testing_enums.Three:
      return 76;
  }
}

testing_enums deserializetesting_enums(int value) {
  switch (value) {
    case 45:
      return testing_enums.one;
    case 3:
      return testing_enums.two;
    case 76:
      return testing_enums.Three;
    default:
      throw thrift.TProtocolError(thrift.TProtocolErrorType.UNKNOWN, "Invalid value '$value' for enum 'testing_enums'");  }
}
