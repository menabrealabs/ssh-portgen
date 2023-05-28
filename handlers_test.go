package main

import (
	"crypto/sha256"
	"reflect"
	"testing"
)

func TestMain_getDigest_should_return_valid_sha256_hash(t *testing.T) {
	hostname := "UncleFester"
	want, err := getDigest(hostname)

	if err != nil {
		t.Errorf("Hostname '%s' threw error: %s", hostname, err)
	}

	// Get expected value to compare with what the function returns.
	b := sha256.New()
	b.Write([]byte(hostname))
	have := b.Sum(nil)

	if !reflect.DeepEqual(have, want) {
		t.Errorf("expected %s, but got %s", want, have)
	}
}

func TestMain_getPort_should_return_correct_port_number(t *testing.T) {
	const want = 1659

	digest, _ := getDigest("UncleFester")
	have := getPort(digest, digestIndex{2, 20})

	if have != want {
		t.Errorf("expected %d, but got %d", want, have)
	}
}
