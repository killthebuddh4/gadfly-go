package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) expression(parent *types.Expression, withSignature bool) (*types.Expression, error) {
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
	case "def", "let", "signal", "emit", "on", "throw", "catch":
		if !accept(p, isIdentifier) {
			return nil, errors.New("expected identifier after def")
		}

		idOp := types.Operator{
			Type:  "string",
			Value: p.previous().Text,
		}

		idExp := types.NewExpression(&root, idOp, []*types.Expression{})

		root.Children = append(root.Children, &idExp)
	}

	for {
		if accept(p, isEnd) {
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
