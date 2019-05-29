package knapsack


//共有十个货物，为了动态规划时方便，第一个用0占用
var Weight = []int { 0,12, 5, 7, 1, 9, 2, 6, 10, 15, 3}
var Value  = []int { 0,3,  1, 6, 4, 2, 5, 9, 7,  8,  7}

const NumofItems int = 11

const KCapacity int = 20

var Choice = [11][21]int {}

func MaxofInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}