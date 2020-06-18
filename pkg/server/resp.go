package server

import (
	"bufio"
	"fmt"
	"strconv"
)

func getRESPCommand(scanner *bufio.Scanner) ([]string, error) {
	/*
		RESP commands are received as an 'Array' of 'Bulk String'

		Example format (Command without arguments): *1\r\n$7\r\nCOMMAND\r\n
		*1 = array of length 1
		$7 = string of length 7
		COMMAND = actual command

		Limitation of implementation:
		Only supporting SimpleString GET & SET commands, hence it will throw an error if the first byte doesn't equal '*'

		Proper implementation:
		Recursively parse the input to allow multiple commands, multi-value commands, different data types, etc..
		Store parsed objects into correct data types (ie: an RESPArray containing RESPBulkString)
	*/

	// Grab string from scanner
	line := scanner.Text()

	// Ensure we are parsing an array of bulk strings
	if line[0] == '*' {
		// Decode length of array expected in remainder of input
		arrLen, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			return nil, err
		}

		// Prep response string slice (not make()-ing the length fixed so we can compare resulting slice length to 'arrLen')
		inputArgs := []string{}

		// Process each bulk string
		for i := 1; i <= arrLen; i++ {
			// Attempt to advance to next line of input
			if ok := scanner.Scan(); !ok {
				return nil, fmt.Errorf("Error parsing inbound line")
			}

			// Grab string from scanner
			line := scanner.Text()

			// Check for $ = string
			if line[0] == '$' {
				// Decode length of string expected in next line of input
				strLength, err := strconv.Atoi(string(line[1:]))
				if err != nil {
					return nil, err
				}

				// Attempt to advance to next line of input
				if ok := scanner.Scan(); !ok {
					return nil, fmt.Errorf("Error parsing inbound line")
				}

				// Grab string from scanner
				line := scanner.Text()

				// Ensure length corresponds
				if len(line) == strLength {
					inputArgs = append(inputArgs, line)
				} else {
					return nil, fmt.Errorf("Length of decoded RESP string did not match expected length")
				}
			} else {
				// Unsupported RESP format
				return nil, fmt.Errorf("Unsupported RESP input received, only Arrays of Bulk Strings are supported")
			}
		}

		// Succesfully parsed an array of bulk strings, return it
		return inputArgs, nil
	}

	// Unsupported RESP format
	return nil, fmt.Errorf("Unsupported RESP input received, only Arrays of Bulk Strings are supported")
}
