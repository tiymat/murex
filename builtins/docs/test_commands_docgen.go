package docs

func init() {

	Definition["test"] = "# _murex_ Shell Docs\n\n## Command Reference: `test`\n\n> _murex_'s test framework - define tests, run tests and debug shell scripts\n\n## Description\n\n`test` is used to define tests, run tests and debug _murex_ shell scripts.\n\n## Usage\n\nDefine an inlined test\n\n    test: define test-name { json-properties }\n    \nDefine a state report\n\n    test: state name { code block }\n    \nDefine a unit test\n\n    test: unit function|private|open|event test-name { json-properties }\n    \nEnable or disable boolean test states (more options available in `config`)\n\n    test: config [ enable|!enable ] [ verbose|!verbose ] [ auto-report|!auto-report ]\n    \nDisable test mode\n\n    !test\n    \nExecute a function with testing enabled\n\n    test: run { code-block }\n    \nExecute unit test(s)\n\n    test: run package/module/test-name|*\n    \nWrite report\n\n    test: report\n\n## Examples\n\nInlined test\n\n    function: hello-world {\n        test: define example {\n            \"StdoutRegex\": (^Hello World$)\n        }\n    \n        out: <test_example> \"Hello Earth\"\n    }\n    \n    test: run { hello-world }\n    \nUnit test\n\n    test: unit function aliases {\n        \"PreBlock\": ({\n            alias ALIAS_UNIT_TEST=example param1 param2 param3\n        }),\n        \"StdoutRegex\": \"([- _0-9a-zA-Z]+ => .*?\\n)+\",\n        \"StdoutType\": \"str\",\n        \"PostBlock\": ({\n            !alias ALIAS_UNIT_TEST\n        })\n    }\n    \n    function: aliases {\n        # Output the aliases in human readable format\n        runtime: --aliases -> formap: name alias {\n            $name -> sprintf: \"%10s => ${esccli @alias}\\n\"\n        } -> cast: str\n    }\n    \n    test: run aliases\n\n## Detail\n\n### Report\n\n`test: report` is only needed if `config: test auto-report` is set false.\nHowever `test: run` automatically enables **auto-report**.\n\nWhen the report is generated, be it automatically or manually triggered, it\nflushes the table of pending reports.\n\n## Synonyms\n\n* `test`\n* `!test`\n\n\n## See Also\n\n* [commands/`<>` / `read-named-pipe`](../commands/namedpipe.md):\n  Reads from a _murex_ named pipe\n* [commands/`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [parser/namedpipe](../parser/namedpipe.md):\n  "

}
