package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	idxBytes = 2
	intBase  = 10
)

type digestIndex [idxBytes]uint8

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (b *digestIndex) String() string {
	return fmt.Sprintf("%d/%d", b[0], b[1])
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
func (b *digestIndex) Set(value string) error {
	numString := strings.Split(value, "/")

	if len(numString) != idxBytes {
		return errors.New(
			"indices parameter must be a string of two integers separated by a forward slash: e.g. '2/20'")
	}

	for i, num := range numString {
		nInt, error := strconv.ParseUint(num, intBase, byteSize)
		if error != nil {
			return error
		}

		// Bounds of the Sha256 digest byte slice is 0..31
		if nInt >= digestByteSize {
			return errors.New("index %d is out of range")
		}

		b[i] = uint8(nInt)
	}

	return nil
}
