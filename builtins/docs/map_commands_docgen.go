package docs

func init() {

	Definition["map"] = "# _murex_ Shell Docs\n\n## Command Reference: `map` \n\n> Creates a map from two data sources\n\n## Description\n\nThis takes two parameters - which are code blocks - and combines them to output a key/value map in JSON.\n\nThe first block is the key and the second is the value.\n\n## Usage\n\n    map { code-block } { code-block } -> <stdout>\n\n## Examples\n\n    » map { tout: json ([\"key 1\", \"key 2\", \"key 3\"]) } { tout: json ([\"value 1\", \"value 2\", \"value 3\"]) } \n    {\n        \"key 1\": \"value 1\",\n        \"key 2\": \"value 2\",\n        \"key 3\": \"value 3\"\n    }\n\n## See Also\n\n* [commands/`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`alter`](../commands/alter.md):\n  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input\n* [commands/`append`](../commands/append.md):\n  Add data to the end of an array\n* [commands/`count`](../commands/count.md):\n  Count items in a map, list or array\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [commands/`prepend` ](../commands/prepend.md):\n  Add data to the start of an array"

}
