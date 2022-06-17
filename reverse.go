package main

import "fmt"

func reverseNumber(num int) int {

	res := 0
	if num > 0 {
		for i := num; i > 0; i = i / 10 {
			remainder := i % 10
			res = (res * 10) + remainder

		}

	} else {
		for i := num; i < 0; i = i / 10 {
			remainder := i % 10
			res = (res * 10) + remainder

		}

	}
	return res
}
func main() {
	fmt.Println(reverseNumber(123))
	fmt.Println(reverseNumber(-123))
	fmt.Println(reverseNumber(120))
}
