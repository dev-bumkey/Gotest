package main

import (
	"bufio" //입력을 읽기 위한 라이브러리
	"fmt"
	"os"      //표준 입력을 받기위한 라이브러리
	"strconv" //문자열을 숫자로 바꿔주는 라이브러리
	"strings" //불필요한 문자를 삭제하기 위한 라이브러리
)

func main() {

	fmt.Println("숫자를 입력하세요")

	var reader = bufio.NewReader(os.Stdin) // reader에 입력 받을 준비
	var line, _ = reader.ReadString('\n')  // 엔터(줄바꿈)할 때 까지 입력 받음
	line = strings.TrimSpace(line)         // 문자열의 공백 제거

	var n1, _ = strconv.Atoi(line) // 입력받은 값을 숫자로 바꿔서 n1에 저장

	//입력 반복
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	var n2, _ = strconv.Atoi(line) // 입력받은 값을 숫자로 바꿔서 n2에 저장

	fmt.Printf("입력하신 숫자는 %d, %d 입니다\n", n1, n2) // n1, n2 출력

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n') // 엔터(줄바꿈)할 때 까지 입력 받음
	line = strings.TrimSpace(line)    // 문자열의 공백 제거

	//연산자 조건 검사
	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if line == "*" {
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}
}
