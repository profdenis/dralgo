package ast

func (op BinaryOp) precedence() int {
	switch op {
	case OpPow:
		return 5
	case OpMul, OpDiv, OpMod:
		return 4
	case OpAdd, OpSub:
		return 3
	case OpLT, OpLTE, OpGT, OpGTE:
		return 2
	case OpEQ, OpNEQ:
		return 1
	default:
		return 0
	}
}

// NeedsParens indique si l'opérande enfant doit être parenthésé
// lorsqu'il est placé sous l'opérateur parent (côté gauche ou droit).
func NeedsParens(parent BinaryOp, child Expr, rightSide bool) bool {
	inner, ok := child.(BinaryExpr)
	if !ok {
		return false
	}
	if inner.Op.precedence() < parent.precedence() {
		return true
	}
	return rightSide && inner.Op.precedence() == parent.precedence()
}
