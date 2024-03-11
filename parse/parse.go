package parse

import (
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func Parse(root *types.Expression, lexemes []types.Lexeme) error {
	p := Parser{
		Lexemes: lexemes,

		Current: 0,
	}

	block := types.Block{}

	for !p.isAtEnd() {
		exp, err := p.parse(root)

		if err != nil {
			return err
		}

		block = append(block, exp)
	}

	root.Arguments["program"] = block

	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		types.Print(root, 0)
	}

	return nil
}

func (p *Parser) parse(parent *types.Expression) (*types.Expression, error) {
	if accept(p, isFunction) {
		return p.function(parent)
	} else {
		return p.predicate(parent)
	}
}
