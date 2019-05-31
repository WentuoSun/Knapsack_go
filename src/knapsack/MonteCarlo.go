package knapsack

import (
	"math/rand"
	"time"
)

func Montecarlo() (int,  int ){
	for true {
		maxv_1, sol_1  := TrySearch()
		maxv_2, sol_2  := TrySearch()
		if maxv_1 == maxv_2 && sol_1 == sol_2 {
			return maxv_1, sol_1
		}
	}
	return 0, 0
}
func TrySearch() (int, int){
	var try int = 1000
	var i, j int
	var CWeight int = 0
	var CValue int = 0
	var MValue int = 0
	var Sollution int = 0
	//RandSource := rand.NewSource(time.Now().Unix())
	rand.Seed( time.Now().Unix() )

	for i = 0 ; i < try ; i++ {
		CValue = 0
		CWeight = 0
		RandNum :=rand.Intn( AllSolution - 1 ) + 1
		for j = 1 ; j < NumofItems && CWeight < KCapacity ; j++ {
			choice := Choose(RandNum, j)
			CWeight += choice * Weight[j]
			CValue += choice * Value[j]
		}
		if MValue < CValue && CWeight <= KCapacity{
			MValue = CValue
			Sollution = RandNum
		}
	}
	return MValue, Sollution
}
