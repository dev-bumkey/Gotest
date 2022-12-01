package main

import "fmt"

func F(n int) int {
	if n < 2 {
		return n
	}
	return F(n-2) + F(n-1)
}

func main() {
	var a int
	n, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println(n, err)
	} else {
		result := F(a)
		fmt.Println(result)
	}
}
