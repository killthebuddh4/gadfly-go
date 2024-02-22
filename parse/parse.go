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
		err := p.parent(root)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parent(parent *types.Expression) error {
	if accept(p, isExpression) {
		def, err := GetExpDef(p.previous().Text)

		if err != nil {
			return err
		}

		exp := types.Expression{
			Parent:       parent,
			Def:          &def,
			Operator:     types.Operator{Type: p.previous().Text, Value: p.previous().Text},
			Parameters:   []*types.Expression{},
			Catches:      []*types.Expression{},
			Returns:      []*types.Expression{},
			Trajectories: []*types.Trajectory{},
		}

		parent.Parameters = append(parent.Parameters, &exp)

		for _, paramDef := range def.Parameters {
			fmt.Println("DEF FOR ", exp.Operator.Value, " ", paramDef.Name)
			paramExp := types.Expression{
				Parent:       parent,
				Def:          &EMPTY,
				Operator:     types.Operator{Type: p.previous().Text, Value: p.previous().Text},
				Parameters:   []*types.Expression{},
				Catches:      []*types.Expression{},
				Returns:      []*types.Expression{},
				Trajectories: []*types.Trajectory{},
			}

			exp.Parameters = append(exp.Parameters, &paramExp)

			fmt.Println("LENG ", len(exp.Parameters))

			for {
				if a(p, p.read().Text, paramDef.EndWords) {
					fmt.Println("BREAKING ", paramExp.Operator.Value, " ", p.previous().Text)
					break
				} else {
					fmt.Println("NOT BREAKING ", paramExp.Operator.Value, " ", p.read().Text)
				}

				if p.isAtEnd() {
					return errors.New(":: BLOCK :: expected end of BLOCK")
				}

				err := p.parent(&paramExp)

				if err != nil {
					return err
				}
			}

			if err != nil {
				return nil
			}
		}

		if err != nil {
			return err
		}

	} else {
		child, err := p.predicate(parent)

		if err != nil {
			return err
		}

		child.Parent = parent
		parent.Parameters = append(parent.Parameters, child)
	}

	return nil
}
