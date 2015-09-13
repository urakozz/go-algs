package main
import (
	"fmt"
)


func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		res := process(n)

		fmt.Println(res)
	}
}

func process(n int) (res int) {

	res = 1
	for i := 1; i <= n; i++{
		res = lcm(res, i)
	}

	return
}

func gcd(a, b int) int {
	for b > 0 {
		temp := b;
		b = a % b; // % is remainder
		a = temp;
	}
	return a;
}

func gcdList(list []int) int {
	res := list[0]
	for i := 1; i < len(list); i++{
		res = gcd(res, list[i])
	}
	return res
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b));
}

func lcmList(list []int) int {
	res := list[0]
	for i := 1; i < len(list); i++{
		res = lcm(res, list[i])
	}
	return res
}