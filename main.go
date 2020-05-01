package main

func main() {
	//saves to hdd
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	//reads from hdd
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	//shuffle cards
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
