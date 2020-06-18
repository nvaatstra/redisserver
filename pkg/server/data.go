package server

import (
	"github.com/nvaatstra/redisserver/pkg/commands"
	"github.com/nvaatstra/redisserver/pkg/datatypes"

	log "github.com/sirupsen/logrus"
)

type dcOp struct {
	cmd        commands.Command
	outputChan chan<- datatypes.Datatype
}

func dataController() chan<- dcOp {
	// Initialize the data structure
	// data := make(map[string]interface{})
	data := make(map[string]datatypes.Datatype)

	// Open input channel to receive ops
	inputChannel := make(chan dcOp)

	// Run processor in a go routine
	go func() {
		for {
			in, ok := <-inputChannel
			if ok {
				// Debugging
				log.Debugf("dataController() command type: %T", in.cmd)
				log.Debugf("dataController() command data: %+v", in.cmd)

				// Process command
				result := in.cmd.Execute(data)

				// Debugging
				log.Debugf("dataController() command result: %s", result)

				// Send response
				in.outputChan <- result

				// Close response channel
				close(in.outputChan)
			}
		}
	}()

	// Pass the input channel back to the server
	return inputChannel
}
