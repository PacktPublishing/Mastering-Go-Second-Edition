package main

import (
	"testing"
)

func Test_test_one(t *testing.T) {
	sleep_with_me()
	value := get_one()
	if value != 1 {
		t.Errorf("Function returned %v", value)
	}
	sleep_with_me()
}

func Test_test_two(t *testing.T) {
	sleep_with_me()
	value := get_two()
	if value != 2 {
		t.Errorf("Function returned %v", value)
	}
}

func Test_that_will_fail(t *testing.T) {
	value := get_one()
	if value != 2 {
		t.Errorf("Function returned %v", value)
	}
}
