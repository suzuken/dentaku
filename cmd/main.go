package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/suzuken/dentaku"
)

func main() {
	l := dentaku.New(os.Stdout)
	l.Init(strings.NewReader(os.Args[1]))
	dentaku.Parse(l)
	fmt.Printf("%#v\n", l.Result)
	fmt.Printf("%#v\n", dentaku.Eval(l.Result))
}
