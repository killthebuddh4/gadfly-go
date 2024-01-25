package parse

import (
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) block(parent *types.Expression) (*types.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing block for lexeme:", p.previous().Text)
	}

	operator, err := types.NewOperator(p.previous().Text)

	if err != nil {
		return nil, err
	}

	root := types.NewExpression(parent, operator, []*types.Expression{})

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
