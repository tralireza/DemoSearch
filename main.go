package main

import (
	"flag"
	"fmt"

	"github.com/tralireza/search"
)

func main() {
	var m, n int
	flag.IntVar(&m, "m", 10, "Number of rows in the maze")
	flag.IntVar(&n, "n", 48, "Number of columns in the maze")

	var i, j int
	flag.IntVar(&i, "i", 5, "Ghosty's start row, between 2 & m-1")
	flag.IntVar(&j, "j", 26, "Ghosty's start column, between 2 & n-1")

	var doors, blocks int
	flag.IntVar(&doors, "exits", 16, "Number of doors to get out of maze")
	flag.IntVar(&blocks, "walls", 128, "Extra bricks inside of maze")

	var doBFS, drawGrid bool
	flag.BoolVar(&doBFS, "BFS", false, "Do a BFS search (otherwise DFS)")
	flag.BoolVar(&drawGrid, "drawMaze", false, "Only draw the maze & exit!")

	flag.Parse()

	d := search.NewDemo(m, n)
	d.AddBlock(blocks)
	d.AddDoor(doors)
	s := search.Point{i - 1, j - 1}
	if drawGrid {
		d.SetStart(s)
		d.Draw()
		d.Stat(0)
		fmt.Print("\n")
		return
	}

	if doBFS {
		d.BFS(s)
	} else {
		d.DFS(s)
	}
}
