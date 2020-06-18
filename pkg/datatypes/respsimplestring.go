package datatypes

import "fmt"

// RESPSimpleString datatype is used to store string objects in the dataset
type RESPSimpleString struct {
	value string
}

// DecodeRESPSimpleString attempts to decode a generic Datatype into a RESPSimpleString
func DecodeRESPSimpleString(input Datatype) (RESPSimpleString, bool) {
	switch input.(type) {
	case RESPSimpleString:
		return input.(RESPSimpleString), true
	default:
		return RESPSimpleString{}, false
	}
}

// EncodeRESPSimpleString encodes a string value into a RESPSimpleString
func EncodeRESPSimpleString(value string) RESPSimpleString {
	return RESPSimpleString{value: value}
}

// Output returns a string according to RESP format that can be sent to the client
func (ss RESPSimpleString) Output() string {
	return fmt.Sprintf("+%s\r\n", ss.value)
}
