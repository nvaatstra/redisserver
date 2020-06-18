package server

import (
	"bufio"
	"net"
	"os"

	"github.com/nvaatstra/redisserver/pkg/commands"
	"github.com/nvaatstra/redisserver/pkg/datatypes"

	log "github.com/sirupsen/logrus"
)

// Server object to handle an instance of the Redis-like server
type Server struct {
	network string
	address string

	cmdRegistry map[string]commands.CommandBuilder
}

// New instantiates a new instance of the Redis-like server
func New(network string, address string) *Server {
	// Initialize server
	s := &Server{
		network:     network,
		address:     address,
		cmdRegistry: commands.CmdRegistry,
	}

	// Return server
	return s
}

// Run starts the server
func (s *Server) Run() {
	// Run single data controller and obtain its input channel
	dc := dataController()

	// Attempt to start listening for inbound connections
	ln, err := net.Listen(s.network, s.address)
	if err != nil {
		log.Fatalf("Unable to start serving: %s", err.Error())
	}

	for {
		// Accept incoming connections
		conn, err := ln.Accept()
		if err != nil {
			log.Errorf("Unable to accept incoming connection: %s", err.Error())
		}

		// Handle connection in goroutine and pass it the channel to the data controller
		go s.handleConnection(conn, dc)
	}
}

// Kill stops the server (not gracefully - used for testcases)
func (s *Server) Kill() {
	os.Exit(0)
}

func (s *Server) handleConnection(conn net.Conn, dc chan<- dcOp) {
	// Grab scanner to read from the connection
	scanner := bufio.NewScanner(conn)

	// Scanner loop to allow interactive session
	for scanner.Scan() {
		// Grab RESP compatible command
		respCmd, err := getRESPCommand(scanner)
		if err != nil {
			conn.Write([]byte(datatypes.NewError("ERR", err.Error()).Output()))
			continue
		}

		// Prepare corresponding Command
		if cmdBuilder, ok := s.cmdRegistry[respCmd[0]]; ok {
			// Build the command
			cmd, respError := cmdBuilder(respCmd[1:])
			if err != nil {
				writeOutput(conn, respError)
				continue
			}

			// Open channel to receive output from controller
			outputChan := make(chan datatypes.Datatype)

			// Prepare op for the controller
			dcOp := dcOp{
				outputChan: outputChan,
				cmd:        cmd,
			}

			// Let controller process the op
			dc <- dcOp

			// Process controller output
			for output := range outputChan {
				writeOutput(conn, output)
			}
		} else {
			// Unsupported command
			writeOutput(conn, datatypes.NewError("ERR", "unknown command"))
		}
	}
}

func writeOutput(conn net.Conn, data datatypes.Datatype) {
	// Attempt to write output to the client
	_, err := conn.Write([]byte(data.Output()))
	if err != nil {
		log.Errorf("Error writing output to client: %s", err.Error())
	}

	// Todo: Above write() fails primarily on connection problems
	// Todo: Could try to force the connection to close here, but cannot inform the client since we apparently cannot write to it
}
