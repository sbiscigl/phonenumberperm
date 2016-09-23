package intstack

import "fmt"

const (
	maxSize = 100
)

/*IntStack implimentaiton of a stack for integers*/
type IntStack struct {
	valueList []int
	maxSize   int
}

/*New returns bew instace of IntStack*/
func New(nums []int) IntStack {
	return IntStack{
		valueList: nums,
		maxSize:   maxSize,
	}
}

/*Pop pops the top value off the stack*/
func (s *IntStack) Pop() int {
	var val int
	if !s.IsEmpty() {
		val = s.valueList[0]
		s.valueList = s.valueList[1:]
	} else {
		fmt.Println("stack is empty")
	}
	return val
}

/*Peek returns top value*/
func (s IntStack) Peek() int {
	return s.valueList[0]
}

/*IsEmpty checks if the stack is empty*/
func (s IntStack) IsEmpty() bool {
	if len(s.valueList) > 0 {
		return false
	}
	return true
}

/*Print prints out the contents of the stack*/
func (s IntStack) Print() {
	for _, element := range s.valueList {
		fmt.Print(element)
	}
	fmt.Print("\n")
}
