package builtins

func init() {
	sourceFile = map[string]string{

		"!":                     "not_commands_docgen.go",
		"(":                     "brace-quote_commands_docgen.go",
		"2darray":               "2darray_commands_docgen.go",
		"(murex named pipe)":    "namedpipe_commands_docgen.go",
		"<stdin>":               "stdin_commands_docgen.go",
		"=":                     "equ_commands_docgen.go",
		">>":                    "greater-than-greater-than_commands_docgen.go",
		">":                     "greater-than_commands_docgen.go",
		"@[":                    "range_commands_docgen.go",
		"[[":                    "element_commands_docgen.go",
		"[":                     "index_commands_docgen.go",
		"a":                     "a_commands_docgen.go",
		"alias":                 "alias_commands_docgen.go",
		"alter":                 "alter_commands_docgen.go",
		"and":                   "and_commands_docgen.go",
		"append":                "append_commands_docgen.go",
		"args":                  "args_commands_docgen.go",
		"autocomplete":          "autocomplete_commands_docgen.go",
		"bexists":               "bexists_commands_docgen.go",
		"bg":                    "bg_commands_docgen.go",
		"cast":                  "cast_commands_docgen.go",
		"catch":                 "catch_commands_docgen.go",
		"cd":                    "cd_commands_docgen.go",
		"config":                "config_commands_docgen.go",
		"count":                 "count_commands_docgen.go",
		"cpuarch":               "cpuarch_commands_docgen.go",
		"cpucount":              "cpucount_commands_docgen.go",
		"datetime":              "datetime_commands_docgen.go",
		"debug":                 "debug_commands_docgen.go",
		"die":                   "die_commands_docgen.go",
		"err":                   "err_commands_docgen.go",
		"escape":                "escape_commands_docgen.go",
		"esccli":                "esccli_commands_docgen.go",
		"eschtml":               "eschtml_commands_docgen.go",
		"escurl":                "escurl_commands_docgen.go",
		"event":                 "event_commands_docgen.go",
		"exec":                  "exec_commands_docgen.go",
		"exit":                  "exit_commands_docgen.go",
		"exitnum":               "exitnum_commands_docgen.go",
		"export":                "export_commands_docgen.go",
		"f":                     "f_commands_docgen.go",
		"false":                 "false_commands_docgen.go",
		"fexec":                 "fexec_commands_docgen.go",
		"fg":                    "fg_commands_docgen.go",
		"fid-kill":              "fid-kill_commands_docgen.go",
		"fid-killall":           "fid-killall_commands_docgen.go",
		"fid-list":              "fid-list_commands_docgen.go",
		"for":                   "for_commands_docgen.go",
		"foreach":               "foreach_commands_docgen.go",
		"formap":                "formap_commands_docgen.go",
		"format":                "format_commands_docgen.go",
		"function":              "function_commands_docgen.go",
		"g":                     "g_commands_docgen.go",
		"get-type":              "get-type_commands_docgen.go",
		"get":                   "get_commands_docgen.go",
		"getfile":               "getfile_commands_docgen.go",
		"global":                "global_commands_docgen.go",
		"history":               "history_commands_docgen.go",
		"if":                    "if_commands_docgen.go",
		"ja":                    "ja_commands_docgen.go",
		"jsplit":                "jsplit_commands_docgen.go",
		"left":                  "left_commands_docgen.go",
		"let":                   "let_commands_docgen.go",
		"lockfile":              "lockfile_commands_docgen.go",
		"man-summary":           "man-summary_commands_docgen.go",
		"map":                   "map_commands_docgen.go",
		"match":                 "match_commands_docgen.go",
		"method":                "method_commands_docgen.go",
		"msort":                 "msort_commands_docgen.go",
		"mtac":                  "mtac_commands_docgen.go",
		"murex-docs":            "murex-docs_commands_docgen.go",
		"murex-package":         "murex-package_commands_docgen.go",
		"murex-parser":          "murex-parser_commands_docgen.go",
		"murex-update-exe-list": "murex-update-exe-list_commands_docgen.go",
		"null":                  "devnull_commands_docgen.go",
		"open-image":            "open-image_commands_docgen.go",
		"open":                  "open_commands_docgen.go",
		"openagent":             "openagent_commands_docgen.go",
		"or":                    "or_commands_docgen.go",
		"os":                    "os_commands_docgen.go",
		"out":                   "out_commands_docgen.go",
		"pipe":                  "pipe_commands_docgen.go",
		"post":                  "post_commands_docgen.go",
		"prefix":                "prefix_commands_docgen.go",
		"prepend":               "prepend_commands_docgen.go",
		"pretty":                "pretty_commands_docgen.go",
		"private":               "private_commands_docgen.go",
		"pt":                    "pt_commands_docgen.go",
		"rand":                  "rand_commands_docgen.go",
		"read":                  "read_commands_docgen.go",
		"regexp":                "regexp_commands_docgen.go",
		"right":                 "right_commands_docgen.go",
		"runtime":               "runtime_commands_docgen.go",
		"rx":                    "rx_commands_docgen.go",
		"set":                   "set_commands_docgen.go",
		"source":                "source_commands_docgen.go",
		"struct-keys":           "struct-keys_commands_docgen.go",
		"suffix":                "suffix_commands_docgen.go",
		"summary":               "summary_commands_docgen.go",
		"switch":                "switch_commands_docgen.go",
		"swivel-datatype":       "swivel-datatype_commands_docgen.go",
		"swivel-table":          "swivel-table_commands_docgen.go",
		"ta":                    "ta_commands_docgen.go",
		"tabulate":              "tabulate_commands_docgen.go",
		"test":                  "test_commands_docgen.go",
		"time":                  "time_commands_docgen.go",
		"tmp":                   "tmp_commands_docgen.go",
		"tout":                  "tout_commands_docgen.go",
		"tread":                 "tread_commands_docgen.go",
		"true":                  "true_commands_docgen.go",
		"try":                   "try_commands_docgen.go",
		"trypipe":               "trypipe_commands_docgen.go",
		"version":               "version_commands_docgen.go",
		"while":                 "while_commands_docgen.go",
	}
}
