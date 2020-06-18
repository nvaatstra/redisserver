package commands

import "github.com/nvaatstra/redisserver/pkg/datatypes"

// CmdRegistry contains the builder functions for each supported command
var CmdRegistry = map[string]CommandBuilder{
	"GET": NewGet,
	"SET": NewSet,
}

// CommandBuilder interface to create Commands
type CommandBuilder func([]string) (Command, datatypes.Datatype)

// Command interface for command implementations
type Command interface {
	Execute(map[string]datatypes.Datatype) datatypes.Datatype
}
