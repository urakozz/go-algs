package main
import (
	"fmt"
)



func main(){
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		_, sum := process(n)

		fmt.Println(sum)
	}
}

func process(n int) (int, int64) {

	var factors []uint64
	var sum int64
	for i := 1; len(factors) <= n; i++ {
		sum = NewArithmetic(int64(1), int64(1)).Sum(int64(i))
		factors = idiotDivision(uint64(sum))
	}

	return len(factors), sum
}

type Arithmetic struct {
	First      int64
	Difference int64
}

func NewArithmetic(a, n int64) *Arithmetic {
	return &Arithmetic{a, n}
}

func (a Arithmetic) Item(n int64) int64 {
	return a.First + (n-int64(1))*a.Difference
}

func (a Arithmetic) Sum(n int64) int64 {
	return n * (a.First + a.Item(n)) / 2
}

func idiotDivision(max uint64) (factors []uint64) {
	_, trial := trialDivision(max)
	factors = append(factors, uint64(1))
	for i := 2; i <= int(max); i++ {
		if max % uint64(i) == 0 {
			factors = append(factors, uint64(i))
		}
	}
	return
}

func trialDivision(max uint64, sieve []uint64) (maxPrime uint64, factors []uint64) {

	if max <= uint64(2){
		return max, []uint64{uint64(2)}
	}

	maxPrime = uint64(2)

	for _, prime := range sieve {
		if max < prime*prime {
			break
		}
		for max % prime == 0 {
			factors = append(factors, prime)
			maxPrime = prime
			max = max / prime
		}
	}
	if max > 1 {
		factors = append(factors, max)
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






