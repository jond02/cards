package main

import "fmt"

func main() {
	cards := []string{"Ace", newCard()}
	cards = append(cards, "Club")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
