package main

import (
	"math/rand"
	"time"
)

type cell struct {
	Column        int
	Row           int
	Ismine        bool
	IsExposed     bool
	touchingMines int
	Label         string
	Neighbors     []cell
}

type board struct {
	Columns [][]cell
	Cols    int
	Rows    int
}

func (b *board) createBoard(columns int, rows int, mines int) {
	b.Cols = columns
	b.Rows = rows
	b.Columns = make([][]cell, columns)
	// create the rows of cells
	for i := range b.Columns {
		b.Columns[i] = make([]cell, rows)
	}
	// initialize the cells
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			b.Columns[i][j] = cell{
				Column:    i,
				Row:       j,
				Ismine:    false,
				IsExposed: false,
				Label:     "",
			}
		}
	}
	// allocate mines to cells randomly
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < mines; {
		v := rand.Intn(columns)
		w := rand.Intn(rows)
		if b.Columns[v][w].Ismine == false {
			b.Columns[v][w].Ismine = true
			i++
		}
	}
	// for each cell, give it a list of neighbors
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			col := b.Columns[i][j].Column
			row := b.Columns[i][j].Row
			if col-1 >= 0 {
				if row-1 >= 0 {
					b.Columns[i][j].Neighbors = append(b.Columns[i][j].Neighbors, cell{Column: i - 1, Row: j - 1})
				}
				b.Columns[i][j].Neighbors = append(b.Columns[i][j].Neighbors, cell{Column: i - 1, Row: j})
				if row+1 < rows {
					b.Columns[i][j].Neighbors = append(b.Columns[i][j].Neighbors, cell{Column: i - 1, Row: j + 1})
				}
			}
			if col+1 < columns {
				if row-1 >= 0 {
					b.Columns[i][j].Neighbors = append(b.Columns[i][j].Neighbors, cell{Column: i + 1, Row: j - 1})
				}
				b.Columns[i][j].Neighbors = append(b.Columns[i][j].Neighbors, cell{Column: i + 1, Row: j})
				if row+1 < rows {
					b.Columns[i][j].Neighbors = append(b.Columns[i][j].Neighbors, cell{Column: i + 1, Row: j + 1})
				}
			}
			if row-1 >= 0 {
				b.Columns[i][j].Neighbors = append(b.Columns[i][j].Neighbors, cell{Column: i, Row: j - 1})
			}
			if row+1 < rows {
				b.Columns[i][j].Neighbors = append(b.Columns[i][j].Neighbors, cell{Column: i, Row: j + 1})
			}
		}
	}
	// for each cell, if isMine is false and the number of touching mines is 0, make a list of all cells that would be opened
	// if the cell is clicked on
	return
}
