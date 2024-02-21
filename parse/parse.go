package parse

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Parse(root *types.Expression, lexemes []types.Lexeme) error {
	p := Parser{
		Lexemes: lexemes,

		Current: 0,
	}

	for !p.isAtEnd() {
		err := p.parse(root)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parse(parent *types.Expression) error {
	if parent == nil {
		return errors.New("cannot parse expression with nil parent")
	}

	// WARNING := The order of these two expressions matters. If you check
	// expression first, then withSig will pick up the signature of the NEXT
	// expression if it exists.

	withSig := accept(p, isSignature)
	isExpr := accept(p, isExpression) || accept(p, isSchema)

	if isExpr || withSig {
		child, err := p.root(parent, withSig)

		if err != nil {
			return err
		}

		child.Parent = parent
		parent.Children = append(parent.Children, child)
	} else {
		child, err := p.predicate(parent)

		if err != nil {
			return err
		}

		child.Parent = parent
		parent.Children = append(parent.Children, child)
	}

	return nil
}
