package commands

import (
	"testing"

	"github.com/nvaatstra/redisserver/pkg/datatypes"
)

func TestSet(t *testing.T) {
	// Prepare testing map
	dataset := map[string]datatypes.Datatype{}
	testKey := "testkey"
	testValue := "testvalue"
	respSimpleString := datatypes.EncodeRESPSimpleString(testValue)

	// Run tests with valid input & data
	cmdPositive, err := NewSet([]string{testKey, testValue})
	if err != nil {
		t.Errorf("NewSet with valid input returned an error: %s", err.Output())
	}

	outputPositive := cmdPositive.Execute(dataset)
	if outputPositive != datatypes.EncodeRESPOutput("OK") {
		t.Error("Output of execute does not match expected output of succesful Set command")
	}

	if dataset[testKey] != respSimpleString {
		t.Error("Stored object does not match expected object")
	}

	// Run test with invalid input
	_, err = NewSet([]string{"", "value"})
	if err == nil {
		t.Error("NewSet with empty key was accepted")
	}

	_, err = NewSet([]string{"key"})
	if err == nil {
		t.Error("NewSet with omitted value was accepted")
	}

	_, err = NewSet([]string{})
	if err == nil {
		t.Error("NewSet with empty string slice was accepted")
	}

	_, err = NewSet(nil)
	if err == nil {
		t.Error("NewSet with nil string slice was accepted")
	}

	_, err = NewSet([]string{"arg1", "arg2", "arg2"})
	if err == nil {
		t.Error("NewSet with too many arguments was accepted")
	}
}
