package main

import "fmt"

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// cards = remainingCards

	// cards.print()
	// hand.print()

	fmt.Println([]byte(cards.toString()))
}
