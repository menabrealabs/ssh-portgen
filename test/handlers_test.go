package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"menabrealabs.com/ssh-portgen/digest"
	"menabrealabs.com/ssh-portgen/handlers"
)

func TestMain_getDigest_should_return_valid_sha256_hash(t *testing.T) {
	hostname := "UncleFester"
	want, err := handlers.GetDigest(hostname)

	if err != nil {
		t.Errorf("Hostname '%s' threw error: %s", hostname, err)
	}

	// Get expected value to compare with what the function returns.
	have := sha256.Sum256([]byte(hostname))

	if hex.EncodeToString(have[:]) != hex.EncodeToString(want) {
		t.Errorf("expected %x, but got %x", want, have)
	}
}

func TestMain_getDigest_should_fail_when_passed_empty_string(t *testing.T) {
	_, err := handlers.GetDigest("")

	if err == nil {
		t.Error("error not returned on an empty hostname argument")
	}
}

func TestMain_getPort_should_return_correct_port_number(t *testing.T) {
	const want = 16592

	d, _ := handlers.GetDigest("UncleFester")
	have, _ := handlers.GetPort(d, digest.Index{2, 20}, false)

	if have != want {
		t.Errorf("expected %d, but got %d", want, have)
	}
}

func TestMain_getPort_should_fail_when_passed_digest_with_insufficient_bytes(t *testing.T) {
	d, _ := handlers.GetDigest("UncleFester")
	malformedDigest := d[:31]
	_, err := handlers.GetPort(malformedDigest, digest.Index{2, 20}, false)

	if err == nil {
		t.Error("failed to raise an error on malformed digest")
	}
}

func TestMain_getPort_should_give_high_port_when_high_flag_set(t *testing.T) {
	const want = 16592

	d, _ := handlers.GetDigest("UncleFester")
	have, _ := handlers.GetPort(d, digest.Index{2, 20}, true)

	if want == have {
		t.Errorf("expected %d, but got %d", want, have)
	}
}
