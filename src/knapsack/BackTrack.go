package knapsack

import "fmt"
var CurrentWeight,  CurrentValue int = 0, 0
var MaxValue int = 0
var index int = 0
var BestChoice[11] int
var CurrentChoice[11] int

func BackTrack(i int ) {
	if i >= NumofItems {
		if MaxValue < CurrentValue {
			MaxValue = CurrentValue
			for index = 0 ; index < NumofItems ; index ++ {
				BestChoice[index] = CurrentChoice[index]
			}
		}
		return
	}
	if (CurrentWeight + Weight[i] <= KCapacity ) {
		CurrentChoice[i] = 1
		CurrentValue += Value[i]
		CurrentWeight += Weight[i]
		BackTrack( i + 1 )
		CurrentValue -= Value[i]
		CurrentWeight -= Weight[i]
	}
	CurrentChoice[i] = 0
	BackTrack( i + 1)
}
func Back() {
	BackTrack(0)
	var index int = 0
	fmt.Println("\n the Back method :")
	for index = 1 ; index < NumofItems ;index ++ {
		if BestChoice[index] == 1 {
			fmt.Printf("[%d, %d]",Weight[index], Value[index])
		}
	}
}