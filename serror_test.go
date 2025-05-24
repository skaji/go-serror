package serror_test

import (
	"strings"
	"testing"

	"github.com/skaji/go-serror"
)

func testFunc1() error {
	return serror.Error("oops")
}

func TestBasic(t *testing.T) {
	err1 := testFunc1()
	err2 := serror.Wrap(err1)
	lines := strings.Split(err2.Error(), "\n")
	if len(lines) != 3 {
		t.Fatal()
	}
	if lines[0] != "oops" {
		t.Fatal()
	}
	if !strings.HasSuffix(lines[1], "serror_test.go:11") {
		t.Fatal()
	}
	if !strings.HasSuffix(lines[2], "serror_test.go:16") {
		t.Fatal()
	}
	t.Log(err2)
}
