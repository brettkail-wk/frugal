# JSON output
The `-gen json` option can be used to generate a summary of a Frugal file and
included files so that it can be introspected.

## File format
The top-level object is a map of namespace names (the basename of each file) to
descriptors of the file:
* `s`: an object containing a map of service name to descriptor
* `c`: an object containing a map of scope names to descriptor
* `t`: an object containing a map of typedef, enum, struct, union, and exception to type descriptor

## Service
* `m`: an object containing a map of method name to method

### Service method
* `p`: an object containing a map of parameter field id (1-based) to field descriptor
* `r`: an object containing a map of result field id (0 for success, 1 for exceptions) to field descriptor

## Field
* `n`: the field name, or absent for the success result of a service method
* `t`: the type descriptor

## Scope
* `p`: the prefix
* `o`: an object containing a map of operation name to type descriptor

## Type
* `a`: an object containing a map of annotation name to value
* One of:
    * `b`: a base type name (for example, `i32`)
    * `k` and `v`: for a `map` type, the key and value type descriptors, respectively
    * `k`: for a `set` type, the element type descriptor
    * `v`: for a `list` type, the element type descriptor
    * `n`: the name of a named type, which will contain a `.` if the type is in another namespace
    * `e`: for an `enum` type, an object containing a map of value to value names
    * `s`: for a `struct` type, an object containing a map of field id to field
    * `u`: for a `union` type, an object containing a map of field id to field
