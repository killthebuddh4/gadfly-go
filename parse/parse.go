package parse

import (
	"errors"

	exp "github.com/killthebuddh4/gadflai/expression"
	"github.com/killthebuddh4/gadflai/lex"
)

func Parse(root *exp.Expression, lexemes []lex.Lexeme) error {
	parser := Parser{
		Lexemes: lexemes,
		Current: 0,
	}

	return parser.program(root)
}

func (p *Parser) program(root *exp.Expression) error {
	for !p.isAtEnd() {
		err := p.expression(root)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) expression(parent *exp.Expression) error {
	if parent == nil {
		return errors.New("cannot parse expression with nil parent")
	}

	if accept(p, isLambda) {
		child, err := p.lambda(parent)

		if err != nil {
			return err
		}

		child.Parent = parent
		parent.Children = append(parent.Children, child)
	} else if accept(p, isBlock) {
		child, err := p.block(parent)

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
