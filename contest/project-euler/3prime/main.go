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

func eratosthenesSieve(max uint64) ([]uint64) {
	mx := make([]bool, max + uint64(1))
	primes := []uint64{}
	for i := uint64(2); i <= max; i++ {
		if mx[i] == false {
			primes = append(primes, uint64(i))
			for j := i; j <= max; j +=i {
				mx[j] = true
			}
		}
	}

	return primes
}