package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println("숫자를 입력하세요")
		c, success := Divide(a, b)
		fmt.Println(c, success)
	}
}
