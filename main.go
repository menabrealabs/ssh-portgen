package main

import (
	"flag"
	"fmt"
	"log"
)

var indicesFlag = digestIndex{2, 20}
var rawFlag bool

func init() {
	flag.Var(&indicesFlag, "indices", "index of two bytes from the digest in the range 0..31, separated by a forward-slash: e.g. 2/20")
	flag.Var(&indicesFlag, "i", "shorthand for -indices flag")
	flag.BoolVar(&rawFlag, "raw", false, "raw output: set to true to print only the raw generated port number")
	flag.BoolVar(&rawFlag, "r", false, "shorthand for -raw flag")
}

func main() {

	flag.Parse()

	hostname, _ := getHostname(flag.Arg(0))

	digest, err := getDigest(hostname)
	if err != nil {
		log.Fatal(err)
	}

	port := getPort(digest, indicesFlag)

	// Print the output
	if rawFlag {
		fmt.Print(port)
	} else {
		fmt.Printf("Hostname: %s\n", hostname)
		fmt.Printf("SHA2 Digest: %x\n", digest)
		fmt.Printf("SSH port number: %d\n", port)
	}

}
