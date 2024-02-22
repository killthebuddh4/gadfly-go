package types

import "fmt"

type Schema func(value Value) (Value, error)

type Predicate func(lexeme Lexeme) bool

type Parameter struct {
	Name     string
	IsThunk  bool
	EndWords []string
}

type ExpDef struct {
	Parameters []Parameter
}

type Expression struct {
	Parent       *Expression
	Operator     Operator
	Def          *ExpDef
	Parameters   []*Expression
	Catches      []*Expression
	Returns      []*Expression
	Trajectories []*Trajectory
}

func NewExpression(parent *Expression, operator Operator, children []*Expression) Expression {
	return Expression{
		Parent:       parent,
		Operator:     operator,
		Parameters:   children,
		Catches:      []*Expression{},
		Returns:      []*Expression{},
		Trajectories: []*Trajectory{},
	}
}

func Returnize(parent *Expression, returns []*Expression) {
	parent.Returns = returns
}

func Print(expression Expression, depth int) error {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	fmt.Println("<" + expression.Operator.Type + ": " + expression.Operator.Value + ">")

	for _, child := range expression.Parameters {
		Print(*child, depth+1)
	}

	return nil
}
