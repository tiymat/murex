package docs

func init() {

	Definition["pipe"] = "# _murex_ Shell Docs\n\n## Command Reference: `pipe`\n\n> Manage _murex_ named pipes\n\n## Description\n\n`pipe` creates and destroys _murex_ named pipes.\n\n## Usage\n\nCreate pipe\n\n    pipe: name [ pipe-type ]\n    \nDestroy pipe\n\n    !pipe: name\n\n## Examples\n\nCreate a standard pipe:\n\n    pipe: example\n    \nDelete a pipe:\n\n    !pipe: example\n    \nCreate a TCP pipe (deleting a pipe is the same regardless of the type of pipe):\n\n    pipe example --tcp-dial google.com:80\n    bg { <example> }\n    out: \"GET /\" -> <example>\n\n## Detail\n\n### What are _murex_ named pipes?\n\nIn POSIX, there is a concept of STDIN, STDOUT and STDERR, these are FIFO files\nwhile are \"piped\" from one executable to another. ie STDOUT for application 'A'\nwould be the same file as STDIN for application 'B' when A is piped to B:\n`A | B`. _murex_ adds a another layer around this to enable support for passing\ndata types and builtins which are agnostic to the data serialization format\ntraversing the pipeline. While this does add overhead the advantage is this new\nwrapper can be used as a primitive for channelling any data from one point to\nanother.\n\n_murex_ named pipes are where these pipes are created in a global store,\ndecoupled from any executing functions, named and can then be used to pass\ndata along asynchronously.\n\nFor example\n\n    pipe: example\n    \n    bg {\n        <example> -> match: Hello\n    }\n    \n    out: \"foobar\"        -> <example>\n    out: \"Hello, world!\" -> <example>\n    out: \"foobar\"        -> <example>\n    \n    !pipe: example\n    \nThis returns `Hello, world!` because `out` is writing to the **example** named\npipe and `match` is also reading from it in the background (`bg`).\n\nNamed pipes can also be inlined into the command parameters with `<>` tags\n\n    pipe: example\n    \n    bg {\n        <example> -> match: Hello\n    }\n    \n    out: <example> \"foobar\"\n    out: <example> \"Hello, world!\"\n    out: <example> \"foobar\"\n    \n    !pipe: example\n    \n> Please note this is also how `test` works.\n\n_murex_ named pipes can also represent network sockets, files on a disk or any\nother read and/or write endpoint. Custom builtins can also be written in Golang\nto support different abstractions so your _murex_ code can work with those read\nor write endpoints transparently.\n\nTo see the different supported types run\n\n    runtime --pipes\n    \n### Namespaces and usage in modules and packages\n\nPipes created via `pipe` are created in the global namespace. This allows pipes\nto be used across different functions easily however it does pose a risk with\nname clashes where _murex_ named pipes are used heavily. Thus is it recommended\nthat pipes created in modules should be prefixed with the name of its package.\n\n## Synonyms\n\n* `pipe`\n* `!pipe`\n\n\n## See Also\n\n* [user-guide/Pipeline](../user-guide/pipeline.md):\n  Overview of what a \"pipeline\" is\n* [commands/`<>` / `read-named-pipe`](../commands/namedpipe.md):\n  Reads from a _murex_ named pipe\n* [commands/`<stdin>` ](../commands/stdin.md):\n  Read the STDIN belonging to the parent code block\n* [commands/`bg`](../commands/bg.md):\n  Run processes in the background\n* [commands/`match`](../commands/match.md):\n  Match an exact value in an array\n* [commands/`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [commands/`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [commands/`test`](../commands/test.md):\n  _murex_'s test framework - define tests, run tests and debug shell scripts\n* [parser/namedpipe](../parser/namedpipe.md):\n  "

}
