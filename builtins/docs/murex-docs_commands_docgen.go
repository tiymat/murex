package docs

func init() {

	Definition["murex-docs"] = "# _murex_ Shell Guide\n\n## Command Reference: `murex-docs`\n\n> Displays the man pages for _murex_ builtins\n\n### Description\n\nDisplays the man pages for _murex_ builtins.\n\n### Usage\n\n    murex-docs: [ flags ] command -> <stdout>\n\n### Examples\n\n    # Output this man page\n    murex-docs: murex-docs\n\n### Flags\n\n* `--summary`\n    Returns an abridged description of the command rather than the entire help page.\n\n### Detail\n\nThese man pages are compiled into the _murex_ executable.\n\n### See Also\n\n* commands/[`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* commands/[`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* commands/[`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* commands/[`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* commands/[`err`](../commands/err.md):\n  Print a line to the STDERR\n* commands/[`out`](../commands/out.md):\n  `echo` a string to the STDOUT with a trailing new line character\n* commands/[`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* commands/[`tread`](../commands/tread.md):\n  `read` a line of input from the user and store as a user defined *typed* variable\n* commands/[sprintf](../commands/sprintf.md):\n  "

}
