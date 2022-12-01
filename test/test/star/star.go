package main

import "fmt"

func main() {
	var s int

	fmt.Print("몇줄을 입력하시겠습니까? :")
	n, err := fmt.Scanf("%d", &s)
	if err != nil {
		fmt.Println(n, err)
	} else {
		for i := 1; i < s+1; i++ {
			for j := 1; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println("^")
		}
	}
}
