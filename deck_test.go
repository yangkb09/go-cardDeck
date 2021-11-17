package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // *testing.T is type; t is the test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spaces, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	// make new deck, save to file
	deck := newDeck()
	deck.saveToFile("_decktesting")

	// load from disk
	loadedDeck := newDeckFromFile("_decktesting")

	// check if deck appropriately loaded
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
