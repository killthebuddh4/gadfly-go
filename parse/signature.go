package parse

import (
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) signature(owner *types.Expression) ([]*types.Expression, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		fmt.Println("Parsing signature for lexeme:", p.previous().Text)
	}

	parameters := []*types.Expression{}

	// for accept(p, isIdentifier) {
	// 	param := types.Operator{
	// 		Type:  "identifier",
	// 		Value: p.previous().Text,
	// 	}

	// 	// paramExp := types.NewExpression(nil, param, []*types.Expression{})

	// 	colon := types.Operator{
	// 		Type:  ":",
	// 		Value: ":",
	// 	}

	// 	var schemaOperator types.Operator

	// 	// If there's no schema identifier, we default to the Identity schema,
	// 	// otherwise we parse the schema's identifier.

	// 	if !accept(p, isColon) {
	// 		schemaOperator = types.Operator{
	// 			Type:  "identifier",
	// 			Value: "Identity",
	// 		}
	// 	} else {
	// 		if !(accept(p, isSchema)) {
	// 			return nil, errors.New("expected identifier after colon")
	// 		}

	// 		schemaOperator = types.Operator{
	// 			Type:  "identifier",
	// 			Value: p.previous().Text,
	// 		}
	// 	}

	// 	schemaExp := types.NewExpression(nil, schemaOperator, []*types.Expression{})

	// 	colonExp := types.NewExpression(nil, colon, []*types.Expression{&paramExp, &schemaExp})

	// 	parameters = append(parameters, &colonExp)
	// }

	// if !accept(p, isEndSignature) {
	// 	return nil, errors.New("expected closing parenthesis")
	// }

	return parameters, nil
}
