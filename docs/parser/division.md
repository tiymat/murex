# `/` Division Operator (expr)

> Divides one numeric value from another

## Description

The Division Operator divides the left hand number by the right hand number in
an expression.



## Examples

#### Expression

```
» 3/2
1.5
```

#### Statement

```
out (3/2)
» 1.5
```

## Detail

### Type Safety

Because shells are historically untyped, you cannot always guarantee that a
numeric-looking value isn't a string. To solve this problem, by default Murex
assumes anything that looks like a number is a number when performing addition.

```
» str = "2"
» int = 3
» $str + $int
1
```

For occasions when type safety is more important than the convenience of silent
data casting, you can disable the above behaviour via `config`:

```
» config set proc strict-types false
» $str + $int
Error in `expr` (0,1): cannot Add with string types
                    > Expression: $str + $int
                    >           : ^
                    > Character : 1
                    > Symbol    : Scalar
                    > Value     : '$str'
```

## See Also

* [`*` Multiplication Operator (expr)](../parser/multiplication.md):
  Multiplies one numeric value with another
* [`+` Addition Operator (expr)](../parser/addition.md):
  Adds two numeric values together
* [`-` Subtraction Operator (expr)](../parser/subtraction.md):
  Subtracts one numeric value from another
* [`/=` Divide By Operator (expr)](../parser/divide-by.md):
  Divides a variable by the right hand value
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`expr`](../commands/expr.md):
  Expressions: mathematical, string comparisons, logical operators
* [`float` (floating point number)](../types/float.md):
  Floating point number (primitive)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)

<hr/>

This document was generated from [gen/expr/division_op_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/expr/division_op_doc.yaml).