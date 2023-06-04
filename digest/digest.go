package digest

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	ByteSize = 2
)

type Index [ByteSize]byte

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (b *Index) String() string {
	return fmt.Sprintf("%d/%d", b[0], b[1])
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
func (b *Index) Set(value string) error {
	numString := strings.Split(value, "/")

	if len(numString) != ByteSize {
		return errors.New(
			"index parameter must be a string of two integers separated by a forward slash: e.g. '2/20'")
	}

	for i, num := range numString {
		nInt, error := strconv.ParseUint(num, 10, sha256.Size)
		if error != nil {
			return error
		}

		// Bounds of the Sha256 digest byte slice is 0..31
		if nInt >= sha256.Size {
			return errors.New("index %d is out of range")
		}

		b[i] = uint8(nInt)
	}

	return nil
}
