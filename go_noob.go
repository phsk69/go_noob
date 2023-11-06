package main

import (
	"fmt"
)

func findLowestHouse(target int) int {
	houses := make([]int, target/10+1)
	for i := 1; i < len(houses); i++ {
		for j := i; j < len(houses); j += i {
			houses[j] += i * 10
		}
	}
	for i, presents := range houses {
		if presents >= target {
			return i
		}
	}
	return 0
}

func main() {
	targetPresents := 36000000 // Replace with your input value
	fmt.Println(findLowestHouse(targetPresents))
}
