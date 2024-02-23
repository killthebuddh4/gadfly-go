package parse

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

func Parse(root *types.Expression, lexemes []types.Lexeme) error {
	p := Parser{
		Lexemes: lexemes,

		Current: 0,
	}

	for !p.isAtEnd() {
		_, err := p.parent(root)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parent(parent *types.Expression) (*types.Expression, error) {
	if accept(p, isExpression) {
		fmt.Println("ACCEPTED", p.previous().Text)
		exp := types.Expression{
			Parent: parent,
			Operator: types.Operator{
				Type:  p.previous().Text,
				Value: p.previous().Text,
			},
			Parameters:   []*types.Expression{},
			Catches:      []*types.Expression{},
			Returns:      []*types.Expression{},
			Trajectories: []*types.Trajectory{},
		}

		for {
			operator := types.Operator{
				Type:  p.previous().Text,
				Value: p.previous().Text,
			}

			param := types.Expression{
				Parent:       parent,
				Operator:     operator,
				Parameters:   []*types.Expression{},
				Catches:      []*types.Expression{},
				Returns:      []*types.Expression{},
				Trajectories: []*types.Trajectory{},
			}

			endWords, err := GetEndwords(operator)

			if err != nil {
				return nil, err
			}

			if acc(p, p.read().Text, endWords) {
				if p.previous().Text == "end" {
					fmt.Println("Breaking on end")
					break
				} else {
					p.backup()
				}
			}

			if p.isAtEnd() {
				return nil, errors.New(":: BLOCK :: expected end of BLOCK for operator <" + operator.Type)
			}

			child, err := p.parent(&param)

			if err != nil {
				return nil, err
			}

			fmt.Println("CHILD", child.Operator.Type)

			if child.Operator.Type == "catch" {
				exp.Catches = append(exp.Catches, child)
			} else {
				exp.Parameters = append(exp.Parameters, child)
			}

			if err != nil {
				return nil, err
			}
		}

		return &exp, nil
	} else {
		return p.predicate(parent)
	}

	return nil, errors.New(":: PARENT :: expected expression or predicate, but got <" + p.read().Text + ">")
}
