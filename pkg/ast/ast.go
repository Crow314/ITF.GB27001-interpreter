package ast

type AstType int
type OpType int

const (
	TBinExpr = AstType(iota)
	TInt
	TSeq
)

const (
	TAdd = OpType(iota)
	TSub
	TMul
	TDiv
)

type IAst interface {
	Type() AstType
}

type IExpr interface {
	IAst
}

type MExpr struct {
	AstType AstType
}

type MBinExpr struct {
	MExpr
	Op  OpType
	Lhs IExpr
	Rhs IExpr
}

type MInt struct {
	MExpr
	Value int
}

type MSeq struct {
	MExpr
	Sequence []MExpr
}

func (expr *MExpr) Type() AstType {
	return expr.AstType
}

func (expr *MExpr) SetType(astType AstType) {
	expr.AstType = astType
}

func (expr *MBinExpr) Init(op OpType, lhs IExpr, rhs IExpr) *MBinExpr {
	expr.SetType(TBinExpr)
	expr.Op = op
	expr.Lhs = lhs
	expr.Rhs = rhs

	return expr
}

func (expr *MInt) Init(value int) *MInt {
	expr.SetType(TInt)
	expr.Value = value

	return expr
}
