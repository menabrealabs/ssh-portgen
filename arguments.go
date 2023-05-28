package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type digestIndex [2]uint8

const (
	byteSize   = 8
	digestSize = 256
	numBytes   = digestSize / byteSize
	numBase    = 10
)

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (b *digestIndex) String() string {
	return fmt.Sprintf("%d/%d", b[0], b[1])
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
func (b *digestIndex) Set(value string) error {
	numString := strings.Split(value, "/")

	if len(numString) < 2 {
		return errors.New(
			"Format Error: indices parameter must be a string of two integers separated by a slash: e.g. '2/20'")
	}

	for i, num := range numString {
		nInt, error := strconv.ParseUint(num, numBase, byteSize)
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
