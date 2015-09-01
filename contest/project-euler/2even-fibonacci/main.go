package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n uint64
		fmt.Scan(&n)
		sum := fiboEvenSum(n)

		fmt.Println(sum)
	}
}

func fiboEvenSum(max uint64) (sum uint64) {

	prev1 := uint64(1)
	prev2 := uint64(2)

	curr := prev2

	sum += prev2

	for {
		curr = prev1 + prev2
		prev1, prev2 = prev2, curr

		if curr >= max {
			break
		}
		if curr%2 == 0 {
			sum += curr
		}

	}

	return
}