package exec

import (
	"errors"

	"github.com/killthebuddh4/gadflai/array"
	"github.com/killthebuddh4/gadflai/eval"
	"github.com/killthebuddh4/gadflai/io"
	"github.com/killthebuddh4/gadflai/maps"
	"github.com/killthebuddh4/gadflai/strings"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

type D func(t *traj.Trajectory, e eval.Eval) (value.Value, error)

func dispatch(trajectory *traj.Trajectory) (D, error) {
	switch trajectory.Expression.Operator.Type {
	case "program":
		return eval.Do, nil
	case "!=":
		return eval.BangEqual, nil
	case "==":
		return eval.EqualEqual, nil
	case ">":
		return eval.Greater, nil
	case ">=":
		return eval.GreaterEqual, nil
	case "<":
		return eval.Less, nil
	case "<=":
		return eval.LessEqual, nil
	case "&&":
		return eval.Conjunction, nil
	case "||":
		return eval.Disjunction, nil
	case "+":
		return eval.Plus, nil
	case "-":
		return eval.Minus, nil
	case "/":
		return eval.Multiply, nil
	case "*":
		return eval.Divide, nil
	case "!":
		return eval.Bang, nil
	case "true":
		return eval.True, nil
	case "false":
		return eval.False, nil
	case "nil":
		return eval.Nil, nil
	case "number":
		return eval.Number, nil
	case "string":
		return eval.String, nil
	case "identifier":
		return eval.Identifier, nil
	case "do":
		return eval.Do, nil
	case "panic":
		return eval.Panic, nil
	case "def":
		return eval.Def, nil
	case "fn":
		return eval.Lambda, nil
	case "@":
		return eval.Call, nil
	case "let":
		return eval.Let, nil
	case "while":
		return eval.While, nil
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
		return eval.If, nil
	case "and":
		return eval.And, nil
	case "or":
		return eval.Or, nil
	case "io.puts":
		return io.Puts, nil
	case "io.http":
		return io.Http, nil
	case "chars":
		return eval.Chars, nil
	case "split":
		return strings.Split, nil
	case "substring":
		return strings.Substring, nil
	case "concat":
		return strings.Concat, nil
	default:
		return nil, errors.New("error dispatching, unknown operator " + trajectory.Expression.Operator.Type)
	}
}
