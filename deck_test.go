package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 48 {
		t.Errorf("expected deck length of 48, got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck, _ := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 48 {
		t.Errorf("Expected 48 cards, got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestPrint(t *testing.T) {
	var in string
	d := newDeck()
	_, err := fmt.Scanln(&in)
	d.print()
	if in != d.toString() {
		t.Error(err)
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()

	initialD := d[0]

	d.shuffle()

	if initialD == string(d[0]) {
		t.Errorf("Expected initialD: %s to be different from %s", initialD, d)
	}
}
