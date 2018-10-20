package main

import "fmt"

func main() {
	cards := deck{"Ace", newCard()}
	cards = append(cards, "Club")

	fmt.Println(cards)
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
