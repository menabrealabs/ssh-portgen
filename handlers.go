package main

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"os"
)

// Get the hostname from the command line or from the current environment.
func getHostname(arg string) (string, error) {
	if arg != "" {
		return arg, nil
	} else {
		// Default to getting the current system hostname
		return os.Hostname()
	}
}

// Get a Sha256 digest of the hostname.
func getDigest(hostname string) ([]byte, error) {
	b := sha256.New()
	_, err := b.Write([]byte(hostname))
	if err != nil {
		return nil, errors.New("Cannot generate SHA-256 digest")
	}
	return b.Sum(nil), nil
}

// Convert digest bytes into unsigned int and remove last digit.
func getPort(digest []byte, idx digestIndex) uint16 {
	digestBytes := []byte{digest[idx[0]], digest[idx[1]]}
	return binary.LittleEndian.Uint16(digestBytes) / 10
}
