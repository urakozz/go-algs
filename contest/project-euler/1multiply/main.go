package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int64
		fmt.Scan(&n)
		sum := int64(0)
		sum += sumAP(n, 3)
		sum += sumAP(n, 5)
		sum -= sumAP(n, 15)

		fmt.Println(sum)
	}
}

func sumAP(max, d int64) int64 {
	if max < d {
		return 0
	}
	n := (max - 1) / d
	sum := (n)*(d + d * n) /2

	return sum
}