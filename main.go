package main

import "fmt"

func main() {
	Nth(1, 30)
}

func Nth(n int, x int) {
	if n < x {
		primes := []int{}
		if n == 1 {
			primes = append(primes, int(2))
		}
		for a := 3; len(primes) <= x; a += 2 {
			isPrime := true
			for b := 2; b < a; b++ {
				if a%b == 0 {
					isPrime = false
					b = 0
					break
				}
			}
			if isPrime {
				primes = append(primes, int(a))

			}
		}
		fmt.Print(primes[n-1 : x])
	} else if n > x {
		primes := []int{}
		if x == 1 {
			primes = append(primes, int(2))
		}
		for a := 3; len(primes) <= n; a += 2 {
			isPrime := true
			for b := 2; b < a; b++ {
				if a%b == 0 {
					isPrime = false
					b = 0
					break
				}
			}
			if isPrime {
				primes = append(primes, int(a))

			}
		}
		fmt.Print(primes[x-1 : n])
	}

}
