package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
)

const (
	idxTopByte    = 20 // index of the digest byte array for the top value.
	idxBottomByte = 2  // index of the digest byte array for the bottom value.
)

func main() {
	hostname := getHostname()
	digest := getDigest(hostname)
	port := getPort(digest)

	// Print the output
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("SHA256 Digest: %x\n", digest)
	fmt.Printf("SSH port number: %d\n", port)
}

// Get the hostname from the command line or from the environment.
func getHostname() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		return os.Getenv("HOSTNAME")
	}
}

// Get a Sha256 digest of the hostname.
func getDigest(hostname string) []byte {
	b := sha256.New()
	b.Write([]byte(hostname))
	return b.Sum(nil)
}

// Convert digest bytes into unsigned int and remove last digit.
func getPort(digest []byte) uint16 {
	digestBytes := []byte{digest[idxBottomByte], digest[idxTopByte]}
	port := binary.LittleEndian.Uint16(digestBytes) / 10
	return port
}
