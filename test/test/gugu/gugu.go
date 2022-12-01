package main

import "fmt"

func main() {
	var dan int

	for {
		fmt.Print("몇단을 보시겠습니까?: ")
		fmt.Scanf("%d", &dan)
		for i := 1; i < 10; i++ {
			fmt.Printf("%d x %d = %d\n", dan, i, dan*i)
		}

		if dan == 0 {
			break
		}
		fmt.Println("종료되었습니다.")
	}
}

/*func main() {
	dan := 2
	for i := 1; i < 10; i++ {
		fmt.Println("2 x ", i, "= ", dan*i)
	}
}
*/
