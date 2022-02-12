package docs

func init() {

	Definition["let"] = "# _murex_ Shell Docs\n\n## Command Reference: `let`\n\n> Evaluate a mathematical function and assign to variable\n\n## Description\n\n`let` evaluates a mathematical function and then assigns it to a locally\nscoped variable (like `set`)\n\n## Usage\n\n    let var_name=evaluation\n    \n    let var_name++\n    \n    let var_name--\n\n## Examples\n\n    » let: age=18\n    » $age\n    18\n    \n    » let: age++\n    » $age\n    19\n    \n    » let: under18=age<18\n    » $under18\n    false\n    \n    » let: under21 = age < 21\n    » $under21\n    true\n\n## Detail\n\n### Other Operators\n\n`let` also supports the following operators (substitute **VAR** with your\nvariable name, and **NUM** with a number):\n\n* `VAR--`, subtract 1 from VAR\n* `VAR++`, add 1 to VAR\n* `VAR -= NUM`, subtract NUM from VAR\n* `VAR += NUM`, add NUM to VAR\n* `VAR /= NUM`, divide VAR by NUM\n* `VAR *= NUM`, multiply VAR by NUM\n\neg\n\n    » let: i=0\n    » let: i++\n    » $i\n    1\n    \n    » let: i+=8\n    » $i\n    9\n    \n    » let: i/=3\n    3\n    \nPlease note these operators are not supported by `=`.\n\n### Variables\n\nThere are two ways you can use variables with the math functions. Either by\nstring interpolation like you would normally with any other function, or\ndirectly by name.\n\nString interpolation:\n\n    » set abc=123\n    » = $abc==123\n    true\n    \nDirectly by name:\n\n    » set abc=123\n    » = abc==123\n    false\n    \nTo understand the difference between the two, you must first understand how\nstring interpolation works; which is where the parser tokenised the parameters\nlike so\n\n    command line: = $abc==123\n    token 1: command (name: \"=\")\n    token 2: parameter 1, string (content: \"\")\n    token 3: parameter 1, variable (name: \"abc\")\n    token 4: parameter 1, string (content: \"==123\")\n    \nThen when the command line gets executed, the parameters are compiled on demand\nsimilarly to this crude pseudo-code\n\n    command: \"=\"\n    parameters 1: concatenate(\"\", GetValue(abc), \"==123\")\n    output: \"=\" \"123==123\"\n    \nThus the actual command getting run is literally `123==123` due to the variable\nbeing replace **before** the command executes.\n\nWhereas when you call the variable by name it's up to `=` or `let` to do the\nvariable substitution.\n\n    command line: = abc==123\n    token 1: command (name: \"=\")\n    token 2: parameter 1, string (content: \"abc==123\")\n    \n    command: \"=\"\n    parameters 1: concatenate(\"abc==123\")\n    output: \"=\" \"abc==123\"\n    \nThe main advantage (or disadvantage, depending on your perspective) of using\nvariables this way is that their data-type is preserved.\n\n    » set str abc=123\n    » = abc==123\n    false\n    \n    » set int abc=123\n    » = abc==123\n    true\n    \nUnfortunately is one of the biggest areas in _murex_ where you'd need to be\ncareful. The simple addition or omission of the dollar prefix, `$`, can change\nthe behavior of `=` and `let`.\n\n### Strings\n\nBecause the usual _murex_ tools for encapsulating a string (`\"`, `'` and `()`)\nare interpreted by the shell language parser, it means we need a new token for\nhandling strings inside `=` and `let`. This is where backtick comes to our\nrescue.\n\n    » set str abc=123\n    » = abc==`123`\n    true\n    \nPlease be mindful that if you use string interpolation then you will need to\ninstruct `=` and `let` that your field is a string\n\n    » set str abc=123\n    » = `$abc`==`123`\n    true\n    \n### Best practice recommendation\n\nAs you can see from the sections above, string interpolation offers us some\nconveniences when comparing variables of differing data-types, such as a `str`\ntype with a number (eg `num` or `int`). However it makes for less readable code\nwhen just comparing strings. Thus the recommendation is to avoid using string\ninterpolation except only where it really makes sense (ie use it sparingly).\n\n### Non-boolean logic\n\nThus far the examples given have been focused on comparisons however `=` and\n`let` supports all the usual arithmetic operators:\n\n    » = 10+10\n    20\n    \n    » = 10/10\n    1\n    \n    » = (4 * (3 + 2))\n    20\n    \n    » = `foo`+`bar`\n    foobar\n    \n### Read more\n\n_murex_ uses the [govaluate package](https://github.com/Knetic/govaluate). More information can be found in it's manual.\n\n### Scoping\n\nVariable scoping is simplified to three layers:\n\n1. Local variables (`set`, `!set`, `let`)\n2. Global variables (`global`, `!global`)\n3. Environmental variables (`export`, `!export`, `unset`)\n\nVariables are looked up in that order of too. For example a the following\ncode where `set` overrides both the global and environmental variable:\n\n    » set:    foobar=1\n    » global: foobar=2\n    » export: foobar=3\n    » out: $foobar\n    1\n    \n#### Local variables\n\nThese are defined via `set` and `let`. They're variables that are persistent\nacross any blocks within a function. Functions will typically be blocks\nencapsulated like so:\n\n    function example {\n        # variables scoped inside here\n    }\n    \n...or...\n\n    private example {\n        # variables scoped inside here\n    }\n    \n    \n...however dynamic autocompletes, events, unit tests and any blocks defined in\n`config` will also be triggered as functions.\n\nCode running inside any control flow or error handing structures will be\ntreated as part of the same part of the same scope as the parent function:\n\n    » function example {\n    »     try {\n    »         # set 'foobar' inside a `try` block\n    »         set: foobar=example\n    »     }\n    »     # 'foobar' exists outside of `try` because it is scoped to `function`\n    »     out: $foobar\n    » }\n    example\n    \nWhere this behavior might catch you out is with iteration blocks which create\nvariables, eg `for`, `foreach` and `formap`. Any variables created inside them\nare still shared with any code outside of those structures but still inside the\nfunction block.\n\nAny local variables are only available to that function. If a variable is\ndefined in a parent function that goes on to call child functions, then those\nlocal variables are not inherited but the child functions:\n\n    » function parent {\n    »     # set a local variable\n    »     set: foobar=example\n    »     child\n    » }\n    » \n    » function child {\n    »     # returns the `global` value, \"not set\", because the local `set` isn't inherited\n    »     out: $foobar\n    » }\n    » \n    » global: $foobar=\"not set\"\n    » parent\n    not set\n    \nIt's also worth remembering that any variable defined using `set` in the shells\nFID (ie in the interactive shell) is localised to structures running in the\ninteractive, REPL, shell and are not inherited by any called functions.\n\n#### Global variables\n\nWhere `global` differs from `set` is that the variables defined with `global`\nwill be scoped at the global shell level (please note this is not the same as\nenvironmental variables!) so will cascade down through all scoped code-blocks\nincluding those running in other threads.\n\n#### Environmental variables\n\nExported variables (defined via `export`) are system environmental variables.\nInside _murex_ environmental variables behave much like `global` variables\nhowever their real purpose is passing data to external processes. For example\n`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):\n\n    » export foo=bar\n    » env -> grep foo\n    foo=bar\n    \n### Function Names\n\nAs a security feature function names cannot include variables. This is done to\nreduce the risk of code executing by mistake due to executables being hidden\nbehind variable names.\n\nInstead _murex_ will assume you want the output of the variable printed:\n\n    » out \"Hello, world!\" -> set hw\n    » $hw\n    Hello, world!\n    \nOn the rare occasions you want to force variables to be expanded inside a\nfunction name, then call that function via `exec`:\n\n    » set cmd=grep\n    » ls -> exec: $cmd main.go\n    main.go\n    \nThis only works for external executables. There is currently no way to call\naliases, functions nor builtins from a variable and even the above `exec` trick\nis considered bad form because it reduces the readability of your shell scripts.\n\n### Usage Inside Quotation Marks\n\nLike with Bash, Perl and PHP: _murex_ will expand the variable when it is used\ninside a double quotes but will escape the variable name when used inside single\nquotes:\n\n    » out \"$foo\"\n    bar\n    \n    » out '$foo'\n    $foo\n    \n    » out ($foo)\n    bar\n\n## See Also\n\n* [user-guide/Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by _murex_\n* [user-guide/Variable and Config Scoping](../user-guide/scoping.md):\n  How scoping works within _murex_\n* [commands/`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [commands/`=` (arithmetic evaluation)](../commands/equ.md):\n  Evaluate a mathematical function\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`export`](../commands/export.md):\n  Define an environmental variable and set it's value\n* [commands/`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [commands/`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value"

}
