package docs

func init() {

	Definition["continue"] = "# `continue` - Command Reference\n\n> Terminate process of a block within a caller function\n\n## Description\n\n`continue` will terminate execution of a block (eg `function`, `private`,\n`foreach`, `if`, etc) right up until the caller function. In iteration loops\nlike `foreach` and `formap` this will result in behavior similar to the\n`continue` statement in other programming languages.\n\n## Usage\n\n    continue block-name\n\n## Examples\n\n    %[1..10] -> foreach i {\n        if { $i == 5 } then {\n            out \"continue\"\n            continue foreach\n            out \"skip this code\"\n        }\n        out $i\n    }\n    \nRunning the above code would output:\n\n    » foo\n    1\n    2\n    3\n    4\n    continue\n    6\n    7\n    8\n    9\n    10\n\n## Detail\n\n`continue` cannot escape the bounds of its scope (typically the function it is\nrunning inside). For example, in the following code we are calling `continue\nbar` (which is a different function) inside of the function `foo`:\n\n    function foo {\n        %[1..10] -> foreach i {\n            out $i\n            if { $i == 5 } then {\n                out \"exit running function\"\n                continue bar\n                out \"ended\"\n            }\n        }\n    }\n    \n    function bar {\n        foo\n    }\n    \nRegardless of whether we run `foo` or `bar`, both of those functions will\nraise the following error:\n\n    Error in `continue` (7,17): no block found named `bar` within the scope of `foo`\n\n## See Also\n\n* [`break`](../commands/break.md):\n  Terminate execution of a block within your processes scope\n* [`exit`](../commands/exit.md):\n  Exit murex\n* [`foreach`](../commands/foreach.md):\n  Iterate through an array\n* [`formap`](../commands/formap.md):\n  Iterate through a map or other collection of data\n* [`function`](../commands/function.md):\n  Define a function block\n* [`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`return`](../commands/return.md):\n  Exits current function scope"

}
