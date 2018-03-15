package shell

import (
	"strings"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/history"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/murex/utils/readline"
)

var manDescript map[string]string = make(map[string]string)

func tabCompletion(line []rune, pos int) (prefix string, items []string) {
	if len(line) > pos-1 {
		line = line[:pos]
	}

	pt, _ := parse(line)

	switch {
	case pt.Variable != "":
		var s string
		if pt.VarLoc < len(line) {
			s = strings.TrimSpace(string(line[pt.VarLoc:]))
		}
		s = pt.Variable + s
		//retPos = len(s)
		prefix = s

		items = autocomplete.MatchVars(s)

	case pt.ExpectFunc:
		var s string
		if pt.Loc < len(line) {
			s = strings.TrimSpace(string(line[pt.Loc:]))
		}
		//retPos = len(s)
		prefix = s
		items = autocomplete.MatchFunction(s)

	default:
		var s string
		if len(pt.Parameters) > 0 {
			s = pt.Parameters[len(pt.Parameters)-1]
		}
		//retPos = len(s)
		prefix = s

		autocomplete.InitExeFlags(pt.FuncName)

		pIndex := 0
		items = autocomplete.MatchFlags(autocomplete.ExesFlags[pt.FuncName], s, pt.FuncName, pt.Parameters, &pIndex)
	}

	v, err := proc.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = 4
	}

	limitSuggestions := v.(int)
	if len(items) < limitSuggestions || limitSuggestions < 0 {
		limitSuggestions = len(items)
	}
	//Instance.Config.MaxCompleteLines = limitSuggestions
	readline.MaxTabCompleterRows = limitSuggestions

	/*suggest = make([][]rune, len(items))
	for i := range items {
		if len(items[i]) == 0 {
			continue
		}

		if !pt.QuoteSingle && !pt.QuoteDouble && len(items[i]) > 1 && strings.Contains(items[i][:len(items[i])-1], " ") {
			items[i] = strings.Replace(items[i], " ", `\ `, -1)
		}

		if items[i][len(items[i])-1] == '/' || items[i][len(items[i])-1] == '=' {
			suggest[i] = []rune(items[i])
		} else {
			suggest[i] = []rune(items[i] + " ")
		}
	}*/

	for i := range items {
		if len(items[i]) == 0 {
			items[i] = " "
		}
		if items[i][len(items[i])-1] != ' ' && items[i][len(items[i])-1] != '=' && items[i][len(items[i])-1] != '/' {
			items[i] += " "
		}
	}

	return
}

func syntaxCompletion(line []rune, pos int) ([]rune, int) {
	pt, _ := parse(line)
	switch {
	case pt.QuoteSingle:
		if pos < len(line)-1 || line[pos] != '\'' {
			return append(line, '\''), pos
		}

	case pt.QuoteDouble:
		if pos < len(line)-1 || line[pos] != '"' {
			return append(line, '"'), pos
		}

	case pt.Bracket > 0:
		if pos < len(line)-1 || line[pos] != '{' {
			return append(line, '}'), pos
		}

	case pos > 0 && line[pos-1] == '[':
		if pos < len(line)-1 {
			r := append(line[:pos+1], ']')
			return append(r, line[pos+2:]...), pos
		}
		return append(line, ']'), pos

	}
	return line, pos
}

func hintText(line []rune, pos int) []rune {
	r, err := history.ExpandVariables(line)
	if err != nil {
		return []rune("Error: " + err.Error())
	}

	vars := variables.Expand(r)
	disclaimer := []rune{}
	if string(r) != string(vars) {
		disclaimer = []rune("(example only) ")
	}
	r = append(disclaimer, vars...)
	if string(line) == string(r) {
		r = []rune{}
	}

	if len(r) > 0 {
		return r
	}

	pt, _ := parse(line)
	s := manDescript[pt.FuncName]
	if s != "" && s != "!" {
		return []rune(manDescript[pt.FuncName])
	}
	if s == "!" {
		return []rune{}
	}
	f := man.GetManPages(pt.FuncName)
	r = []rune(man.ParseDescription(f))
	if len(r) == 0 {
		manDescript[pt.FuncName] = "!"
	} else {
		manDescript[pt.FuncName] = string(r)
	}
	return r
}
