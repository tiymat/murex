# _murex_ Language Guide

## Command reference: or

> Returns `true` or `false` depending on whether one code-block out of multiple
ones supplied is successful or unsuccessful.

### Description

Returns a boolean results (`true` or `false`) depending on whether any of the
code-blocks included as parameters are successful or not.

### Usage

    or: { code-block } { code-block } -> <stdout>

    !or: { code-block } { code-block } -> <stdout>

`or` supports as many or as few code-blocks as you wish.

### Examples

    if { or { = 1+1==2 } { = 2+2==5 } } then {
        out: At least one of those equations are correct
    }

### Details

`or` does not set the exit number on failure so it is safe to use inside a `try`
or `trypipe` block.

If `or` is prefixed by a bang (`!or`) then it returns `true` when one or more
code-blocks are unsuccessful (ie the opposite of `or`).

#### Code-Block Testing

* `or` only executes code-blocks up until one of the code-blocks is successful
  then it exits the function and returns `true`.

* `!or` only executes code-blocks while the code-blocks are successful. Once one
  is unsuccessful `!or` exits and returns `true` (ie it `not`s every code-block).

### Synonyms

* !or

### See also

* [`and`](and.md): Returns `true` or `false` depending on whether multiple conditions are met
* [`catch`](catch.md): Handles the exception code raised by `try` or `trypipe`
* [`if`](if.md): Conditional statement to execute different blocks of code depending on the
result of the condition
* `not`
* [`try`](try.md): Handles errors inside a block of code
* [`trypipe`](trypipe.md): Checks state of each function in a pipeline and exits block on error