package bitboard

import (
	"fmt"
)

var result int

func Solve(N int) {
	if N > 8 {
		fmt.Println("Now version, specify N <= 8")
		return
	}
	backtrack(0)
	fmt.Println(result)
}

func check() bool {
	return true
}

func backtrack(q uint64) {

}

func countBit(i uint64) int {
	i -= ((i >> 1) & 0x5555555555555555)
	i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f0f0f0f0f
	i += i >> 8
	i += i >> 16
	i += i >> 32
	return int(i & 0x7f)
}

func getPlacableCell(q uint64) uint64 {
	var b uint64 = 0

	var temp uint64
	var sheild uint64

	// left
	sheild = 0x7f7f7f7f7f7f7f7f
	temp = q
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	temp |= (temp & sheild) << 1
	b |= temp

	// up left
	sheild = 0x007f7f7f7f7f7f7f
	temp = q
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	temp |= (temp & sheild) << 9
	b |= temp

	// up
	sheild = 0x00ffffffffffffff
	temp = q
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	temp |= (temp & sheild) << 8
	b |= temp

	// up right
	sheild = 0x00fefefefefefefe
	temp = q
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	temp |= (temp & sheild) << 7
	b |= temp

	// right
	sheild = 0xfefefefefefefefe
	temp = q
	temp |= (temp & sheild) >> 1
	temp |= (temp & sheild) >> 1
	temp |= (temp & sheild) >> 1
	temp |= (temp & sheild) >> 1
	temp |= (temp & sheild) >> 1
	temp |= (temp & sheild) >> 1
	temp |= (temp & sheild) >> 1
	b |= temp

	// down right
	sheild = 0xfefefefefefefefe
	temp = q
	temp |= (temp & sheild) >> 9
	temp |= (temp & sheild) >> 9
	temp |= (temp & sheild) >> 9
	temp |= (temp & sheild) >> 9
	temp |= (temp & sheild) >> 9
	temp |= (temp & sheild) >> 9
	temp |= (temp & sheild) >> 9
	b |= temp

	// down
	sheild = 0xffffffffffffffff
	temp = q
	temp |= (temp & sheild) >> 8
	temp |= (temp & sheild) >> 8
	temp |= (temp & sheild) >> 8
	temp |= (temp & sheild) >> 8
	temp |= (temp & sheild) >> 8
	temp |= (temp & sheild) >> 8
	temp |= (temp & sheild) >> 8
	b |= temp

	// down left
	// down right
	sheild = 0x7f7f7f7f7f7f7f7f
	temp = q
	temp |= (temp & sheild) >> 7
	temp |= (temp & sheild) >> 7
	temp |= (temp & sheild) >> 7
	temp |= (temp & sheild) >> 7
	temp |= (temp & sheild) >> 7
	temp |= (temp & sheild) >> 7
	temp |= (temp & sheild) >> 7
	b |= temp

	b = ^b
	fmt.Printf("%064b\n", b)
	return b
}
