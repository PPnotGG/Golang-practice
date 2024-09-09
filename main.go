package main

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// cards = remainingCards

	cards.shuffle()
	cards.print()
	// hand.print()

	// fmt.Println([]byte(cards.toString()))

	// cards.saveToFile("my_cards")
}
