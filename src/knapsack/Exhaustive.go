package knapsack

import "fmt"
import "math"

func Exhaustive(){
	var i = 1
	var j = 1
	var CurrentWeight = 0
	var CurrentValue  = 0
	var MaxValue = 0
	var BestSolution = 0
	var AllSolution   = int( math.Exp2( float64(NumofItems) ) )
	for i < AllSolution {
		j = 0
		CurrentWeight = 0
		CurrentValue = 0
		for j < NumofItems && CurrentWeight < KCapacity {
			choise := Choose(i,j)
			CurrentWeight += Weight[j] * choise
			CurrentValue += Value[j] *choise
			j += 1
		}
		if MaxValue < CurrentValue && CurrentWeight< KCapacity {
			MaxValue = CurrentValue
			BestSolution = i
		}
		i += 1

	}
	j = 0
	fmt.Println("\nthe Max Value is :" ,MaxValue)
	fmt.Println("The Best Choice is ")
	for j < NumofItems {
		if Choose(BestSolution, j) == 1 {
			fmt.Printf("[%d, %d]", Weight[j], Value[j])
		}
		j += 1
	}

}

//获得 x % 2 ^ y
func Choose(x ,y int) int {
	var tmpy = float64(y)
	var tmpz = int( math.Exp2(tmpy) )
	if (x & tmpz) == 0 {
		return 0
	} else {
		return 1
	}
}