package types

import "fmt"

type Expression struct {
	Parent       *Expression
	Operator     Operator
	Keyword      []*Expression
	Siblings     []*Expression
	Parameters   []*Expression
	Returns      []*Expression
	Trajectories []*Trajectory
}

func NewExpression(parent *Expression, operator Operator, children []*Expression) Expression {
	return Expression{
		Parent:       parent,
		Operator:     operator,
		Keyword:      children,
		Siblings:     []*Expression{},
		Parameters:   []*Expression{},
		Returns:      []*Expression{},
		Trajectories: []*Trajectory{},
	}
}

func ExpandExp(parent *Expression, children []*Expression) {
	parent.Keyword = append(parent.Keyword, children...)
}

func Parameterize(parent *Expression, parameters []*Expression) {
	parent.Parameters = parameters
}

func Returnize(parent *Expression, returns []*Expression) {
	parent.Returns = returns
}

func Print(expression Expression, depth int) error {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	fmt.Println("<" + expression.Operator.Type + ": " + expression.Operator.Value + ">")

	for _, child := range expression.Keyword {
		Print(*child, depth+1)
	}

	return nil
}
