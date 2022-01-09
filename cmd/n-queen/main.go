package main

import (
	"flag"
	"github.com/grugrut/n-queen/internal/bigbitboard"
	"github.com/grugrut/n-queen/internal/bitboard"
	"github.com/grugrut/n-queen/internal/naive"
)

func main() {
	solver := flag.String("solver", "naive", "solver")
	size := flag.Int("n", 8, "board size")

	flag.Parse()

	if *solver == "bitboard" {
		bitboard.Solve(*size)
	} else if *solver == "bigbitboard" {
		bigbitboard.Solve(*size)
	} else {
		naive.Solve(*size)
	}
}
