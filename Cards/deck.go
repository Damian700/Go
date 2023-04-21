/* You'll find bellow the code that describes what a deck is and
how it works

New type of 'Deck' which is a slice of strings
*/

package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardPalo := []string{"Espadas", "Oro", "Copa", "Bastos"}
	cardValue := []string{"Ancho", "Dos", "Tres", "Cuatro"}
	for _, palo := range cardPalo {
		for _, value := range cardValue {
			cards = append(cards, value+" de "+palo)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
