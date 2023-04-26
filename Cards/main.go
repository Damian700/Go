/* here you'll find the code to create and manipulate a deck of cards*/

package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}
