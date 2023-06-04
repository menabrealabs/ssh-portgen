package handlers

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"os"

	"menabrealabs.com/ssh-portgen/digest"
)

// Get the hostname from the command line or from the current environment.
func GetHostname(arg string) (string, error) {
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
func GetDigest(hostname string) ([]byte, error) {
	if hostname == "" {
		return []byte{}, errors.New("hostname is empty")
	}

	sum := sha256.Sum256([]byte(hostname))
	return sum[:], nil
}

// Convert digest bytes into unsigned int and remove last digit.
func GetPort(digest []byte, idx digest.Index, low bool) (uint16, error) {
	if len(digest) < sha256.Size {
		return 0, errors.New("the digest value is malformed with an insufficient number of bytes")
	}

	if idx[0] >= sha256.Size || idx[1] >= sha256.Size {
		return 0, errors.New("digest index is out of range")
	}

	digestBytes := []byte{digest[idx[0]], digest[idx[1]]}
	port := binary.LittleEndian.Uint16(digestBytes)
	if low {
		return port / 10, nil
	}
	return port, nil
}
