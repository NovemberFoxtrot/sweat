package main

import (
	"log"
	"os"
	"sir"
	"strconv"
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

var f []int

func main() {
	if len(os.Args) < 1 {
		log.Fatal("not enough args")
	}

	max := 0

	for i := 1; i < len(os.Args); i++ {
		arg, err := strconv.Atoi(os.Args[i])

		sir.CheckError(err)

		if arg > max {
			max = arg
		}
	}

	if len(os.Args[1:]) < max {
		log.Fatal(max, " points to a location not in the array")
	}

	log.Println(max)

	f = make([]int, 0)

	for i := 1; i < len(os.Args); i++ {
		arg, err := strconv.Atoi(os.Args[i])

		sir.CheckError(err)

		f = append(f, arg)
	}

	log.Println("\t", "[start pos]", "[cycle start pos]", "[cycle len]")

	for i := 0; i < len(f); i++ {
		fpos, flen := floyd(i)
		log.Println("floyd", i, "\t", fpos, "\t", flen)

		bpos, blen := brent(i)
		log.Println("brent", i, "\t", bpos, "\t", blen)

		if fpos != bpos || flen != blen {
			log.Fatal("[error]", "floyd and brent do not agree!")
		}
	}

	log.Println(f)
}
