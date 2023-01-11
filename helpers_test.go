package main

import "testing"

func TestCreateBoard(t *testing.T) {
	board := board{}
	mines := 99
	board.createBoard(16, 24, mines)
	// count the number of mines
	counter := 0
	for _, c := range board.Columns {
		for _, r := range c {
			if r.Ismine {
				counter += 1
			}
		}
	}
	if counter != mines {
		t.Fatalf("there was not 99 mines present, there were %d", counter)
	}
}
