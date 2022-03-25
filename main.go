package main

func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()
	// // cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards2 := newDeckFromFile("my_cards")
	// fmt.Println(cards2)

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
