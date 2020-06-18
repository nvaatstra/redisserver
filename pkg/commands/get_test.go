package commands

import (
	"testing"

	"github.com/nvaatstra/redisserver/pkg/datatypes"
)

func TestGet(t *testing.T) {
	// Prepare testing map
	dataset := map[string]datatypes.Datatype{}
	respSimpleString := datatypes.EncodeRESPSimpleString("testvalue")

	dataset["testkey"] = respSimpleString

	// Run tests with valid input & data
	cmdPositive, err := NewGet([]string{"testkey"})
	if err != nil {
		t.Errorf("NewGet with valid input returned an error: %s", err.Output())
	}

	outputPositive := cmdPositive.Execute(dataset)
	if outputPositive != respSimpleString {
		t.Error("Retrieved object does not match original object")
	}

	// Run test with invalid input
	_, err = NewGet([]string{""})
	if err == nil {
		t.Error("NewGet with empty key was accepted")
	}

	_, err = NewGet([]string{})
	if err == nil {
		t.Error("NewGet with empty string slice was accepted")
	}

	_, err = NewGet(nil)
	if err == nil {
		t.Error("NewGet with nil string slice was accepted")
	}

	_, err = NewGet([]string{"arg1", "arg2"})
	if err == nil {
		t.Error("NewGet with too many arguments was accepted")
	}
}
