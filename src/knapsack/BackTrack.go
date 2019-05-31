package knapsack

import "fmt"


var index int = 0
var bestChoice[11] int
var currentChoice[11] int
var currentweight int = 0
var currentvalue  int = 0
var MaxValue      int = 0
func BackTrack(i int ) {
	if i >= NumofItems {
		if MaxValue < currentvalue {
			MaxValue = currentvalue
			for index = 0 ; index < NumofItems ; index ++ {
				bestChoice[index] = currentChoice[index]
			}
		}
		return
	}
	if (currentweight + Weight[i] <= KCapacity ) {
		currentChoice[i] = 1
		currentvalue += Value[i]
		currentweight += Weight[i]
		BackTrack( i + 1 )
		currentvalue -= Value[i]
		currentweight -= Weight[i]
	}
	upper := MaxUpperBound(i)
	if currentvalue + upper < MaxValue {
		return
	}
	currentChoice[i] = 0
	BackTrack( i + 1)
}
func Back() {
	BackTrack(0)
	var index int = 0
	fmt.Println("\n the Back method :")
	for index = 1 ; index < NumofItems ;index ++ {
		if bestChoice[index] == 1 {
			fmt.Printf("[%d, %d]",Weight[index], Value[index])
		}
	}
}

func MaxUpperBound(k int ) int {
	var index int
	var Upper int = 0
	for index = k + 1 ; index < NumofItems ; index ++ {
		Upper += Value[index]
	}
	return Upper
}