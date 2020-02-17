package main

var card string

func main() {
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
