/* here you'll find the code to create and manipulate a deck of cards*/

package main

import "fmt"

func main() {
	/* 	cards := newDeck()

	   	mano, restoMazo := deal(cards, 5)

	   	mano.print()
	   	restoMazo.print()

	*/
	cards := newDeck()

	fmt.Println(cards.deckToString())
}
