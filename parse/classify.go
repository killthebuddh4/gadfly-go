package parse

import "github.com/killthebuddh4/gadflai/types"

/*
 * Note that these classifiers are very much coupled to the grammar and must be
 * used with a clear understanding of the grammar. This is a very direct
 * violation of encapsulation, but I think it's fine in this case.
 */

func isExpression(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "fn":
		return true
	case "do", "panic":
		return true
	case "def", "val", "let", ".":
		return true
	case "if", "and", "or", "while":
		return true
	case "array", "array.read", "array.write", "array.for", "array.map", "array.filter", "array.reduce", "array.push", "array.pop", "array.shift", "array.unshift", "array.segment", "array.find", "array.splice", "array.reverse", "array.sort":
		return true
	case "map", "map.merge", "map.delete", "map.keys", "map.values", "map.read", "map.write", "map.extract":
		return true
	case "split", "substring", "concat", "chars":
		return true
	case "std.write", "http":
		return true
	case "GADFLY", "DAEMON", "GHOST", "ORACLE", "THEORY", "MUSE", "RAPTURE", "@":
		return true
	case "Array", "Number", "Hash", "String", "Function", "Identity":
		return true
	default:
		return false
	}
}

func isEnd(lexeme types.Lexeme) bool {
	return lexeme.Text == "end"
}

func isTrue(lexeme types.Lexeme) bool {
	return lexeme.Text == "true"
}

func isFalse(lexeme types.Lexeme) bool {
	return lexeme.Text == "false"
}

func isNil(lexeme types.Lexeme) bool {
	return lexeme.Text == "nil"
}

func isIdentifier(lexeme types.Lexeme) bool {
	switch string(lexeme.Text[0]) {
	case ".", "_", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
		return true
	default:
		return false
	}
}

func isSchema(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "Array", "Number", "Hash", "String", "Function", "Identity":
		return true
	default:
		return false
	}
}

func isPipe(lexeme types.Lexeme) bool {
	return lexeme.Text == "|"
}

func isColon(lexeme types.Lexeme) bool {
	return lexeme.Text == ":"
}

func isArrow(lexeme types.Lexeme) bool {
	return lexeme.Text == "->"
}

func isLogical(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "&&", "||":
		return true
	default:
		return false
	}
}

func isEquality(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "==", "!=":
		return true
	default:
		return false
	}
}

func isComparison(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "<", "<=", ">", ">=":
		return true
	default:
		return false
	}
}

func isTerm(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "+", "-":
		return true
	default:
		return false
	}
}

func isFactor(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "*", "/":
		return true
	default:
		return false
	}
}

func isUnary(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "-", "!":
		return true
	default:
		return false
	}
}

func isString(lexeme types.Lexeme) bool {
	return string(lexeme.Text[0]) == "\""
}

func isNumber(lexeme types.Lexeme) bool {
	switch string(lexeme.Text[0]) {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	default:
		return false
	}
}

// Is an atom anything that is not more specific?
func isAtom(lexeme types.Lexeme) bool {
	if isExpression(lexeme) {
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
