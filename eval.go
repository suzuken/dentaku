package dentaku

import (
	"strconv"
)

func Eval(e Expression) int {
	switch e.(type) {
	case BinOpExpr:
		left := Eval(e.(BinOpExpr).left)
		right := Eval(e.(BinOpExpr).right)
		switch e.(BinOpExpr).operator {
		case '+':
			return left + right
		case '-':
			return left - right
		case '*':
			return left * right
		case '/':
			return left / right
		case '^':
			r := 1
			for i := 0; i < right; i++ {
				r = r * left
			}
			return r
		}
	case NumExpr:
		i, _ := strconv.Atoi(e.(NumExpr).literal)
		return i
	}
	return 0
}
