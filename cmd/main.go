package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/suzuken/dentaku"
)

func main() {
	l := new(dentaku.Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	dentaku.Parse(l)
	fmt.Printf("%#v\n", l.Result)
}
