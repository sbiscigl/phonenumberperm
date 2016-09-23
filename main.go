package main

import (
	"github.com/sbiscigl/phonenumberperm/intstack"
	"github.com/sbiscigl/phonenumberperm/permutations"
)

func main() {
	num := []int{6, 5, 1, 2, 7, 1, 4, 5, 5, 2}
	stack := intstack.New(num)
	permutations.CalcWords(stack, "")
}
