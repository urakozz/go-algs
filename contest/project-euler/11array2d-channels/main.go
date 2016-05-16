package main
import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"math"
)

const N=20
func main() {

	sieve := getSieve()

	trans := trans(sieve)
	diagonal := diagonal(sieve)
	diagonalBack := diagonalBack(sieve)



	products := make(chan int, 20)
	finish := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go GreatestHorizontalProduct(sieve, 4, products, wg)
	go GreatestHorizontalProduct(trans, 4, products, wg)
	go GreatestHorizontalProduct(diagonal, 4, products, wg)
	go GreatestHorizontalProduct(diagonalBack, 4, products, wg)

	go func(products chan int, wg *sync.WaitGroup){
		wg.Wait()
		close(products)
	}(products, wg)

	go func(finish chan<- struct {}, wg *sync.WaitGroup){
		gr := 0
		for p:= range products {
			if p > gr {
				gr = p
			}
		}
		fmt.Println(gr)
		finish <- struct{}{}
	}(finish, wg)


	go func(f chan<- struct{}){
		c := make(chan os.Signal, 1)
		 signal.Notify(c, os.Interrupt)
		<-c
		f <- struct{}{}
	}(finish)
	<-finish

}

func _main(){
	a := getSieve()
	var ans float64
	for i := 3; i < N+3; i++ {
		for j := 3; j < N+3; j++ {
			ans = math.Max(ans, float64(a[i][j]*a[i+1][j]*a[i+2][j]*a[i+3][j]))
			ans = math.Max(ans, float64(a[i][j]*a[i][j+1]*a[i][j+2]*a[i][j+3]))
			ans = math.Max(ans, float64(a[i][j]*a[i+1][j+1]*a[i+2][j+2]*a[i+3][j+3]))
			ans = math.Max(ans, float64(a[i][j]*a[i-1][j+1]*a[i-2][j+2]*a[i-3][j+3]))
		}
	}
	fmt.Println(ans)
}

func getSieve() (sieve [][]int){
	sieve = make([][]int, 26)
	for k, _ := range sieve {
		sieve[k] = make([]int, 26)
	}

	for i := 3; i < N+3; i++ {
		for j := 3; j < N+3; j++ {
			var n int
			fmt.Scan(&n)

			sieve[i][j]=n
		}
	}
	return
}

func GreatestHorizontalProduct(sieve [][]int, seq int, out chan<- int, wg *sync.WaitGroup) (greatest int) {
	for _, row := range sieve {
		GreatestRowProduct(row, seq, out)
	}
	wg.Done()
	return
}

func trans(sieve [][]int) ([][]int) {
	trans := make([][]int, len(sieve))
	for _, row := range sieve {
		for j, elem := range row {
			trans[j] = append(trans[j], elem)
		}
	}
	return trans
}
func diagonal(sieve [][]int) ([][]int) {
	trans := make([][]int, len(sieve)*2)
	for i:= 0; i < len(sieve); i++ {
		for j :=i; j>=0;j--{
			x := i-j
			y := j
//			fmt.Println(i,j, y, x)
			trans[i] = append(trans[i], sieve[y][x])
		}
	}
	for i:= 0; i < len(sieve)-1; i++ {
		for j :=i; j>=0;j--{
			x := len(sieve)-1-(i-j)
			y := len(sieve)-1-j
//			fmt.Println(i,j, y, x)
			trans[i] = append(trans[i], sieve[y][x])
		}
	}
	return trans
}

func diagonalBack(sieve [][]int) ([][]int) {
	trans := make([][]int, len(sieve)*2)
	for i:= 0; i < len(sieve); i++ {
		for j :=i; j>=0; j--{
			x := len(sieve) -1 -j
			y := i-j
//			fmt.Println(i,j, y, x)
			trans[i] = append(trans[i], sieve[y][x])
		}
	}
	for i:= 0; i < len(sieve)-1; i++ {
		for j :=i; j>=0;j--{
			x := j
			y := len(sieve)-1-(i-j)
//			fmt.Println(i,j, y, x)
			trans[i] = append(trans[i], sieve[y][x])
		}
	}
	return trans
}

func GreatestRowProduct(row []int, seq int, out chan<- int) {
	for i := 0; i < len(row) - seq; i++ {

		product := 1
		for j := i; j < i + seq; j++ {
			product *= row[j]
			if product == 0 {
				break
			}
		}
		out <- product
	}
}




