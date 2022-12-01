package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{100, 200, 300, 400}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a
	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}

// func main() {
// 	//var names [3]string = [3]string{"a","b","c"}
// 	names := [...]string{"a", "b", "c", "d", "e"}
// 	for idx, name := range names {
// 		fmt.Printf("변수 %d번째 알파벳은 %s 입니다 \n", idx+1, name)
// 	}
// }

// func main() {
// 	//var names [3]string = [3]string{"a","b","c"}
// 	names := [...]string{"a", "b", "c", "d", "e"}
// 	for _, name := range names {
// 		fmt.Printf("알파벳은 %s 입니다 \n", name)
// 	}
// }
