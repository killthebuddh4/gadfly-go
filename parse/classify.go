package parse

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

/*
 * Note that these classifiers are very much coupled to the grammar and must be
 * used with a clear understanding of the grammar. This is a very direct
 * violation of encapsulation, but I think it's fine in this case.
 */

func arityForOperator(lexeme types.Lexeme) (int, error) {
	switch lexeme.Text {
	case "fn":
		return 1, nil
	case "do":
		return -1, nil
	case "panic":
		return 1, nil
	case "def":
		return 2, nil
	case "val":
		return 2, nil
	case "let":
		return 2, nil
	case ".":
		return -1, nil
	case "if":
		return 3, nil
	case "and":
		return -1, nil
	case "or":
		return -1, nil
	case "while":
		return 2, nil
	case "when":
		return 2, nil
	case "array":
		return -1, nil
	case "array.read":
		return 2, nil
	case "array.write":
		return 3, nil
	case "array.for":
		return 2, nil
	case "array.map":
		return 2, nil
	case "array.filter":
		return 2, nil
	case "array.reduce":
		return 2, nil
	case "array.push":
		return 2, nil
	case "array.pop":
		return 1, nil
	case "array.shift":
		return 1, nil
	case "array.unshift":
		return 2, nil
	case "array.segment":
		return 3, nil
	case "array.find":
		return 2, nil
	case "array.splice":
		return 3, nil
	case "array.reverse":
		return 1, nil
	case "array.sort":
		return 2, nil
	case "map":
		return -1, nil
	case "map.merge":
		return 2, nil
	case "map.delete":
		return 2, nil
	case "map.keys":
		return 1, nil
	case "map.values":
		return 1, nil
	case "map.read":
		return 2, nil
	case "map.write":
		return 3, nil
	case "map.extract":
		return 2, nil
	case "split":
		return 2, nil
	case "substring":
		return 3, nil
	case "concat":
		return 2, nil
	case "chars":
		return 1, nil
	case "std.write":
		return 1, nil
	case "http":
		return 1, nil
	case "GADFLY":
		return 1, nil
	case "DAEMON":
		return 1, nil
	case "GHOST":
		return 1, nil
	case "ORACLE":
		return 1, nil
	case "THEORY":
		return 1, nil
	case "MUSE":
		return 1, nil
	case "RAPTURE":
		return 1, nil
	case "@":
		return 1, nil
	case "signal":
		return 2, nil
	case "emit":
		return -1, nil
	case "on":
		return 2, nil
	case "catch":
		return 2, nil
	case "throw":
		return 1, nil
	default:
		return 0, errors.New("unknown operator, could not get arity " + lexeme.Text)
	}
}

func IsThunk(lexeme types.Lexeme, index int) bool {
	switch lexeme.Text {
	case "if", "and", "or", "while", "when":
		if (lexeme.Text == "if" || lexeme.Text == "when") && index > 0 {
			return true
		}

		if lexeme.Text == "and" || lexeme.Text == "or" {
			return true
		}

		if lexeme.Text == "while" {
			return true
		}

		return false
	default:
		return false
	}
}

func isExpression(lexeme types.Lexeme) bool {
	switch lexeme.Text {
	case "fn":
		return true
	case "do", "panic":
		return true
	case "def", "val", "let", ".":
		return true
	case "if", "and", "or", "while", "when":
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
	case "signal", "emit", "on", "catch", "throw":
		return true
	default:
		return false
	}
}

func isThen(lexeme types.Lexeme) bool {
	return lexeme.Text == "then"
}

func isElse(lexeme types.Lexeme) bool {
	return lexeme.Text == "else"
}

func isCatch(lexeme types.Lexeme) bool {
	return lexeme.Text == "catch"
}

func isValue(lexeme types.Lexeme) bool {
	return lexeme.Text == "value"
}

func isIdentifier(lexeme types.Lexeme) bool {
	switch string(lexeme.Text[0]) {
	case "_", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
		return true
	case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
		return true
	default:
		return false
	}
}

func isSchema(lexeme types.Lexeme) bool {
	switch string(lexeme.Text[0]) {
	case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
		if len(lexeme.Text) == 1 {
			return true
		} else {
			switch string(lexeme.Text[1]) {
			case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
				return true
			default:
				return false
			}
		}
	default:
		return false
	}
}

func isConstant(lexeme types.Lexeme) bool {
	switch string(lexeme.Text[0]) {
	case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
		switch string(lexeme.Text[1]) {
		case "_", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
			return true
		default:
			return false
		}
	default:
		return false
	}
}

func isSignature(lexeme types.Lexeme) bool {
	return lexeme.Text == "("
}

func isReturn(lexeme types.Lexeme) bool {
	return lexeme.Text == "(->"
}

func isEndSignature(lexeme types.Lexeme) bool {
	return lexeme.Text == ")"
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

func isPipe(lexeme types.Lexeme) bool {
	return lexeme.Text == "|"
}

func isColon(lexeme types.Lexeme) bool {
	return lexeme.Text == ":"
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
