package main
import (
	"fmt"

	"math"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n uint64
		fmt.Scan(&n)
		prime := trialDivision(n)

		fmt.Println(prime)
	}
}

func trialDivision(max uint64) (maxPrime uint64) {

	if max <= uint64(2){
		return max
	}

	sieve := eratosthenesSieve(uint64(math.Sqrt(float64(max)))+ uint64(1))

	maxPrime = uint64(2)

	for _, prime := range sieve {
		if max < prime*prime {
			break
		}
		for max % prime == 0 {
			maxPrime = prime
			max = max / prime
		}
	}
	if max > 1 {
		maxPrime = max
	}
	return
}

func eratosthenesSieve(max int) []int {
	mx := make([]bool, max + 1)
	for i := 2; i <= max; i += 2 {
		mx[i] = true
	}
	n := 1
	for i := 3; i <= max; i += 2 {
		if mx[i] == false {
			n++
			for j := i*i; j <= max; j += i {
				mx[j] = true
			}
		}
	}

	primes := make([]int, n)
	primes[0] = 2
	j := 0
	for i := 3; i <= max; i += 2 {
		if mx[i] == false {
			j++
			primes[j]= i
		}
	}

	return primes
}