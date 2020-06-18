package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Server details
	network := "tcp"
	address := ":56123"

	// Prepare test server
	testServer := New(network, address)

	// Start server
	go testServer.Run()

	// Allow server to fully boot
	time.Sleep(2 * time.Second)

	// Run tests and ensure outcome is passed
	os.Exit(m.Run())
}

func TestServer(t *testing.T) {
	// Server details
	network := "tcp"
	address := ":56123"

	// Open a TCP connection
	conn, err := net.Dial(network, address)
	if err != nil {
		t.Errorf("Unable to connect to testServer: %s", err.Error())
	}
	defer conn.Close()

	// Grab scanner to read from the connection
	scanner := bufio.NewScanner(conn)

	// Prepare commands
	testKey := "testkey"
	testValue := "testvalue"
	cmdSet := fmt.Sprintf("*3\r\n$3\r\nSET\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(testKey), testKey, len(testValue), testValue)
	cmdGet := fmt.Sprintf("*2\r\n$3\r\nGET\r\n$%d\r\n%s\r\n", len(testKey), testKey)

	// Send Set
	_, err = conn.Write([]byte(cmdSet))
	if err != nil {
		t.Errorf("Unable to send Set command to testServer: %s", err.Error())
	}

	// Read Set reply
	if scanner.Scan() {
		line := scanner.Text()

		if strings.Compare(line, "+OK") != 0 {
			t.Errorf("Expected '+OK' as response to Set command, instead got: %s", line)
		}
	} else {
		t.Error("Expected a response to Set command, but did not receive a reply")
	}

	// Send Get
	_, err = conn.Write([]byte(cmdGet))
	if err != nil {
		t.Errorf("Unable to send Get command to testServer: %s", err.Error())
	}

	// Read Get reply
	if scanner.Scan() {
		line := scanner.Text()

		if strings.Compare(line, fmt.Sprintf("+%s", testValue)) != 0 {
			t.Errorf("Expected '+%s' as response to Get command, instead got: %s", testValue, line)
		}
	} else {
		t.Error("Expected a response to Get command, but did not receive a reply")
	}

}
