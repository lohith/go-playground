package main

import "testing"

func TestSum(t *testing.T) {
	got := mySum(3, 4, 6)
	if got != 13 {
		t.Errorf("Expected 13, got %v", got)
	}

}
