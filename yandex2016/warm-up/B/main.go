// Copyright 2016 Kozyrev Yury
package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"bufio"
	//"time"
)

func main() {

	//t := time.Now()
	var n, c, min, max int
	file, _ := os.Open("input.txt")
	in := bufio.NewReader(file)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		bytes, _ := in.ReadBytes(' ')
		c = toInt(bytes)

		if i == 0 {
			max, min = c, c
		} else if c > max {
			max = c
		} else if c < min {
			min = c
		}
		//if (i+1) % 100000 == 0 {
		//	fmt.Println(i, time.Since(t))
		//}
	}
	fmt.Println(max - min)
}
func toInt(buf []byte) (n int) {
	//fmt.Println(string(buf))
	for _, v := range buf {
		if v > '9' || v < '0'{
			continue
		}
		n = n*10 + int(v-'0')
	}

	//fmt.Println(n)
	return
}

func _main() {
	file, _ := os.OpenFile("input.txt", os.O_RDWR | os.O_CREATE | os.O_TRUNC | os.O_SYNC, 0775)
	file.WriteString("1000000\n")
	for i := 0; i < 1000000; i++ {
		file.WriteString(strconv.Itoa(int(rand.Int31())))
		file.Write([]byte(" "))
		if i % 10000 == 0 {
			fmt.Println(i)
		}
	}
}

