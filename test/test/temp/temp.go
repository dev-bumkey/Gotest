package main

import "fmt"

func main() {
	temp := 30
	rain := 40

	if temp >= 25 {
		if rain >= 80 {
			fmt.Println("비가옵니다")
		} else if rain >= 20 {
			fmt.Println("덥고 습합니다")
		} else {
			fmt.Println("야외 가능")
		}
	} else if temp < 10 || rain >= 80 {
		fmt.Println("야외활동 ㄴㄴ")
	} else {
		fmt.Println("좋은날씨")
	}
}
