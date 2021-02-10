package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 but got %v", len(d))
	}
}
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedFile := newDeckFromFile("_decktesting")

	if len(loadedFile) != 16 {
		t.Errorf("Expected 16 but got %v", len(loadedFile))
	}
	os.Remove("_decktesting")
}
