package dentaku

import (
	"strconv"
)

func Eval(e Expression) int {
	switch e.(type) {
	case BinOpExpr:
		left := Eval(e.(BinOpExpr).left)
		right := Eval(e.(BinOpExpr).right)
		return left + right
	case NumExpr:
		i, _ := strconv.Atoi(e.(NumExpr).literal)
		return i
	}
	return 0
}
