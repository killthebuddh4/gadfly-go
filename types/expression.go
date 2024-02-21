package types

import "fmt"

type Expression struct {
	Parent       *Expression
	Operator     Operator
	Children     []*Expression
	Siblings     []*Expression
	Parameters   []*Expression
	Returns      []*Expression
	Trajectories []*Trajectory
}

func NewExpression(parent *Expression, operator Operator, children []*Expression) Expression {
	return Expression{
		Parent:       parent,
		Operator:     operator,
		Children:     children,
		Siblings:     []*Expression{},
		Parameters:   []*Expression{},
		Returns:      []*Expression{},
		Trajectories: []*Trajectory{},
	}
}

func ExpandExp(parent *Expression, children []*Expression) {
	parent.Children = append(parent.Children, children...)
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

	for _, child := range expression.Children {
		Print(*child, depth+1)
	}

	return nil
}
