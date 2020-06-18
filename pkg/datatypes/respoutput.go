package datatypes

import "fmt"

// RESPOutput datatype is used to relay non-specific text back to the client (confirmations of actions, such as '+OK')
type RESPOutput struct {
	value string
}

// EncodeRESPOutput encodes a string value into a RESPOutput
func EncodeRESPOutput(value string) RESPOutput {
	return RESPOutput{value: value}
}

// Output returns a string according to RESP format that can be sent to the client
func (o RESPOutput) Output() string {
	return fmt.Sprintf("+%s\r\n", o.value)
}
