package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i := i; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
