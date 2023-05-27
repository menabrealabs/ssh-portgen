package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type digestIndices [2]uint8

const numBytes = 256 / 8

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (b *digestIndices) String() string {
	return fmt.Sprint(*b)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
func (b *digestIndices) Set(value string) error {
	numString := strings.Split(value, "/")

	if len(numString) < 2 {
		return errors.New(
			"Format Error: indices parameter must be a string of two integers separated by a slash: e.g. '2/20'")
	}

	for i, num := range numString {
		nInt, error := strconv.ParseUint(num, 10, 8)
		if error != nil {
			return error
		}

		// Bounds of the Sha256 digest byte slice is 0..31
		if nInt >= numBytes || nInt < 0 {
			return errors.New("Bounds Error: index %d is out of range")
		}

		b[i] = uint8(nInt)
	}

	return nil
}
