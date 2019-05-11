package knapsack

import "fmt"
import "math"


//穷举法解决背包问题
func Exhaustive()(int, int){
	var i int = 2
	var j int = 1
	var CurrentWeight int = 0
	var CurrentValue  int = 0
	var MaxValue  int = 0
	var BestSolution int = 0
	var AllSolution   = int( math.Exp2( float64(NumofItems) ) )
	for i < AllSolution { //i = 0显然不是最优解，故直接从1开始，并到AllSolution - 1 停止
		j = 1
		CurrentWeight = 0
		CurrentValue = 0
		for j < NumofItems && CurrentWeight < KCapacity {
			var choise = Choose(i, j)
			CurrentWeight += choise * Weight[j]
			CurrentValue += choise * Value[j]
			j += 1
		}
		if MaxValue < CurrentValue && CurrentWeight< KCapacity {
			MaxValue = CurrentValue
			BestSolution = i
		}
		i += 1
	}
	return BestSolution, MaxValue
}

//获得 x 的二进制表示的第 y位的值，
func Choose (x , y int) int{
	var tmpx uint = uint(x) >> uint(y)
	return int(tmpx & 1)
}

//输出穷举法结果
func ExhaustivePrint(best, max int) {
	var j int = 0
	fmt.Println("the Exhaustive method : ")
	fmt.Println("the Max Value is :" ,max)
	fmt.Println("The Best Choice is ")
	for j < NumofItems {
		if Choose(best, j) == 1 {
			fmt.Printf("[%d, %d]", Weight[j], Value[j])
		}
		j += 1
	}
}