package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var WriteHash types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	base, ok := arguments[0].(map[string]types.Value)

	if !ok {
		return nil, errors.New("WriteHash :: not a map")
	}

	key, ok := arguments[1].(string)

	if !ok {
		return nil, errors.New("Writehash :: not a string")
	}

	written := make(map[string]types.Value)

	for k, v := range base {
		written[k] = v
	}

	written[key] = arguments[2]

	return written, nil
}
