package main

import (
	"os"
	"testing"
)

/*
- The go test runner will run automatically the functions with arguments *testing.T
that are inside a "_test.go" file

- Take care of any cleaning you have to make after testing
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
	if d[0] != "Ancho de Espadas" {
		t.Errorf("Expected first card to be Ancho de Espadas, but got %v", d[0])
	}
	if d[len(d)-1] != "Cuatro de Bastos" {
		t.Errorf("Expected last card to be Cuatro de Bastos, but got %v", d[len(d)-1])
	}
}

// placeholders https://pkg.go.dev/fmt (%d, %v, etc)

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// first delete "_decktesting"
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expexted 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
