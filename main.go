package main

func main() {
	cards := newDeckFromFile("my-cards")

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()

	cards.print()

}