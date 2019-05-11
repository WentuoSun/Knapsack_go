package knapsack

import "fmt"

var choice = [11][21]int {}

func Maxofuint(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Dynamic() {
	var i int = 0
	var j int = 1

	for i = 0 ; i <= KCapacity ; i ++ {
		choice[0][i] = 0
	}
	for i = 0 ; i< NumofItems ; i ++ {
		choice[i][0] = 0
	}

	for i = 1 ; i < NumofItems ; i++ {
		for j = 1 ; j <= KCapacity ; j ++ {
			choice[i][j] = choice[i-1][j]
			if Weight[i] <= j {
				choice[i][j] = Maxofuint( choice[i-1][j], choice[ i-1 ][ j-Weight[i] ] + Value[i] )
			}
		}
	}
}

//输出动态规划结果
func DynamicPrint() {
	var j int = KCapacity
	fmt.Println("\nthe Dynamic method : ")
	for i := NumofItems - 1; i >= 1; i-- {
		if j >= Weight[i] {
			if choice[i][j] == choice[i-1][j-Weight[i]]+Value[i] {
				fmt.Printf("[%d, %d]", Weight[i], Value[i])
				j -= Weight[i]
			}
		}
	}
}