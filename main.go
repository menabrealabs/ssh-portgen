package main

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	idxTopByte    = 20 // index of the digest byte array for the top value.
	idxBottomByte = 2  // index of the digest byte array for the bottom value.
)

func main() {
	hostname, err := getHostname()
	if err != nil {
		log.Fatalf("Cannot read hostname: \n%s\n", usage())
	}

	digest, err := getDigest(hostname)
	if err != nil {
		log.Fatal(err)
	}

	port := getPort(digest)

	// TODO: Check whether the port number is bound on current host

	// Print the output
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("SHA2 Digest: %x\n", digest)
	fmt.Printf("SSH port number: %d\n", port)
}

func usage() string {
	return "Usage: ssh-portgen [hostname]"
}

// Get the hostname from the command line or from the environment.
func getHostname() (string, error) {
	if len(os.Args) > 1 {
		return os.Args[1], nil
	} else {
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
func getPort(digest []byte) uint16 {
	digestBytes := []byte{digest[idxBottomByte], digest[idxTopByte]}
	port := binary.LittleEndian.Uint16(digestBytes) / 10
	return port
}
