package main

import "fmt"

func Multiple(a, b int) (t int) {
	t = a * b
	return t
}

func main() {
	var a, b int
	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		success := Multiple(a, b)
		fmt.Println(success)
	}
}
