package main

import (
	"sync"

	"github.com/sbiscigl/phonenumberperm/intstack"
	"github.com/sbiscigl/phonenumberperm/permutations"
)

// func consumerFunc(ch chan string) {
// 	for {
// 		fmt.Println(<-ch)
// 	}
// }

var wg sync.WaitGroup

func main() {
	num := []int{2, 7, 1, 4, 5, 5, 2}

	stack := intstack.New(num)
	// start := time.Now()
	//permutations.CalcWords(stack, "")
	// diff := time.Since(start).Nanoseconds()
	// fmt.Printf("time elapsed : %d\n", diff)

	//start := time.Now()
	permutationChannel := make(chan string)
	wg.Add(1)
	go permutations.ThreadSafeCalcWords(stack, "", permutationChannel, &wg)
	wg.Wait()
	//consumerFunc(permutationChannel)
	//diff := time.Since(start).Nanoseconds()
	//fmt.Printf("time elapsed : %d\n", diff)
}
