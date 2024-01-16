package parse

import (
	"fmt"
	"os"

	exp "github.com/killthebuddh4/gadflai/expression"
)

func (p *Parser) block(parent *exp.Expression) (*exp.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing block for lexeme:", p.previous().Text)
	}

	operator, err := exp.NewOperator(p.previous().Text)

	if err != nil {
		return nil, err
	}

	root := exp.NewExpression(parent, operator, []*exp.Expression{})

	for {
		if accept(p, isEnd) {
			break
		}

		if p.isAtEnd() {
			break
		}

		err := p.expression(&root)

		if err != nil {
			return nil, err
		}
	}

	return &root, nil
}
