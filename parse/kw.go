package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) kw(root *types.Expression) error {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing kw for lexeme:", p.previous().Text)
	}

	var endPredicates []Predicate = []Predicate{}

	// some keywords have specific siblings that are required, the rest can have
	// end or catch.
	switch root.Operator.Type {
	case "when", "if":
		endPredicates = append(endPredicates, isThen)
	case "def", "let":
		endPredicates = append(endPredicates, isValue)
	default:
		endPredicates = append(endPredicates, isEnd, isCatch)
	}

	for {
		if accept(p, endPredicates...) {
			break
		}

		if p.isAtEnd() {
			return errors.New(":: root :: expected end of expression")
		}

		err := p.parse(root)

		if err != nil {
			return nil
		}
	}

	return nil
}
