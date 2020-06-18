package datatypes

import "fmt"

// RESPNull datatype is used to relay non-existing keys (ie: (nil)) back to the client
type RESPNull struct {
	value string
}

// NewNull can be used to create a new RESPNull
func NewNull() RESPNull {
	n := RESPNull{
		value: "$-1",
	}

	return n
}

// Output returns a string according to RESP format that can be sent to the client
func (n RESPNull) Output() string {
	return fmt.Sprintf("%s\r\n", n.value)
}
