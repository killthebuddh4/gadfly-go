package main

func (p *Parser) ParseIdentifier(parent *Expression, operator Token) error {
	root := Expr(parent, VARIANTS.Call, operator)

	defn, err := Resolve(parent, operator.Lexeme)

	if err != nil {
		return err
	}

	if defn.Arity == 0 && !defn.Variadic {
		return nil
	}

	for {
		if accept(p, isEnd) {
			break
		}

		if p.isAtEnd() {
			break
		}

		err := p.expression(&root)

		if err != nil {
			return nil
		}
	}

	return nil
}
