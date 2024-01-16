package parse

import (
	"errors"
	"fmt"
	"os"

	exp "github.com/killthebuddh4/gadflai/expression"
)

func (p *Parser) lambda(parent *exp.Expression) (*exp.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing lambda for lexeme:", p.previous().Text)
	}

	operator, err := exp.NewOperator(p.previous().Text)

	if err != nil {
		return nil, err
	}

	root := exp.NewExpression(parent, operator, []*exp.Expression{})

	if accept(p, isPipe) {
		parameters := []string{}

		for accept(p, isIdentifier) {
			parameters = append(parameters, p.previous().Text)
		}

		if !accept(p, isPipe) {
			return nil, errors.New("expected closing pipe")
		}

		exp.Parameterize(&root, parameters)
	}

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
