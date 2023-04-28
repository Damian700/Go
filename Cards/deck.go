/* You'll find bellow the code that describes what a deck is and
how it works
New type of 'Deck' which is a slice of strings
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func deal(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

func (d deck) deckToString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// func WriteFile(filename string, data []byte, perm fs.FileMode) error
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	// func Split(s, sep string) []string
	return deck(ss)
}
