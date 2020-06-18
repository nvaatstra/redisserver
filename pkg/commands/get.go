package commands

import (
	"github.com/nvaatstra/redisserver/pkg/datatypes"
)

// Get command
type Get struct {
	key string
}

// NewGet returns a Get command if arguments are valid, an error if arguments are invalid
func NewGet(args []string) (Command, datatypes.Datatype) {
	cmd := Get{}

	// Get accepts exactly 1 argument
	if len(args) != 1 || len(args[0]) == 0 {
		return cmd, datatypes.NewError("ERR", "GET only accepts 1 non-empty argument: GET [key]")
	}

	cmd.key = args[0]

	return cmd, nil
}

// Execute performs the Get action on the input 'data' map
func (cmd Get) Execute(data map[string]datatypes.Datatype) datatypes.Datatype {
	if value, exists := data[cmd.key]; exists {
		ss, valid := datatypes.DecodeRESPSimpleString(value)

		if valid {
			// Send value as response
			return ss
		}

		// Send WRONGTYPE response
		return datatypes.NewError("WRONGTYPE", "Operation against a key holding the wrong kind of value")
	}

	return datatypes.NewNull()
}
