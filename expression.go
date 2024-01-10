package main

type Lambda func(args ...Value) (Value, error)

type Expression struct {
	Parent       *Expression
	Operator     Token
	Parameters   []string
	Children     []*Expression
	Trajectories []*Trajectory
}

func Expr(parent *Expression, operator Token) Expression {
	return Expression{
		Parent:       parent,
		Operator:     operator,
		Children:     []*Expression{},
		Parameters:   []string{},
		Trajectories: []*Trajectory{},
	}
}

func RootExpr() Expression {
	return Expr(nil, Token{
		Type:   "ROOT",
		Lexeme: "ROOT",
		Start:  0,
		Length: 0,
	})
}
