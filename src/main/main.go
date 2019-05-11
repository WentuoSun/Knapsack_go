package main

import "knapsack"

func main() {
	//start := time.Now().UnixNano()

	//end := time.Now().UnixNano()
	//fmt.Printf("\nthe Exhaustive method :%d us", (end-start) /1000 )
	best, max := knapsack.Exhaustive()
	knapsack.ExhaustivePrint(best, max)
	knapsack.Dynamic()
	knapsack.DynamicPrint()




}

