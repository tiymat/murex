package docs

func init() {

	Definition["read"] = "# _murex_ Shell Guide\n\n## Command Reference: `read`\n\n> `read` a line of input from the user and store as a variable\n\n### Description\n\nA readline function to allow a line of data inputed from the terminal.\n\n### Usage\n\n    read: \"prompt\" var_name\n    \n    <stdin> -> read: var_name\n\n### Examples\n\n    read: \"What is your name? \" name\n    out: \"Hello $name\"\n    \n    out: What is your name? -> read: name\n    out: \"Hello $name\"\n\n### Detail\n\nIf `read` is called as a method then the prompt string is taken from STDIN.\nOtherwise the prompt string will be the first parameter. However if no prompt\nstring is given then `read` will not write a prompt.\n\nThe last parameter will be the variable name to store the string read by `read`.\nThis variable cannot be prefixed by dollar, `$`, otherwise the shell will write\nthe output of that variable as the last parameter rather than the name of the\nvariable.\n\nThe data type the `read` line will be stored as is `str` (string). If you\nrequire this to be different then please use `tread` (typed read).\n\n### See Also\n\n* commands/[`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* commands/[`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* commands/[`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* commands/[`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* commands/[`err`](../commands/err.md):\n  Print a line to the STDERR\n* commands/[`out`](../commands/out.md):\n  `echo` a string to the STDOUT with a trailing new line character\n* commands/[`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* commands/[`tread`](../commands/tread.md):\n  `read` a line of input from the user and store as a user defined *typed* variable\n* commands/[sprintf](../commands/sprintf.md):\n  "

}
