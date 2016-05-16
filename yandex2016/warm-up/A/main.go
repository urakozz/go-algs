// Copyright 2016 Kozyrev Yury
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	word := s.Bytes()
	fmt.Println(solve(word))
}

func solve(word []byte) string {
	var v, c []byte
	for _, r := range word {
		if isVowel(r) {
			v = append(v, r)
		} else {
			c = append(c, r)
		}
	}
	vs := string(v)
	cs := string(c)
	if vs > cs {
		return "Vowel"
	} else if cs > vs {
		return "Consonant"
	} else if len(v) > len(c) {
		return "Vowel"
	} else {
		return "Consonant"
	}
}

func isVowel(r byte) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return true
	}
	return false
}

