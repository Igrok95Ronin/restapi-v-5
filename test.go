package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func multiplication(num int) []int {
	results := []int{}
	for i := 1; i <= 10; i++ {
		result := num * i
		results = append(results, result)
	}
	return results
}

func main() {
	ints := multiplication(2)
	fmt.Println(ints)
}
