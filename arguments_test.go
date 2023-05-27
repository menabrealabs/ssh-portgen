package main

import (
	"testing"
)

func TestArgs_should_set_var_to_default_value(t *testing.T) {
	var indices digestIndices
	err := indices.Set("2/20")
	if err != nil {
		t.Error(err)
	}
}

func TestArgs_one_value_should_fail(t *testing.T) {
	var indices digestIndices
	err := indices.Set("2")
	if err == nil {
		t.Errorf("Expected error on single value but got %s", indices.String())
	}
}

func TestArgs_above_upper_bound_value_should_fail(t *testing.T) {
	var indices digestIndices
	err := indices.Set("2/33")
	if err == nil {
		t.Errorf("Expected error when asking for index out of upper bounds but got %s", indices.String())
	}
}

func TestArgs_above_lower_bound_value_should_fail(t *testing.T) {
	var indices digestIndices
	err := indices.Set("-1/33")
	if err == nil {
		t.Errorf("Expected error when asking for index out of lower bounds but got %s", indices.String())
	}
}
