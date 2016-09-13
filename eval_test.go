package dentaku

import (
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	l := new(Lexer)
	l.Init(strings.NewReader(`1 + 1`))
	yyParse(l)
	if actual := Eval(l.Result); actual != 2 {
		t.Fatalf("1 + 1 should 2 but got %d", actual)
	}
}

func TestEval2(t *testing.T) {
	l := new(Lexer)
	l.Init(strings.NewReader(`1 + 2 * 10`))
	yyParse(l)
	if actual := Eval(l.Result); actual != 21 {
		t.Fatalf("1 + 2 * 10 should 21 but got %d", actual)
	}
}
