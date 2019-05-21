package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {

	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card \"Spades of Ace\" but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs of King" {
		t.Errorf("Expected first card \"Clubs of King\" but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndnewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Deck length not as expected")
	}

	os.Remove("_decktesting")
}
