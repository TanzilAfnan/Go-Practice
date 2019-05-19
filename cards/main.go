package main

func main() {

	// cards := newDeck()

	// cards.print()

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}
