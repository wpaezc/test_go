package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck lenght of 16, but got %d", len(d))
	}

	if d[0] != "One of Espadas" {
		t.Errorf("Bad first element %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck lenght of 16, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
