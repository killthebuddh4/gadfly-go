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

	root := types.NewExpression(parent, types.Operator{}, []*types.Expression{})

	if accept(p, isSignature) {
		signature, err := p.signature(parent)

		if err != nil {
			return err
		}

		root.Parameters = append(root.Parameters, signature...)
	}

	if accept(p, isSchema) {
		root.Operator = types.Operator{
			Type:  "schema",
			Value: p.previous().Text,
		}

		err := p.kw(&root)

		if err != nil {
			return err
		}

		parent.Keyword = append(parent.Keyword, &root)
	} else if accept(p, isExpression) {
		root.Operator = types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}

		err := p.kw(&root)

		if err != nil {
			return err
		}

		if root.Operator.Type == "when" {
			if p.previous().Text != "then" {
				return errors.New(":: parse :: expected then after when")
			}

			thenExp, err := p.sibling(&root)

			if err != nil {
				return err
			}

			root.Siblings = append(root.Siblings, thenExp)
		} else if root.Operator.Type == "if" {
			if p.previous().Text != "then" {
				return errors.New(":: parse :: expected then after if")
			}

			thenExp, err := p.sibling(&root)

			if err != nil {
				return err
			}

			root.Siblings = append(root.Siblings, thenExp)

			if p.previous().Text != "else" {
				return errors.New(":: parse :: expected else after then")
			}

			elseExp, err := p.sibling(&root)

			if err != nil {
				return err
			}

			root.Siblings = append(root.Siblings, elseExp)
		} else if root.Operator.Type == "def" || root.Operator.Type == "let" {
			if p.previous().Text != "value" {
				return errors.New(":: parse :: expected value after " + root.Operator.Type)
			}

			valueExp, err := p.sibling(&root)

			if err != nil {
				return err
			}

			root.Siblings = append(root.Siblings, valueExp)
		}

		for p.previous().Text == "catch" {
			catchExp, err := p.sibling(&root)

			if err != nil {
				return err
			}

			root.Siblings = append(root.Siblings, catchExp)
		}

		if accept(p, isReturn) {
			if !accept(p, isSchema) {
				return errors.New("expected schema idf after arrow")
			}

			schema := types.Operator{
				Type:  "identifier",
				Value: p.previous().Text,
			}

			schemaExp := types.NewExpression(nil, schema, []*types.Expression{})

			types.Returnize(&root, []*types.Expression{&schemaExp})

			if !accept(p, isEndSignature) {
				return errors.New("expected end of return signature")
			}
		}

		parent.Keyword = append(parent.Keyword, &root)
	} else {
		child, err := p.predicate(parent)

		if err != nil {
			return err
		}

		child.Parent = parent
		parent.Keyword = append(parent.Keyword, child)
	}

	return nil
}
