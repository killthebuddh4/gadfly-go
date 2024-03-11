package parse

import (
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) block(parent *types.Expression, endWords []string) (types.Block, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	startWord := p.previous().Text

	if debug {
		fmt.Println(":: PARSE :: BLOCK :: ", startWord, "with end words", endWords)
	}

	predicates := []Predicate{}

	for _, word := range endWords {
		predicates = append(predicates, func(lexeme types.Lexeme) bool {
			return lexeme.Text == word
		})
	}

	expressions := types.Block{}

	for !accept(p, predicates...) {
		if p.isAtEnd() {
			return nil, fmt.Errorf("expected one of %v, got EOF", endWords)
		}

		expr, err := p.parse(parent)

		if err != nil {
			return nil, err
		}

		expressions = append(expressions, expr)
	}

	if debug {
		fmt.Println(":: DONE :: PARSE :: BLOCK :: ", startWord, "with end:", p.previous().Text)
	}

	return expressions, nil
}
