package docs

func init() {

	Definition["@["] = "# _murex_ Shell Docs\n\n## Command Reference: `@[` (range) \n\n> Outputs a ranged subset of data from STDIN\n\n## Description\n\nThis will read from STDIN and output a subset of data in a defined range.\n\nThe range can be defined as a number of different range types - such as the\ncontent of the array or it's index / row number. You can also omit either\nthe start or the end of the search criteria to cover all items before or\nafter the remaining search criteria.\n\n## Usage\n\n    <stdin> -> @[start..end]flags -> <stdout>\n\n## Examples\n\nRange over all months after March:\n\n    » a: [January..December] -> @[March..]se\n    March\n    April\n    May\n    June\n    July\n    August\n    September\n    October\n    November\n    December\n    \nRange from the 6th to the 10th month (indexes start from zero, `0`):\n\n    » a: [January..December] -> @[5..9]\n    June\n    July\n    August\n    September\n    October\n\n## Flags\n\n* `e`\n    exclude the start and end search criteria from the range\n* `n`\n    array index\n* `r`\n    regexp match\n* `s`\n    exact string match\n\n## Synonyms\n\n* `@[`\n\n\n## See Also\n\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`alter`](../commands/alter.md):\n  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input\n* [commands/`append`](../commands/append.md):\n  Add data to the end of an array\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [commands/`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [commands/len](../commands/len.md):\n  "

}
