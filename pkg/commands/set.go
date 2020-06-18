package commands

import (
	"github.com/nvaatstra/redisserver/pkg/datatypes"
)

// Set command
type Set struct {
	key   string
	value datatypes.RESPSimpleString
}

// NewSet returns a Set command if arguments are valid, an error if arguments are invalid
func NewSet(args []string) (Command, datatypes.Datatype) {
	cmd := Set{}

	// Set accepts exactly 2 arguments
	if len(args) != 2 {
		return cmd, datatypes.NewError("ERR", "SET only accepts 2 arguments: SET [key] [value]")
	}

	// Set requires a valid key
	if len(args[0]) == 0 {
		return cmd, datatypes.NewError("ERR", "SET requires a non-empty key: SET [key] [value]")
	}

	cmd.key = args[0]
	cmd.value = datatypes.EncodeRESPSimpleString(args[1])

	return cmd, nil
}

// Execute performs the Set action on the input 'data' map
func (cmd Set) Execute(data map[string]datatypes.Datatype) datatypes.Datatype {
	data[cmd.key] = cmd.value

	return datatypes.EncodeRESPOutput("OK")
}
