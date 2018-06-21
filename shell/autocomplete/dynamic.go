package autocomplete

import (
	"bytes"
	"sort"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
)

type dynamicArgs struct {
	exe    string
	params []string
	float  int
}

func matchDynamic(f *Flags, partial string, args dynamicArgs) (items []string) {
	//if len(f.Dynamic) == 0 {
	if f.Dynamic == "" {
		return
	}

	if !types.IsBlock([]byte(f.Dynamic)) {
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Dynamic autocompleter is not a code block.")
		return
	}
	block := []rune(f.Dynamic[1 : len(f.Dynamic)-1])

	branch := proc.ShellProcess.BranchFID()
	branch.Process.Scope = branch.Process
	branch.Process.Parent = branch.Process
	branch.Process.Name = args.exe
	branch.Process.Parameters = parameters.Parameters{Params: args.params}
	defer branch.Close()

	stdout := streams.NewStdin()
	exitNum, err := lang.RunBlockNewConfigSpace(block, nil, stdout, nil, branch.Process)

	if err != nil {
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Dynamic autocomplete code could not compile: "+err.Error())
	}
	if exitNum != 0 && debug.Enable {
		ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "Dynamic autocomplete returned a none zero exit number."+utils.NewLineString)
	}

	stdout.ReadArray(func(b []byte) {
		s := string(bytes.TrimSpace(b))
		if len(s) == 0 {
			return
		}
		if strings.HasPrefix(s, partial) {
			items = append(items, s[len(partial):])
		}
	})

	if f.AutoBranch {
		autoBranch(items)
		items = dedup(items)
	}

	return
}

func autoBranch(tree []string) {
	//debug.Json("tree", tree)
	for branch := range tree {

		for i := 0; i < len(tree[branch])-1; i++ {
			if tree[branch][i] == '/' {
				tree[branch] = tree[branch][:i+1]
				break
			}
		}

	}
}

func dedup(items []string) []string {
	m := make(map[string]bool)
	for i := range items {
		m[items[i]] = true
	}

	new := []string{}
	for s := range m {
		new = append(new, s)
	}

	sort.Strings(new)
	return new
}
