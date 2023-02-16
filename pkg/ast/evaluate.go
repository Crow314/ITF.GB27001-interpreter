package ast

func Evaluate(expr IExpr) IExpr {
	if binExpr, ok := expr.(*MBinExpr); ok {
		op := binExpr.Op

		switch op {
		case TAdd:
			res := new(MInt).Init(Evaluate(binExpr.Lhs).(*MInt).Value + Evaluate(binExpr.Rhs).(*MInt).Value)
			return &res

		case TSub:
			res := new(MInt).Init(Evaluate(binExpr.Lhs).(*MInt).Value - Evaluate(binExpr.Rhs).(*MInt).Value)
			return &res

		case TMul:
			res := new(MInt).Init(Evaluate(binExpr.Lhs).(*MInt).Value * Evaluate(binExpr.Rhs).(*MInt).Value)
			return &res

		case TDiv:
			res := new(MInt).Init(Evaluate(binExpr.Lhs).(*MInt).Value / Evaluate(binExpr.Rhs).(*MInt).Value)
			return &res
		}
	} else if _, ok := expr.(*MInt); ok {
		return expr
	}

	panic("Unexpected Error")
}
