package knapsack

import "fmt"

//var choice = [11][21]int {}



func Dynamic() {
	var i int = 0
	var j int = 1

	for i = 0 ; i <= KCapacity ; i ++ {
		Choice[0][i] = 0
	}
	for i = 0 ; i< NumofItems ; i ++ {
		Choice[i][0] = 0
	}

	for i = 1 ; i < NumofItems ; i++ {
		for j = 1 ; j <= KCapacity ; j ++ {
			Choice[i][j] = Choice[i-1][j]
			if Weight[i] <= j {
				Choice[i][j] = MaxofInt( Choice[i-1][j], Choice[ i-1 ][ j-Weight[i] ] + Value[i] )
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
			if Choice[i][j] == Choice[i-1][j-Weight[i]]+Value[i] {
				fmt.Printf("[%d, %d]", Weight[i], Value[i])
				j -= Weight[i]
			}
		}
	}
}