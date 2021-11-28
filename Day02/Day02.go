package Day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Computer struct {
	CurrentPtr   int
	Instructions *[]int
	Message      string
}

func (c *Computer) printState() {
	fmt.Printf("ptr at %v\n", c.CurrentPtr)
	fmt.Printf("instr: %v\n", c.Instructions)
	if c.Message != "" {
		fmt.Printf("instr: %v\n", c.Message)
	}
}

func (c *Computer) canProcess() bool {
	if c.CurrentPtr >= len(*(c.Instructions)) {
		return false
	}
	if !isKnownOp((*(c.Instructions))[c.CurrentPtr]) {
		return false
	}
	return true
}

func (c *Computer) processNextOp() {
	nextOp := (*c.Instructions)[c.CurrentPtr]
	if nextOp == 1 {
		c.processAdd()
	} else if nextOp == 2 {
		c.processMuli()
	} else if nextOp == 99 {
		c.Message = fmt.Sprintf("completed at ptr %v", c.CurrentPtr)
		c.CurrentPtr = len(*(c.Instructions))
	}

}

func (c *Computer) processAdd() {
	if c.CurrentPtr+3 > len((*c.Instructions)) {
		c.Message = fmt.Sprintf("illegal add at ptr %v", c.CurrentPtr)
		c.CurrentPtr = len(*(c.Instructions))
		return
	}
	for i := 1; i <= 3; i++ {
		opPtr := (*c.Instructions)[c.CurrentPtr+i]
		if opPtr > len(*(c.Instructions)) {
			c.Message = fmt.Sprintf("illegal add at ptr %v", c.CurrentPtr)
			c.CurrentPtr = len(*(c.Instructions))
			return
		}
	}

	theSum := (*c.Instructions)[(*c.Instructions)[c.CurrentPtr+1]] + (*c.Instructions)[(*c.Instructions)[c.CurrentPtr+2]]
	(*c.Instructions)[(*c.Instructions)[c.CurrentPtr+3]] = theSum
	c.CurrentPtr += 4
}

func (c *Computer) processMuli() {
	if c.CurrentPtr+3 > len((*c.Instructions)) {
		c.Message = fmt.Sprintf("illegal muli at ptr %v", c.CurrentPtr)
		c.CurrentPtr = len(*(c.Instructions))
		return
	}
	for i := 1; i <= 3; i++ {
		opPtr := (*c.Instructions)[c.CurrentPtr+i]
		if opPtr > len(*(c.Instructions)) {
			c.Message = fmt.Sprintf("illegal muli at ptr %v", c.CurrentPtr)
			c.CurrentPtr = len(*(c.Instructions))
			return
		}
	}

	theSum := (*c.Instructions)[(*c.Instructions)[c.CurrentPtr+1]] * (*c.Instructions)[(*c.Instructions)[c.CurrentPtr+2]]
	(*c.Instructions)[(*c.Instructions)[c.CurrentPtr+3]] = theSum
	c.CurrentPtr += 4
}

func isKnownOp(theOp int) bool {
	knownOps := [3]int{1, 2, 99}
	for _, v := range knownOps {
		if theOp == v {
			return true
		}
	}
	return false
}

func solvePt1(inputLines []string) {

	inputParts := strings.Split(inputLines[0], ",")

	intArray := make([]int, len(inputParts))

	for k, v := range inputParts {
		singleInt, _ := strconv.Atoi(v)
		intArray[k] = singleInt
	}

	fmt.Printf("%v\n", intArray)

	shipPc := Computer{
		CurrentPtr:   0,
		Instructions: &intArray,
	}

	for shipPc.canProcess() {
		shipPc.processNextOp()
		shipPc.printState()
	}
}

func solvePt2(inputLines []string) {
	inputParts := strings.Split(inputLines[0], ",")

	intArray := make([]int, len(inputParts))

	for k, v := range inputParts {
		singleInt, _ := strconv.Atoi(v)
		intArray[k] = singleInt
	}

	for x := 0; x <= 99; x++ {

		for y := 0; y <= 99; y++ {
			arrCopy := make([]int, len(intArray))
			copy(arrCopy, intArray)
			arrCopy[1] = x
			arrCopy[2] = y
			shipPc := Computer{
				CurrentPtr:   0,
				Instructions: &arrCopy,
			}

			for shipPc.canProcess() {
				shipPc.processNextOp()
				//shipPc.printState()
			}
			if (*shipPc.Instructions)[0] == 19690720 {
				fmt.Printf("found\n")
				fmt.Printf("found %v and %v\n", x, y)
				fmt.Printf("result %v\n", 100*x+y)
			}

		}
	}

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
