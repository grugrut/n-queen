package naive

import (
	"fmt"
)

var result int

func Solve(N int) {
	q := make([]int, N)
	backtrack(N, q, 0, 0)
	fmt.Println(result)
}

func check(N int, q []int, idx int) bool {
	ix := q[idx] % N
	iy := q[idx] / N

	for j := 0; j <= idx-1; j++ {
		// left
		for x, y := ix, iy; x >= 0; x = x - 1 {
			if x+N*y == q[j] {
				return false
			}
		}

		// left up
		for x, y := ix, iy; x >= 0 && y >= 0; x, y = x-1, y-1 {
			if x+N*y == q[j] {
				return false
			}
		}

		// up
		for x, y := ix, iy; y >= 0; y = y - 1 {
			if x+N*y == q[j] {
				return false
			}
		}

		// right up
		for x, y := ix, iy; x < N && y >= 0; x, y = x+1, y-1 {
			if x+N*y == q[j] {
				return false
			}
		}
	}

	return true
}

func backtrack(N int, q []int, idx int, next int) {
	for i := next; i < N*N; i++ {
		q[idx] = i
		if check(N, q, idx) {
			if N == idx+1 {
				fmt.Println(q, "OK")
				result++
			} else {
				backtrack(N, q, idx+1, i+1)
			}
		}
	}
}
