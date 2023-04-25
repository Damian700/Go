/* here you'll find the code to create and manipulate a deck of cards*/

package main

func main() {
	cards := newDeck()

	mano, restoMazo := deal(cards, 5)

	mano.print()
	restoMazo.print()
}
