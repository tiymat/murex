// +build !windows

package man

import (
	"testing"
)

// TestMan tests the builtins package
func TestMan(t *testing.T) {
	files := GetManPages("cat")
	if len(files) == 0 {
		t.Error("Could not find any man pages")
	}

	flags := ParseFlags(files)
	if len(flags) == 0 {
		t.Error("No flags returned for `cat`")
	}

	s := ParseDescription(files)
	if s == "" {
		t.Error("No description returned for `cat`")
	}
}