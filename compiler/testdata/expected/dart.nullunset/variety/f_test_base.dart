// Autogenerated by Frugal Compiler (3.14.4)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

// ignore_for_file: unused_import
// ignore_for_file: unused_field
import 'dart:typed_data' show Uint8List;

import 'package:collection/collection.dart';
import 'package:thrift/thrift.dart' as thrift;
import 'package:variety/variety.dart' as t_variety;
import 'package:actual_base_dart/actual_base_dart.dart' as t_actual_base_dart;
import 'package:intermediate_include/intermediate_include.dart' as t_intermediate_include;
import 'package:validStructs/validStructs.dart' as t_validStructs;
import 'package:ValidTypes/ValidTypes.dart' as t_ValidTypes;
import 'package:subdir_include_ns/subdir_include_ns.dart' as t_subdir_include_ns;

class TestBase implements thrift.TBase {
  static final thrift.TStruct _STRUCT_DESC = thrift.TStruct('TestBase');
  static final thrift.TField _BASE_STRUCT_FIELD_DESC = thrift.TField('base_struct', thrift.TType.STRUCT, 1);

  t_actual_base_dart.thing base_struct;
  static const int BASE_STRUCT = 1;


  bool isSetBase_struct() => this.base_struct != null;

  unsetBase_struct() {
    this.base_struct = null;
  }

  @override
  getFieldValue(int fieldID) {
    switch (fieldID) {
      case BASE_STRUCT:
        return this.base_struct;
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  setFieldValue(int fieldID, Object value) {
    switch (fieldID) {
      case BASE_STRUCT:
        if (value == null) {
          unsetBase_struct();
        } else {
          this.base_struct = value as t_actual_base_dart.thing;
        }
        break;

      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  // Returns true if the field corresponding to fieldID is set (has been assigned a value) and false otherwise
  @override
  bool isSet(int fieldID) {
    switch (fieldID) {
      case BASE_STRUCT:
        return isSetBase_struct();
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
        case BASE_STRUCT:
          if (field.type == thrift.TType.STRUCT) {
            this.base_struct = t_actual_base_dart.thing();
            base_struct.read(iprot);
          } else {
            thrift.TProtocolUtil.skip(iprot, field.type);
          }
          break;
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
    if (isSetBase_struct()) {
      oprot.writeFieldBegin(_BASE_STRUCT_FIELD_DESC);
      this.base_struct.write(oprot);
      oprot.writeFieldEnd();
    }
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @override
  String toString() {
    StringBuffer ret = StringBuffer('TestBase(');

    ret.write('base_struct:');
    if (this.base_struct == null) {
      ret.write('null');
    } else {
      ret.write(this.base_struct);
    }

    ret.write(')');

    return ret.toString();
  }

  @override
  bool operator ==(Object o) {
    if (o is TestBase) {
      return this.base_struct == o.base_struct;
    }
    return false;
  }

  @override
  int get hashCode {
    var value = 17;
    value = (value * 31) ^ this.base_struct.hashCode;
    return value;
  }

  TestBase clone({
    t_actual_base_dart.thing base_struct,
  }) {
    return TestBase()
      ..base_struct = base_struct ?? this.base_struct;
  }

  validate() {
  }
}