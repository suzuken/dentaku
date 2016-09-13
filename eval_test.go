package dentaku

import (
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	type c struct {
		in  string
		out int
	}
	cases := []c{
		{in: "1 + 1", out: 2},
		{in: "1 + 2 * 10", out: 21},
		{in: "1 + 10 / 5", out: 3},
		{in: "1 + 10 ^ 2", out: 101},
	}
	for _, r := range cases {
		l := new(Lexer)
		l.Init(strings.NewReader(r.in))
		t.Logf("test: %s = %d", r.in, r.out)
		yyParse(l)
		if actual := Eval(l.Result); actual != r.out {
			t.Fatalf("%s should %d but got %d", r.in, r.out, actual)
		}
	}
}
