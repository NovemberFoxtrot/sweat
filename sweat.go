package main

import (
	"fmt"
)

func floyd(start int) (int, int) {
	tortoise := f[start]
	hare := f[f[start]]

	for tortoise != hare {
		tortoise = f[tortoise]
		hare = f[f[hare]]
	}

	mu := 0
	tortoise = start

	for tortoise != hare {
		tortoise = f[tortoise]
		hare = f[hare]
		mu += 1
	}

	lam := 1
	hare = f[tortoise]

	for tortoise != hare {
		hare = f[hare]
		lam += 1
	}

	return mu, lam
}

func brent(start int) (int, int) {
	power := 1
	lam := 1

	tortoise := start
	hare := f[start]

	for tortoise != hare {
		if power == lam {
			tortoise = hare
			power *= 2
			lam = 0
		}

		hare = f[hare]
		lam += 1
	}

	mu := 0
	tortoise = start
	hare = start

	for i := 0; i < lam; i++ {
		hare = f[hare]
	}

	for tortoise != hare {
		tortoise = f[tortoise]
		hare = f[hare]
		mu += 1
	}

	return mu, lam
}

// var f = [...]int{6, 6, 0, 1, 4, 3, 3, 4, 0}
var f = [...]int{1, 5, 5, 2, 5, 9, 16, 16, 11, 6, 6, 19, 1, 12, 6, 0, 13, 8, 7, 16}

func main() {
	for i := 0; i < len(f); i++ {
		position, length := floyd(i)
		fmt.Println("floyd", "search start", i, "cycle start position", position, "cycle length", length)
		position, length = brent(i)
		fmt.Println("brent", "search start", i, "cycle start position", position, "cycle length", length)
	}
}
