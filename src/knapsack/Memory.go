package knapsack

import (
	"fmt"
)

var x[11] int

func MemoryInitial() {
	var ii int = 0
	var jj int = 0
	for ii = 0; ii < 11; ii++ {
		for jj = 0; jj < 21; jj++ {
			Choice[ii][jj] = -1
		}
	}
}
func Memory(i, j int ) int{
	if Choice[i][j] != -1 {
		return Choice[i][j]
	}
	if i == NumofItems -1 {
		if j >= Weight[i] {
			Choice[i][j] = Value[i]
			return Value[i]
		}else {
			Choice[i][j] = 0
			return 0
		}
	} else {
		if j >= Weight[i] {
			var selected int  = Memory(i + 1, j -Weight[i] ) + Value[i]
			var unselected int = Memory(i + 1, j)
			result := MaxofInt(selected,unselected)
			Choice[i][j] = result
			return result
		} else {
			result := Memory(i + 1 , j )
			Choice[i][j] = result
			return result
		}
	}
}

func MemoryPrint() {
	var index int = 0
	MemoryTrace(1,KCapacity)
	fmt.Println("\nthe Memory method :")
	for index = 1 ; index < NumofItems ; index ++ {
		if x[index] == 1 {
			fmt.Printf("[%d, %d]",Weight[index], Value[index])
		}
	}

}

func MemoryTrace(i, j int ) {
	if i == NumofItems -1 {
		if Choice[i][j] == Value[i] {
			x[i] = 1
		} else {
			x[i] = 0
		}
		return
	} else {
		if Choice[i][j] == Choice[i+1][j] {
			x[i] = 0
			MemoryTrace( i + 1,j )
		} else {
			x[i] = 1
			MemoryTrace(i+1, j - Weight[i])
		}
	}
}