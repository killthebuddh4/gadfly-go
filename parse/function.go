package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) function(parent *types.Expression) (*types.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println(":: PARSE :: function :: ", p.previous().Text)
	}

	exp := types.Expression{
		Parent: parent,
		Operator: types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		},
		Signature: &types.Signature{
			Parameters: make(map[string]string),
			Returns:    []*types.Expression{},
		},
		Arguments:    make(map[string][]*types.Expression),
		Catches:      []types.Block{},
		Trajectories: []*types.Trajectory{},
	}

	if accept(p, isSignature) {
		for accept(p, isIdentifier) {
			name := p.previous().Text

			var schema string

			if !accept(p, isColon) {
				schema = "Identity"
			} else {
				if !accept(p, isSchema) {
					return nil, errors.New("expected schema after colon")
				}

				schema = p.previous().Text
			}

			exp.Signature.Parameters[name] = schema
		}

		if !accept(p, isEndSignature) {
			return nil, errors.New("expected closing parenthesis")
		}
	}

	names, err := GetParameters(exp.Operator)

	if err != nil {
		return nil, err
	}

	if len(names) == 0 || names[0] != exp.Operator.Type {
		return nil, errors.New("first parameter must be the function name")
	}

	for i, name := range names {
		endWords := []string{}

		// if there is another parameter then it is the end word,
		// otherwise it is either catch or end
		if i+1 < len(names) {
			endWords = append(endWords, names[i+1])
		} else {
			endWords = append(endWords, "catch", "end")
		}

		block, err := p.block(parent, endWords)

		if err != nil {
			return nil, err
		}

		exp.Arguments[name] = block
	}

	// for !accepted(p, isEnd) {
	// 	if !accept(p, isCatch) {
	// 		return nil, errors.New("expected catch or end")
	// 	}

	// 	catch, err := p.block(parent, []string{"end"})

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	exp.Catches = append(exp.Catches, catch)
	// }

	if debug {
		fmt.Println(":: DONE :: PARSE :: function :: ", exp.Operator.Value)
	}

	return &exp, nil
}
