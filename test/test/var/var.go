package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	var a, b int
	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		c, success := Divide(a, b)
		fmt.Println(c, success)
	}
}
