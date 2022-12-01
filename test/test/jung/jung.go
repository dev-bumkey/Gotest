package main

import "fmt"

// 부자가 있는지
func habeRich() bool {
	return true
}

// 친구 수
func countFriend() int {
	return 3
}

func main() {
	var price int
	fmt.Print("가격을 입력하세요 :")
	n, err := fmt.Scanf("%d", &price)
	if err != nil {
		fmt.Println(n, err)
	} else {
		if price > 50000 {
			if habeRich() {
				fmt.Println("신발끈 묶자")
			} else {
				fmt.Println("돈 나눠내자")
			}
		} else if price >= 30000 && price < 50000 {
			if countFriend() > 3 {
				fmt.Println("신발끈 묶자")
			} else {
				fmt.Println("돈 나눠내자")
			}
		} else {
			fmt.Println("알수없음")
		}
	}
}
