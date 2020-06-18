package datatypes

import "fmt"

// RESPError datatype is used to relay errors back to the client
type RESPError struct {
	class   string
	message string
}

// NewError can be used to create a new RESPError
func NewError(class string, message string) RESPError {
	e := RESPError{
		class:   class,
		message: message,
	}

	return e
}

// Output returns a string according to RESP format that can be sent to the client
func (e RESPError) Output() string {
	return fmt.Sprintf("-%s %s\r\n", e.class, e.message)
}
