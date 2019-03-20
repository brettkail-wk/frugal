// Autogenerated by Frugal Compiler (3.1.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

class HealthCondition {
  /// This docstring gets added to the generated code because it
  /// has the @ sign.
  static const int PASS = 1;
  /// This docstring also gets added to the generated code
  /// because it has the @ sign.
  static const int WARN = 2;
  /// Deprecated: use something else
  @deprecated
  static const int FAIL = 3;
  /// This is a docstring comment for a deprecated enum value that has been
  /// spread across two lines.
  /// Deprecated: don't use this; use "something else"
  @deprecated
  static const int UNKNOWN = 4;

  static final Set<int> VALID_VALUES = new Set.from([
    PASS,
    WARN,
    FAIL,
    UNKNOWN,
  ]);

  static final Map<int, String> VALUES_TO_NAMES = {
    PASS: 'PASS',
    WARN: 'WARN',
    FAIL: 'FAIL',
    UNKNOWN: 'UNKNOWN',
  };
}
