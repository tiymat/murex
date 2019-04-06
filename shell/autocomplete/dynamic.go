package autocomplete

import (
	"bytes"
	"sort"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/readline"
)

type dynamicArgs struct {
	exe    string
	params []string
	float  int
}

func matchDynamic(f *Flags, partial string, args dynamicArgs, defs *map[string]string, tdt *readline.TabDisplayType) (items []string) {
	// Default to building up from Dynamic field. Fall back to DynamicDefs
	dynamic := f.Dynamic
	if f.Dynamic == "" {
		dynamic = f.DynamicDesc
	}
	if dynamic == "" {
		return
	}

	if !types.IsBlock([]byte(dynamic)) {
		lang.ShellProcess.Stderr.Writeln([]byte("Dynamic autocompleter is not a code block"))
		return
	}
	block := []rune(dynamic[1 : len(dynamic)-1])

	fork := lang.ShellFork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.Name = args.exe
	fork.Parameters = parameters.Parameters{Params: args.params}
	fork.Module = ExesFlagsMod[args.exe]
	exitNum, err := fork.Execute(block)

	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Dynamic autocomplete code could not compile: " + err.Error()))
	}
	if exitNum != 0 && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte("Dynamic autocomplete returned a none zero exit number." + utils.NewLineString))
	}

	if f.Dynamic != "" {
		fork.Stdout.ReadArray(func(b []byte) {
			s := string(bytes.TrimSpace(b))
			if len(s) == 0 {
				return
			}
			if strings.HasPrefix(s, partial) {
				items = append(items, s[len(partial):])
			}
		})

	} else {
		if f.ListView {
			*tdt = readline.TabDisplayList
		}

		fork.Stdout.ReadMap(lang.ShellProcess.Config, func(key string, value string, last bool) {
			if strings.HasPrefix(key, partial) {
				items = append(items, key[len(partial):])
				value = strings.Replace(value, "\r", "", -1)
				value = strings.Replace(value, "\n", " ", -1)
				(*defs)[key[len(partial):]+" "] = value
				sort.Strings(items)
			}
		})
	}

	if f.AutoBranch {
		autoBranch(&items)
	}

	return
}
