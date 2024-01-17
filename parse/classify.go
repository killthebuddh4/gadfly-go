package parse

import "github.com/killthebuddh4/gadflai/lex"

/*
 * Note that these classifiers are very much coupled to the grammar and must be
 * used with a clear understanding of the grammar. This is a very direct
 * violation of encapsulation, but I think it's fine in this case.
 */

func isLambda(lexeme lex.Lexeme) bool {
	return lexeme.Text == "fn"
}

func isBlock(lexeme lex.Lexeme) bool {
	switch lexeme.Text {
	case "do":
		return true
	case "def", "@", "val", "let":
		return true
	case "if", "and", "or", "while":
		return true
	case "array", "get", "set", "for", "map", "filter", "reduce", "push", "pop", "shift", "unshift", "segment", "find", "splice", "reverse":
		return true
	case "record", "merge", "delete", "keys", "values", "entries", "read", "write", "extract":
		return true
	case "puts", "chars":
		return true
	default:
		return false
	}
}

func isEnd(lexeme lex.Lexeme) bool {
	return lexeme.Text == "end"
}

func isTrue(lexeme lex.Lexeme) bool {
	return lexeme.Text == "true"
}

func isFalse(lexeme lex.Lexeme) bool {
	return lexeme.Text == "false"
}

func isNil(lexeme lex.Lexeme) bool {
	return lexeme.Text == "nil"
}

func isIdentifier(lexeme lex.Lexeme) bool {
	switch string(lexeme.Text[0]) {
	case "_", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
		return true
	default:
		return false
	}
}

func isPipe(lexeme lex.Lexeme) bool {
	return lexeme.Text == "|"
}

func isLogical(lexeme lex.Lexeme) bool {
	switch lexeme.Text {
	case "&&", "||":
		return true
	default:
		return false
	}
}

func isEquality(lexeme lex.Lexeme) bool {
	switch lexeme.Text {
	case "==", "!=":
		return true
	default:
		return false
	}
}

func isComparison(lexeme lex.Lexeme) bool {
	switch lexeme.Text {
	case "<", "<=", ">", ">=":
		return true
	default:
		return false
	}
}

func isTerm(lexeme lex.Lexeme) bool {
	switch lexeme.Text {
	case "+", "-":
		return true
	default:
		return false
	}
}

func isFactor(lexeme lex.Lexeme) bool {
	switch lexeme.Text {
	case "*", "/":
		return true
	default:
		return false
	}
}

func isUnary(lexeme lex.Lexeme) bool {
	switch lexeme.Text {
	case "-", "!":
		return true
	default:
		return false
	}
}

func isString(lexeme lex.Lexeme) bool {
	return string(lexeme.Text[0]) == "\""
}

func isNumber(lexeme lex.Lexeme) bool {
	switch string(lexeme.Text[0]) {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	default:
		return false
	}
}

// Is an atom anything that is not more specific?
func isAtom(lexeme lex.Lexeme) bool {
	if isBlock(lexeme) {
		return false
	}

	if isLambda(lexeme) {
		return false
	}

	if isPipe(lexeme) {
		return false
	}

	if isEnd(lexeme) {
		return false
	}

	if isLogical(lexeme) {
		return false
	}

	if isEquality(lexeme) {
		return false
	}

	if isComparison(lexeme) {
		return false
	}

	if isTerm(lexeme) {
		return false
	}

	if isFactor(lexeme) {
		return false
	}

	if isUnary(lexeme) {
		return false
	}

	if isTrue(lexeme) {
		return true
	}

	if isFalse(lexeme) {
		return true
	}

	if isNil(lexeme) {
		return true
	}

	if isIdentifier(lexeme) {
		return true
	}

	if isString(lexeme) {
		return true
	}

	if isNumber(lexeme) {
		return true
	}

	return false
}
