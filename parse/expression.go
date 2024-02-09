package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) expression(parent *types.Expression) (*types.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing block for lexeme:", p.previous().Text)
	}

	operator, err := types.NewOperator(p.previous().Text)

	if err != nil {
		return nil, err
	}

	root := types.NewExpression(parent, operator, []*types.Expression{})

	parameters := []*types.Expression{}

	if accept(p, isPipe) {
		for accept(p, isIdentifier) {
			param, err := types.NewOperator(p.previous().Text)

			if err != nil {
				return nil, err
			}

			paramExp := types.NewExpression(nil, param, []*types.Expression{})

			var colon types.Operator
			var schema types.Operator

			if !accept(p, isColon) {
				colonOp, err := types.NewOperator(":")

				if err != nil {
					return nil, err
				}

				colon = colonOp

				schemaOp, err := types.NewOperator("Identity")

				if err != nil {
					return nil, err
				}

				schema = schemaOp
			} else {
				colonOp, err := types.NewOperator(p.previous().Text)

				if err != nil {
					return nil, err
				}

				colon = colonOp

				if !(accept(p, isIdentifier) || accept(p, isSchema)) {
					return nil, errors.New("expected identifier after colon")
				}

				schemaOp, err := types.NewOperator(p.previous().Text)

				if err != nil {
					return nil, err
				}

				schema = schemaOp
			}

			schemaExp := types.NewExpression(nil, schema, []*types.Expression{})

			validator := types.NewExpression(&root, colon, []*types.Expression{&paramExp, &schemaExp})

			parameters = append(parameters, &validator)
		}

		types.Parameterize(&root, parameters)

		if !accept(p, isPipe) {
			return nil, errors.New("expected closing pipe")
		}

	}

	for {
		if accept(p, isEnd) {
			break
		}

		if p.isAtEnd() {
			break
		}

		err := p.parse(&root)

		if err != nil {
			return nil, err
		}
	}

	if accept(p, isArrow) {
		if !(accept(p, isIdentifier) || accept(p, isSchema)) {
			return nil, errors.New("expected identifier after arrow")
		}

		schema, err := types.NewOperator(p.previous().Text)

		if err != nil {
			return nil, err
		}

		schemaExp := types.NewExpression(nil, schema, []*types.Expression{})

		types.Returnize(&root, []*types.Expression{&schemaExp})
	}

	return &root, nil
}
