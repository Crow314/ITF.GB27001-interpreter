package ast

func Evaluate(expr IExpr, env map[string]*MExpr) IExpr {
	if expr.Type() == TBinExpr {
		binExpr := expr.(*MBinExpr)
		op := binExpr.Op

		switch op {
		case TAdd:
			return new(MInt).Init(Evaluate(binExpr.Lhs, env).(*MInt).Value + Evaluate(binExpr.Rhs, env).(*MInt).Value)

		case TSub:
			return new(MInt).Init(Evaluate(binExpr.Lhs, env).(*MInt).Value - Evaluate(binExpr.Rhs, env).(*MInt).Value)

		case TMul:
			return new(MInt).Init(Evaluate(binExpr.Lhs, env).(*MInt).Value * Evaluate(binExpr.Rhs, env).(*MInt).Value)

		case TDiv:
			return new(MInt).Init(Evaluate(binExpr.Lhs, env).(*MInt).Value / Evaluate(binExpr.Rhs, env).(*MInt).Value)
		}
	} else if expr.Type() == TSeq {
		seq := expr.(*MSeq).Sequence
		var e *MExpr

		for _, *e = range seq {
			Evaluate(e, env)
		}

		return e
	} else if expr.Type() == TInt {
		return expr
	}

	panic("Unexpected Error")
}
