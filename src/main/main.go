package main

import (
	"fmt"
	"knapsack"
)
import "time"

func main() {
	start := time.Now().UnixNano()
	knapsack.Exhaustive()
	end := time.Now().UnixNano()
	fmt.Printf("\nthe Exhaustive method :%d us", (end-start) /1000 )
}

