package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Hello")
	fmt.Println(findThaiID("1234567890121"))
}

func findThaiID(id string) bool {
	if len(id) != 13 {
		return false
	}
	lastIndex := len(id) - 1
	lastNum, _ := strconv.Atoi(string(id[lastIndex]))

	sum := 0
	for i := 0; i <= 11; i++ {
		num, _ := strconv.Atoi(string(id[i]))
		sum += num * (13 - i)
	}

	return (11-sum%11)%10 == lastNum
}
