package main

func main() {

	// cards := newDeck()

	// cards.print()

	// cards := newDeck()

	// fmt.Println(cards.toString())

	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("My_File")
	// hand, remainingCards := deal(cards, 5)

	cards := newDeckFromFile("My_File")
	cards.print()
	cards.shuffle()
	cards.print()

}
