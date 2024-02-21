package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) sibling(parent *types.Expression) (*types.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing sibling for lexeme:", p.previous().Text)
	}

	operator := types.Operator{
		Type:  p.previous().Text,
		Value: p.previous().Text,
	}

	root := types.NewExpression(parent, operator, []*types.Expression{})

	var endPredicates []Predicate = []Predicate{}

	if operator.Type == "then" {
		if parent.Operator.Type == "when" {
			endPredicates = append(endPredicates, isEnd)
		} else if parent.Operator.Type == "if" {
			endPredicates = append(endPredicates, isElse)
		} else {
			return nil, errors.New(":: parse :: then must be inside when or if")
		}
	} else {
		endPredicates = append(endPredicates, isEnd, isCatch)
	}

	for {
		if accept(p, endPredicates...) {
			break
		}

		if p.isAtEnd() {
			return nil, errors.New(":: sibling :: expected end of expression")
		}

		err := p.parse(&root)

		if err != nil {
			return nil, err
		}
	}

	return &root, nil
}
