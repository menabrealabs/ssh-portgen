package main

import (
	"testing"
)

func TestMain_should_set_var_to_default_value(t *testing.T) {
	var index digestIndex
	err := index.Set("2/20")
	if err != nil {
		t.Error(err)
	}
}

func TestMain_one_value_should_fail(t *testing.T) {
	var index digestIndex
	err := index.Set("2")
	if err == nil {
		t.Errorf("Expected error on single value but got %s", index.String())
	}
}

func TestMain_above_upper_bound_value_should_fail(t *testing.T) {
	var index digestIndex
	err := index.Set("2/33")
	if err == nil {
		t.Errorf("Expected error when asking for index out of upper bounds but got %s", index.String())
	}
}

func TestMain_above_lower_bound_value_should_fail(t *testing.T) {
	var index digestIndex
	err := index.Set("-1/33")
	if err == nil {
		t.Errorf("Expected error when asking for index out of lower bounds but got %s", index.String())
	}
}
