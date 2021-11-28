package Day02

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePt1(inputLines []string) {

	inputParts := strings.Split(inputLines[0], ",")

	intArray := make([]int, len(inputParts))

	for k, v := range inputParts {
		singleInt, _ := strconv.Atoi(v)
		intArray[k] = singleInt
	}

	fmt.Printf("%v\n",intArray)

	//fuelSum := 0
	//
	//for _, singleLine := range inputLines {
	//	singleMass, _ := strconv.Atoi(singleLine)
	//	fuelAmt := getFuel(singleMass)
	//	fuelSum += fuelAmt
	//	fmt.Printf("mass: %v, fuel: %v\n", singleMass, fuelAmt)
	//}
	//
	//fmt.Printf("Sum is %v\n", fuelSum)

}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
