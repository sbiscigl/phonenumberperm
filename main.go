package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/sbiscigl/phonenumberperm/intstack"
	"github.com/sbiscigl/phonenumberperm/permutations"
)

var wg sync.WaitGroup

func main() {
	num := []int{2, 7, 1, 4, 5, 5, 2}

	stack := intstack.New(num)
	start := time.Now()
	permutations.CalcWords(stack, "")
	diff := time.Since(start).Nanoseconds()
	fmt.Printf("time elapsed : %d\n", diff)

	newStart := time.Now()
	wg.Add(1)
	go permutations.ThreadSafeCalcWords(stack, "", &wg)
	wg.Wait()
	newDiff := time.Since(newStart).Nanoseconds()
	fmt.Printf("time elapsed : %d\n", newDiff)
}
