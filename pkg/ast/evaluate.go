package ast

func Evaluate(expr IExpr) IExpr {
	if expr.Type() == TBinExpr {
		binExpr := expr.(*MBinExpr)
		op := binExpr.Op

		switch op {
		case TAdd:
			return new(MInt).Init(Evaluate(binExpr.Lhs).(*MInt).Value + Evaluate(binExpr.Rhs).(*MInt).Value)

		case TSub:
			return new(MInt).Init(Evaluate(binExpr.Lhs).(*MInt).Value - Evaluate(binExpr.Rhs).(*MInt).Value)

		case TMul:
			return new(MInt).Init(Evaluate(binExpr.Lhs).(*MInt).Value * Evaluate(binExpr.Rhs).(*MInt).Value)

		case TDiv:
			return new(MInt).Init(Evaluate(binExpr.Lhs).(*MInt).Value / Evaluate(binExpr.Rhs).(*MInt).Value)
		}
	} else if expr.Type() == TInt {
		return expr
	}

	panic("Unexpected Error")
}
