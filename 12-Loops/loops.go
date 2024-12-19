package main

import (
	"fmt"
	"time"
)

func main() {

	num := 0

	for num < 10 {
		fmt.Println("Incrementando num - ", num)
		num++
		time.Sleep(time.Second)
	}
}
