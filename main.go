package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	byteSize       = 0x0008                   // 8
	digestBitSize  = 0x0100                   // 256
	digestByteSize = digestBitSize / byteSize // 32
)

var indexFlag = digestIndex{2, 20}
var rawFlag bool
var lowFlag bool

func init() {
	flag.Var(&indexFlag, "indices", "index of two bytes from the digest in the range 0..31, separated by a forward-slash: e.g. 2/20")
	flag.Var(&indexFlag, "i", "shorthand for -indices flag")
	flag.BoolVar(&rawFlag, "raw", false, "raw output: set to true to print only the raw generated port number")
	flag.BoolVar(&rawFlag, "r", false, "shorthand for -raw flag")
	flag.BoolVar(&lowFlag, "low", false, "allow high (5 digit) port numbers")
	flag.BoolVar(&lowFlag, "l", false, "shorthand for -high flag")
}

func main() {

	flag.Parse()

	hostname, _ := getHostname(flag.Arg(0))

	digest, err := getDigest(hostname)
	logFatal(err)

	port, err := getPort(digest, indexFlag, lowFlag)
	logFatal(err)

	// Print the output
	if rawFlag {
		fmt.Print(port)
	} else {
		fmt.Printf("Hostname: %s\n", hostname)
		fmt.Printf("SHA2 Digest: %x\n", digest)
		fmt.Printf("SSH port number: %d\n", port)
	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
