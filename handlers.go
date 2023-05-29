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
		hostname, err := os.Hostname()
		if err != nil {
			return "", err
		}
		return hostname, nil
	}
}

// Get a Sha256 digest of the hostname.
func getDigest(hostname string) ([]byte, error) {
	if len(hostname) < 1 {
		return nil, errors.New("hostname is empty")
	}

	b := sha256.New()
	_, err := b.Write([]byte(hostname))
	if err != nil {
		return nil, errors.New("failed to generate SHA-256 digest")
	}
	return b.Sum(nil), nil
}

// Convert digest bytes into unsigned int and remove last digit.
func getPort(digest []byte, idx digestIndex, low bool) (uint16, error) {
	if len(digest) < digestByteSize {
		return 0, errors.New("the digest value is malformed with an insufficient number of bytes")
	}

	if idx[0] >= digestByteSize || idx[1] >= digestByteSize {
		return 0, errors.New("digest index is out of range")
	}

	digestBytes := []byte{digest[idx[0]], digest[idx[1]]}
	port := binary.LittleEndian.Uint16(digestBytes)
	if low {
		return port / 10, nil
	}
	return port, nil
}
