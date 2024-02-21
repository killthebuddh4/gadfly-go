package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) root(parent *types.Expression, withSignature bool) (*types.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing block for lexeme:", p.previous().Text)
	}

	var operator types.Operator

	if isSchema(p.previous()) {
		operator = types.Operator{
			Type:  "schema",
			Value: p.previous().Text,
		}
	} else {
		operator = types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}
	}

	root := types.NewExpression(parent, operator, []*types.Expression{})

	parameters := []*types.Expression{}

	if withSignature {
		for accept(p, isIdentifier) {
			param := types.Operator{
				Type:  "identifier",
				Value: p.previous().Text,
			}

			paramExp := types.NewExpression(nil, param, []*types.Expression{})

			colon := types.Operator{
				Type:  ":",
				Value: ":",
			}

			var schema types.Operator

			if !accept(p, isColon) {
				schema = types.Operator{
					Type:  "identifier",
					Value: "Identity",
				}
			} else {
				if !(accept(p, isIdentifier) || accept(p, isSchema)) {
					return nil, errors.New("expected identifier after colon")
				}

				schema = types.Operator{
					Type:  "identifier",
					Value: p.previous().Text,
				}
			}

			schemaExp := types.NewExpression(nil, schema, []*types.Expression{})

			validator := types.NewExpression(&root, colon, []*types.Expression{&paramExp, &schemaExp})

			parameters = append(parameters, &validator)
		}

		types.Parameterize(&root, parameters)

		if !accept(p, isEndSignature) {
			return nil, errors.New("expected closing parenthesis")
		}

		if !accept(p, isExpression) {
			return nil, errors.New("expected expression after signature")
		}
	}

	if isSchema(p.previous()) {
		operator = types.Operator{
			Type:  "schema",
			Value: p.previous().Text,
		}
	} else {
		operator = types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}
	}

	root.Operator = operator

	switch operator.Type {
	case "def", "let", "signal", "emit", "on", "catch":
		if !accept(p, isIdentifier) {
			return nil, errors.New("expected identifier after operator")
		}

		idOp := types.Operator{
			Type:  "string",
			Value: p.previous().Text,
		}

		idExp := types.NewExpression(&root, idOp, []*types.Expression{})

		root.Children = append(root.Children, &idExp)
	}

	var endPredicates []Predicate = []Predicate{}

	switch operator.Type {
	case "when", "if":
		endPredicates = append(endPredicates, isThen)
	default:
		endPredicates = append(endPredicates, isEnd, isCatch)
	}

	for {
		if accept(p, endPredicates...) {
			break
		}

		if p.isAtEnd() {
			return nil, errors.New("expected end of expression")
		}

		err := p.parse(&root)

		if err != nil {
			return nil, err
		}
	}

	if operator.Type == "when" && p.previous().Text == "then" {
		thenExp, err := p.sibling(&root)

		if err != nil {
			return nil, err
		}

		root.Siblings = append(root.Siblings, thenExp)
	}

	if operator.Type == "if" && p.previous().Text == "then" {
		thenExp, err := p.sibling(&root)

		if err != nil {
			return nil, err
		}

		root.Siblings = append(root.Siblings, thenExp)

		elseExp, err := p.sibling(&root)

		if err != nil {
			return nil, err
		}

		root.Siblings = append(root.Siblings, elseExp)
	}

	for p.previous().Text == "catch" {
		catchExp, err := p.sibling(&root)

		if err != nil {
			return nil, err
		}

		root.Siblings = append(root.Siblings, catchExp)
	}

	if accept(p, isReturn) {
		if !(accept(p, isIdentifier) || accept(p, isSchema)) {
			return nil, errors.New("expected identifier after arrow")
		}

		schema := types.Operator{
			Type:  "identifier",
			Value: p.previous().Text,
		}

		schemaExp := types.NewExpression(nil, schema, []*types.Expression{})

		types.Returnize(&root, []*types.Expression{&schemaExp})

		if !accept(p, isEndSignature) {
			return nil, errors.New("expected end of return signature")
		}
	}

	return &root, nil
}
