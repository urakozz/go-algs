
package main
import (

	"fmt"
//	"bufio"
//	"os"
//	"text/scanner"
	"math"
)

func main() {
	var y, x int64
	var pos int
	fmt.Scan(&y)
	fmt.Scan(&x)
	fmt.Scan(&pos)
	mx := make([][]int64, y)
	for k, _ := range mx {
		row := make([]int64, x)
		for kr, _ := range row {
			var i int64
			fmt.Scan(&i)
			row[kr] = i
		}
//		fmt.Println(row)
		mx[k] = row
	}

	circleN := int(math.Min(float64(x), float64(y))) / 2

	circles := make([][]int64, circleN)
	for i := 0; i < circleN; i++ {
		circles[i] = circleMe(mx, i)
	}

	for k, v := range circles {
		circles[k] = RotateMe(v, pos)
	}

	fmt.Println(circles)
	mx = MatrixMe(mx, circles)
	for _, row := range mx {
		for _, v := range row {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}


}

func circleMe(mx [][]int64, n int) (circle []int64) {
//	fmt.Println(">>>")
//	fmt.Println(n)
	firstRow := mx[n]
	for i := n; i <= len(firstRow) - n - 1; i++ { // n=0, len(firstRow)=4 -> i = 0; i < 4-0-1; i++
		circle = append(circle, firstRow[i])
	}
//	fmt.Println(circle)

	if len(mx) > (n + 1) * 2 {
		for i := n + 1 ; i < len(mx) - n - 1; i++ { // n=0, len(mx)=4 -> i = 1; i < 4-0-1; i++
			circle = append(circle, mx[i][len(mx[i]) - n - 1]) // n=0, len(mx[i])=4 -> mx[i][4-0-1]
		}
	}
//	fmt.Println(circle)

	lastRow := mx[len(mx) - n - 1] // n=0, len(mx)=4 -> mx[4-0-1]
	for i := len(lastRow) - n - 1; i >= n; i-- { // n=0, len(lastRow)=4 -> i = 4-0-1; i >= 0; i--
		circle = append(circle, lastRow[i])
	}
//	fmt.Println(circle)

	if len(mx) > (n + 1) * 2 {
		for i := len(mx) - n - 2; i > n; i-- { // n=0, len(mx)=4 -> i = 4-0-2; i > 0; i--
			circle = append(circle, mx[i][n]) // n=0 -> mx[i][0]
		}
	}
//	fmt.Println(circle)

	return
}

func RotateMe(circle []int64, pos int) (rotated []int64) {
	l := len(circle)
	rotated = make([]int64, l)
	for k, v := range circle {

		rotated[(k + l - pos%l) % l] = v
	}
	return
}

func MatrixMe(origMX, circles [][]int64) (mx [][]int64) {
	mx = make([][]int64, len(origMX))
	for k, _ := range mx {
		mx[k] = make([]int64, len(origMX[0]))
	}
	x := len(mx[0])
	y := len(mx)

	for n, circle := range circles {
		rightSideIdx := x - (n * 2)
		lastRowIdx := rightSideIdx  + y - 2 - (n * 2)
		leftSideIdx := lastRowIdx + x - (n * 2)


		for ci, v := range circle {

			if ci < rightSideIdx {
				mx[n][ci + n] = v
			}else if ci < lastRowIdx {
				cur := ci - rightSideIdx
				mx[n + 1 + cur][x - 1 - n] = v
			}else if ci < leftSideIdx {
				cur := ci - lastRowIdx
				mx[y - 1 - n][x - 1 - n - cur] = v
			}else {
				cur := ci - leftSideIdx + 1
				mx[y -1 -n -cur][n] = v
			}
		}
	}
	return
}



