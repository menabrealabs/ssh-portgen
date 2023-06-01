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

func TestMain_getDigest_should_fail_when_passed_empty_string(t *testing.T) {
	_, err := getDigest("")

	if err == nil {
		t.Error("error not returned on an empty hostname argument")
	}
}

func TestMain_getPort_should_return_correct_port_number(t *testing.T) {
	const want = 16592

	digest, _ := getDigest("UncleFester")
	have, _ := getPort(digest, digestIndex{2, 20}, false)

	if have != want {
		t.Errorf("expected %d, but got %d", want, have)
	}
}

func TestMain_getPort_should_fail_when_passed_digest_with_insufficient_bytes(t *testing.T) {
	digest, _ := getDigest("UncleFester")
	malformedDigest := digest[:31]
	_, err := getPort(malformedDigest, digestIndex{2, 20}, false)

	if err == nil {
		t.Error("failed to raise an error on malformed digest")
	}
}

func TestMain_getPort_should_give_high_port_when_high_flag_set(t *testing.T) {
	const want = 16592

	digest, _ := getDigest("UncleFester")
	have, _ := getPort(digest, digestIndex{2, 20}, true)

	if want == have {
		t.Errorf("expected %d, but got %d", want, have)
	}
}
