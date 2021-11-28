package Day01

import (
	"fmt"
	"strconv"
)

func getFuelPt2(mass int) int {

	partialFuelMass := getFuel(mass)
	if partialFuelMass > 0 {
		return partialFuelMass + getFuelPt2(partialFuelMass)
	}
	return 0
}

func getFuel(mass int) int {
	return mass/3 - 2
}

func solvePt1(inputLines []string) {

	fuelSum := 0

	for _, singleLine := range inputLines {
		singleMass, _ := strconv.Atoi(singleLine)
		fuelAmt := getFuel(singleMass)
		fuelSum += fuelAmt
		fmt.Printf("mass: %v, fuel: %v\n", singleMass, fuelAmt)
	}

	fmt.Printf("Sum is %v\n", fuelSum)

}

func solvePt2(inputLines []string) {

	fuelSum := 0

	for _, singleLine := range inputLines {
		singleMass, _ := strconv.Atoi(singleLine)
		fuelAmt := getFuelPt2(singleMass)
		fuelSum += fuelAmt
		fmt.Printf("mass: %v, fuel: %v\n", singleMass, fuelAmt)
	}

	fmt.Printf("Sum is %v\n", fuelSum)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
