package main

import (
	"testing"
)

func TestMainOneOne(t *testing.T) {
	if one_one() != 3448043 {
		t.Fatalf("Does not equal %d", 3448043)
	}
}

func TestMainOneTwo(t *testing.T) {
	result := one_two()
	if result == 0 {
		t.Fatal("No result")
	} 
}
