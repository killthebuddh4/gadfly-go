package expression

import "fmt"

type Expression struct {
	Parent     *Expression
	Operator   Operator
	Children   []*Expression
	Parameters []string
}

func NewExpression(parent *Expression, operator Operator, children []*Expression) Expression {
	return Expression{
		Parent:     parent,
		Operator:   operator,
		Children:   children,
		Parameters: []string{},
	}
}

func Expand(parent *Expression, children []*Expression) {
	parent.Children = append(parent.Children, children...)
}

func Parameterize(parent *Expression, parameters []string) {
	parent.Parameters = parameters
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
