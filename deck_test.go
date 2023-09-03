package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected Length of Deck is 16 but got %v ", len(d))
	}
	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first cards of Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected last cards of four of Spades but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {

	// os.Remove("_decktesting")

	deck := newDeck()

	deck.toSaveToFile("decktesting")

	loadedDeck := getDeckfromFile("decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck got %v ", len(loadedDeck))
	}

	os.Remove("decktesting")

}
