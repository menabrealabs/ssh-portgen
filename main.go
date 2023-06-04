package main

import (
	"flag"
	"fmt"
	"log"

	"menabrealabs.com/ssh-portgen/digest"
	"menabrealabs.com/ssh-portgen/handlers"
)

var indexFlag = digest.Index{2, 20}
var rawFlag bool
var lowFlag bool

func init() {
	flag.Var(&indexFlag, "index", "index of two bytes from the digest in the range 0..31, separated by a forward-slash: e.g. 2/20")
	flag.Var(&indexFlag, "i", "shorthand for -index flag")
	flag.BoolVar(&rawFlag, "raw", false, "raw output: set to true to print only the raw generated port number")
	flag.BoolVar(&rawFlag, "r", false, "shorthand for -raw flag")
	flag.BoolVar(&lowFlag, "low", false, "allow high (5 digit) port numbers")
	flag.BoolVar(&lowFlag, "l", false, "shorthand for -high flag")
}

func main() {

	flag.Parse()

	hostname, _ := handlers.GetHostname(flag.Arg(0))

	digest, err := handlers.GetDigest(hostname)
	logFatal(err)

	port, err := handlers.GetPort(digest, indexFlag, lowFlag)
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
