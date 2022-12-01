package main

import "fmt"

func PrintAvgScore(name string, mat int, eng int, his int) {
	total := mat + eng + his
	avg := total / 3
	fmt.Println(name, "님 평균점수는 ", avg, "입니다")
}
func main() {
	PrintAvgScore("김일등", 39, 59, 23)
	PrintAvgScore("박이등", 45, 47, 75)
	PrintAvgScore("조삼등", 85, 97, 86)
}
