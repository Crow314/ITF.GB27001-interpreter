package main

import "testing"
import "github.com/Crow314/ITF.GB27001-interpreter/pkg/ast"

func TestAdd(t *testing.T) {
	expr := new(ast.MBinExpr)
	lhs := new(ast.MInt)
	rhs := new(ast.MInt)

	lhs.Init(3)
	rhs.Init(7)
	expr.Init(ast.TAdd, lhs, rhs)

	res := ast.Evaluate(expr).(*ast.MInt)

	if res.Type() != ast.TInt {
		t.Errorf("Result is not TInt")
	}

	if res.Value != 10 {
		t.Errorf("3 + 7\nExpected: 10\nResult: %d", res.Value)
	}
}

func TestSub(t *testing.T) {
	expr := new(ast.MBinExpr)
	lhs := new(ast.MInt)
	rhs := new(ast.MInt)

	lhs.Init(4)
	rhs.Init(9)
	expr.Init(ast.TSub, lhs, rhs)

	res := ast.Evaluate(expr).(*ast.MInt)

	if res.Type() != ast.TInt {
		t.Errorf("Result is not TInt")
	}

	if res.Value != -5 {
		t.Errorf("4 - 9\nExpected: -5\nResult: %d", res.Value)
	}
}

func TestMul(t *testing.T) {
	expr := new(ast.MBinExpr)
	lhs := new(ast.MInt)
	rhs := new(ast.MInt)

	lhs.Init(20)
	rhs.Init(12)
	expr.Init(ast.TMul, lhs, rhs)

	res := ast.Evaluate(expr).(*ast.MInt)

	if res.Type() != ast.TInt {
		t.Errorf("Result is not TInt")
	}

	if res.Value != 240 {
		t.Errorf("20 * 12\nExpected: 240\nResult: %d", res.Value)
	}
}

func TestDiv(t *testing.T) {
	expr := new(ast.MBinExpr)
	lhs := new(ast.MInt)
	rhs := new(ast.MInt)

	lhs.Init(100)
	rhs.Init(10)
	expr.Init(ast.TDiv, lhs, rhs)

	res := ast.Evaluate(expr).(*ast.MInt)

	if res.Type() != ast.TInt {
		t.Errorf("Result is not TInt")
	}

	if res.Value != 10 {
		t.Errorf("100 / 10\nExpected: 10\nResult: %d", res.Value)
	}
}
