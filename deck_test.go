package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	length := 16

	if len(d) != length {
		t.Errorf("Expected deck length of %v, but got %v", length, len(d))
	}

	card := "Ace of Spades"
	actualCard := d[0]

	if actualCard != card {
		t.Errorf("Expected first card %v, but got %v",card, actualCard)
	}

	card = "Four of Clubs"
	actualCard = d[len(d) - 1]

	if actualCard != card {
		t.Errorf("Expected last card %v, but got %v", card, actualCard)
	}
}