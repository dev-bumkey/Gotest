package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요")
		var num int
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("숫자를 입력하세요")
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다\n", num)
		if num%2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료되었습니다")
}
