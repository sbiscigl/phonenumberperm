package permutations

import (
	"sync"

	"github.com/sbiscigl/phonenumberperm/intstack"
)

var letterMap = map[int][]string{
	2: []string{"a", "b", "c"},
	3: []string{"d", "e", "f"},
	4: []string{"g", "h", "i"},
	5: []string{"j", "k", "l"},
	6: []string{"m", "n", "o"},
	7: []string{"p", "q", "r", "s"},
	8: []string{"t", "u", "v"},
	9: []string{"w", "x", "y", "z"},
}

/*CalcWords uses recursion to find posisble words*/
func CalcWords(s intstack.IntStack, word string) {
	if s.IsEmpty() {
		/* Commented out timing is pure calculation*/
		//fmt.Println(word)
	} else {
		/*Check to see if the values are 1 or zero as they*/
		/*have no letters associated with them*/
		if s.Peek() == 1 || s.Peek() == 0 {
			s.Pop()
			CalcWords(s, word)
		} else {
			for _, letter := range letterMap[s.Pop()] {
				CalcWords(s, word+letter)
			}
		}
	}
}

/*ThreadSafeCalcWords uses recursion to find posisble words*/
/*Uses channels for multi threading*/
func ThreadSafeCalcWords(s intstack.IntStack, word string, wg *sync.WaitGroup) {
	defer wg.Done()
	if s.IsEmpty() {
		/* Commented out timing is pure calculation*/
		//fmt.Println(word)
	} else {
		/*Check to see if the values are 1 or zero as they*/
		/*have no letters associated with them*/
		if s.Peek() == 1 || s.Peek() == 0 {
			s.Pop()
			wg.Add(1)
			go ThreadSafeCalcWords(s, word, wg)
		} else {
			for _, letter := range letterMap[s.Pop()] {
				wg.Add(1)
				go ThreadSafeCalcWords(s, word+letter, wg)
			}
		}
	}

}
