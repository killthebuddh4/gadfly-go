package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) expression(parent *types.Expression) (*types.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing block for lexeme:", p.previous().Text)
	}

	operator, err := types.NewOperator(p.previous().Text)

	if err != nil {
		return nil, err
	}

	root := types.NewExpression(parent, operator, []*types.Expression{})

	if accept(p, isPipe) {
		parameters := []string{}

		for accept(p, isIdentifier) {
			parameters = append(parameters, p.previous().Text)
		}

		if !accept(p, isPipe) {
			return nil, errors.New("expected closing pipe")
		}

		types.Parameterize(&root, parameters)
	}

	for {
		if accept(p, isEnd) {
			break
		}

		if p.isAtEnd() {
			break
		}

		err := p.parse(&root)

		if err != nil {
			return nil, err
		}
	}

	return &root, nil
}
