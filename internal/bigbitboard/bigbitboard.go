package bigbitboard

import (
	"fmt"
	"math"
)

var result int

func Solve(N int) {
	if N > 16 {
		fmt.Println("Now version, specify N <= 16")
		return
	}
	var checked [4]uint64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			n := 4*i + j
			if n < N {
				checked[i] = (checked[i] << 16) | (uint64(math.Pow(2, float64(16-N)) - 1))
			} else {
				checked[i] = (checked[i] << 16) | 0xffff
			}
		}
	}
	var q [4]uint64
	backtrack(N, &q, &checked)
	fmt.Println(result)
}

func check() bool {
	return true
}

func backtrack(N int, q *[4]uint64, checked *[4]uint64) {
	/*
		for {
			b := (getPlacableCell(q) & ^checked)
			if b == 0 {
				if countBit(q) == N {
					fmt.Printf("%064b OK\n", q)
					result++
				}
				return
			}
			next := (b & -b)
			checked |= next
			backtrack(N, q|next, checked)
		}
	*/
}

func countBit(q *[4]uint64) int {
	result := 0
	for n := 0; n < 4; n++ {
		i := q[n]
		i -= ((i >> 1) & 0x5555555555555555)
		i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
		i = (i + (i >> 4)) & 0x0f0f0f0f0f0f0f0f
		i += i >> 8
		i += i >> 16
		i += i >> 32

		result += int(i & 0x7f)
	}
	return result
}

func getPlacableCell(q *[4]uint64) [4]uint64 {
	var b [4]uint64

	var temp [4]uint64
	var sheild uint64

	// left
	sheild = 0x7fff7fff7fff7fff
	temp[3] = q[3]
	temp[2] = q[2]
	temp[1] = q[1]
	temp[0] = q[0]
	for i := 0; i < 15; i++ {
		temp[3] |= (temp[3] & sheild) << 1
		temp[2] |= (temp[2] & sheild) << 1
		temp[1] |= (temp[1] & sheild) << 1
		temp[0] |= (temp[0] & sheild) << 1
	}
	b[3] |= temp[3]
	b[2] |= temp[2]
	b[1] |= temp[1]
	b[0] |= temp[0]

	// up left
	sheild = 0x00007fff7fff7fff
	temp[3] = q[3]
	temp[2] = q[2] |
		((0x0fff & q[3]) << 4) | ((0x1fff & (q[3] >> 16)) << 3) | ((0x3fff & (q[3] >> 32)) << 2) | ((0x7fff & (q[3] >> 48)) << 1)
	temp[1] = q[1] |
		((0x0fff & q[2]) << 4) | ((0x1fff & (q[2] >> 16)) << 3) | ((0x3fff & (q[2] >> 32)) << 2) | ((0x7fff & (q[2] >> 48)) << 1) |
		((0x00ff & q[3]) << 8) | ((0x01ff & (q[3] >> 16)) << 7) | ((0x03ff & (q[3] >> 32)) << 6) | ((0x07ff & (q[3] >> 48)) << 5)
	temp[0] = q[0] |
		((0x0fff & q[1]) << 4) | ((0x1fff & (q[1] >> 16)) << 3) | ((0x3fff & (q[1] >> 32)) << 2) | ((0x7fff & (q[1] >> 48)) << 1) |
		((0x00ff & q[2]) << 8) | ((0x01ff & (q[2] >> 16)) << 7) | ((0x03ff & (q[2] >> 32)) << 6) | ((0x07ff & (q[2] >> 48)) << 5) |
		((0x000f & q[3]) << 12) | ((0x01f & (q[3] >> 16)) << 11) | ((0x003f & (q[3] >> 32)) << 10) | ((0x007f & (q[3] >> 48)) << 9)

	for i := 0; i < 3; i++ {
		temp[3] |= (temp[3] & sheild) << 17
		temp[2] |= (temp[2] & sheild) << 17
		temp[1] |= (temp[1] & sheild) << 17
		temp[0] |= (temp[0] & sheild) << 17
	}
	b[3] |= temp[3]
	b[2] |= temp[2]
	b[1] |= temp[1]
	b[0] |= temp[0]

	// up
	sheild = 0x0000ffffffffffff
	temp[3] = q[3]
	temp[2] = q[2] |
		(0xffff & q[3]) | (0xffff & (q[3] >> 16)) | (0xffff & (q[3] >> 32)) | (0xffff & (q[3] >> 48))
	temp[1] = q[1] |
		(0xffff & q[2]) | (0xffff & (q[2] >> 16)) | (0xffff & (q[2] >> 32)) | (0xffff & (q[2] >> 48)) |
		(0xffff & q[3]) | (0xffff & (q[3] >> 16)) | (0xffff & (q[3] >> 32)) | (0xffff & (q[3] >> 48))
	temp[0] = q[0] |
		(0xffff & q[1]) | (0xffff & (q[1] >> 16)) | (0xffff & (q[1] >> 32)) | (0xffff & (q[1] >> 48)) |
		(0xffff & q[2]) | (0xffff & (q[2] >> 16)) | (0xffff & (q[2] >> 32)) | (0xffff & (q[2] >> 48)) |
		(0xffff & q[3]) | (0xffff & (q[3] >> 16)) | (0xffff & (q[3] >> 32)) | (0xffff & (q[3] >> 48))

	for i := 0; i < 3; i++ {
		temp[3] |= (temp[3] & sheild) << 16
		temp[2] |= (temp[2] & sheild) << 16
		temp[1] |= (temp[1] & sheild) << 16
		temp[0] |= (temp[0] & sheild) << 16
	}
	b[3] |= temp[3]
	b[2] |= temp[2]
	b[1] |= temp[1]
	b[0] |= temp[0]

	// up right
	sheild = 0x0000fffefffefffe
	temp[3] = q[3]
	temp[2] = q[2] |
		((0xfff0 & q[3]) >> 4) | ((0xfff8 & (q[3] >> 16)) >> 3) | ((0xfffc & (q[3] >> 32)) >> 2) | ((0xfffe & (q[3] >> 48)) >> 1)
	temp[1] = q[1] |
		((0xfff0 & q[2]) >> 4) | ((0xfff8 & (q[2] >> 16)) >> 3) | ((0xfffc & (q[2] >> 32)) >> 2) | ((0xfffe & (q[2] >> 48)) >> 1) |
		((0xff00 & q[3]) >> 8) | ((0xff80 & (q[3] >> 16)) >> 7) | ((0xffc0 & (q[3] >> 32)) >> 6) | ((0xffe0 & (q[3] >> 48)) >> 5)
	temp[0] = q[0] |
		((0xfff0 & q[1]) >> 4) | ((0xfff8 & (q[1] >> 16)) >> 3) | ((0xfffc & (q[1] >> 32)) >> 2) | ((0xfffe & (q[1] >> 48)) >> 1) |
		((0xff00 & q[2]) >> 8) | ((0xff80 & (q[2] >> 16)) >> 7) | ((0xffc0 & (q[2] >> 32)) >> 6) | ((0xffe0 & (q[2] >> 48)) >> 5) |
		((0xf000 & q[3]) >> 12) | ((0xf800 & (q[3] >> 16)) >> 11) | ((0xfc00 & (q[3] >> 32)) >> 10) | ((0xfe00 & (q[3] >> 48)) >> 9)

	for i := 0; i < 3; i++ {
		temp[3] |= (temp[3] & sheild) << 15
		temp[2] |= (temp[2] & sheild) << 15
		temp[1] |= (temp[1] & sheild) << 15
		temp[0] |= (temp[0] & sheild) << 15
	}
	b[3] |= temp[3]
	b[2] |= temp[2]
	b[1] |= temp[1]
	b[0] |= temp[0]

	// reverse
	b[3] = ^b[3]
	b[2] = ^b[2]
	b[1] = ^b[1]
	b[0] = ^b[0]

	return b
}
