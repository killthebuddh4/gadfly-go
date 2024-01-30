package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval/ai"
	"github.com/killthebuddh4/gadflai/eval/array"
	"github.com/killthebuddh4/gadflai/eval/io"
	"github.com/killthebuddh4/gadflai/eval/lambda"
	"github.com/killthebuddh4/gadflai/eval/maps"
	"github.com/killthebuddh4/gadflai/eval/strings"
	"github.com/killthebuddh4/gadflai/types"
)

type D func(t *types.Trajectory, e types.Exec) (types.Value, error)

func dispatch(trajectory *types.Trajectory) (D, error) {
	switch trajectory.Expression.Operator.Type {
	case "program":
		return Do, nil
	case "!=":
		return BangEqual, nil
	case "==":
		return EqualEqual, nil
	case ">":
		return Greater, nil
	case ">=":
		return GreaterEqual, nil
	case "<":
		return Less, nil
	case "<=":
		return LessEqual, nil
	case "&&":
		return Conjunction, nil
	case "||":
		return Disjunction, nil
	case "+":
		return Plus, nil
	case "-":
		return Minus, nil
	case "/":
		return Multiply, nil
	case "*":
		return Divide, nil
	case "!":
		return Bang, nil
	case "true":
		return True, nil
	case "false":
		return False, nil
	case "nil":
		return Nil, nil
	case "number":
		return Number, nil
	case "string":
		return String, nil
	case "identifier":
		return Identifier, nil
	case "do":
		return Do, nil
	case "panic":
		return Panic, nil
	case "def":
		return Def, nil
	case "fn":
		return lambda.Lambda, nil
	case "@":
		return lambda.Call, nil
	case "let":
		return Let, nil
	case "while":
		return While, nil
	case "map":
		return maps.Map, nil
	case "map.merge":
		return maps.Merge, nil
	case "map.delete":
		return maps.Delete, nil
	case "map.extract":
		return maps.Extract, nil
	case "map.keys":
		return maps.Keys, nil
	case "map.values":
		return maps.Values, nil
	case "map.read":
		return maps.Read, nil
	case "map.write":
		return maps.Write, nil
	case "array":
		return array.Array, nil
	case "array.read":
		return array.Read, nil
	case "array.for":
		return array.For, nil
	case "array.filter":
		return array.Filter, nil
	case "array.map":
		return array.Map, nil
	case "array.reduce":
		return array.Reduce, nil
	case "array.push":
		return array.Push, nil
	case "array.pop":
		return array.Pop, nil
	case "array.write":
		return array.Write, nil
	case "array.shift":
		return array.Shift, nil
	case "array.unshift":
		return array.Unshift, nil
	case "array.segment":
		return array.Segment, nil
	case "array.find":
		return array.Find, nil
	case "array.splice":
		return array.Splice, nil
	case "array.reverse":
		return array.Reverse, nil
	case "array.sort":
		return array.Sort, nil
	case "if":
		return If, nil
	case "and":
		return And, nil
	case "or":
		return Or, nil
	case "io.puts":
		return io.Puts, nil
	case "io.http":
		return io.Http, nil
	case "chars":
		return Chars, nil
	case "split":
		return strings.Split, nil
	case "substring":
		return strings.Substring, nil
	case "concat":
		return strings.Concat, nil
	case "GADFLY":
		return ai.Gadfly, nil
	case "DAEMON":
		return ai.Daemon, nil
	case "GHOST":
		return ai.Ghost, nil
	case "ORACLE":
		return ai.Oracle, nil
	default:
		return nil, errors.New("error dispatching, unknown operator " + trajectory.Expression.Operator.Type)
	}
}
