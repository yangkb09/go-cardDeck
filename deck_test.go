package main

import "testing"

func TestNewDeck(t *testing.T) { // *testing.T is type; t is the test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d)) // %v will be taken from len(d)
	}
}
