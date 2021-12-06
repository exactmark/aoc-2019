package Day04

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidPw(pw int) bool {
	pwStr := fmt.Sprintf("%v", pw)
	if len(pwStr) != 6 {
		return false
	}

	foundDouble := false
	for i := 0; i < len(pwStr)-1; i++ {
		if pwStr[i] == pwStr[i+1] {
			foundDouble = true
		}
		if pwStr[i] > pwStr[i+1] {
			return false
		}
	}

	if !foundDouble {
		return false
	}

	return true
}

func solvePt1(inputLines []string) {
	//SUT := 111111
	//fmt.Printf("%v:%v\n", SUT, isValidPw(SUT))
	//SUT = 223450
	//fmt.Printf("%v:%v\n", SUT, isValidPw(SUT))
	//SUT = 123789
	//fmt.Printf("%v:%v\n", SUT, isValidPw(SUT))

	inParts := strings.Split(inputLines[0], "-")
	start, _ := strconv.Atoi(inParts[0])
	end, _ := strconv.Atoi(inParts[1])

	validCount := 0

	for i := start; i < end; i++ {
		if isValidPw(i) {
			validCount++
		}
	}
	fmt.Printf("Valid: %v\n", validCount)
}

func isValidPw2(pw int) bool {
	pwStr := fmt.Sprintf("%v", pw)
	if len(pwStr) != 6 {
		return false
	}

	for i := 0; i < len(pwStr)-1; i++ {

		if pwStr[i] > pwStr[i+1] {
			return false
		}
	}

	if !hasValidDouble(pwStr) {
		return false
	}

	return true
}

func hasValidDouble(pwStr string) bool{

	currentRun:=1
	for i:=1;i<len(pwStr);i++{
		if pwStr[i-1]==pwStr[i]{
			currentRun++
		}else if currentRun==2{
			return true
		}else {
			currentRun=1
		}
	}
	if currentRun==2{
		return true
	}
	return false
}

func solvePt2(inputLines []string) {
	SUT := 112233
	fmt.Printf("%v:%v\n", SUT, isValidPw2(SUT))
	SUT = 123444
	fmt.Printf("%v:%v\n", SUT, isValidPw2(SUT))
	SUT = 111122
	fmt.Printf("%v:%v\n", SUT, isValidPw2(SUT))

	inParts := strings.Split(inputLines[0], "-")
	start, _ := strconv.Atoi(inParts[0])
	end, _ := strconv.Atoi(inParts[1])

	validCount := 0

	for i := start; i < end; i++ {
		if isValidPw2(i) {
			validCount++
		}
	}
	fmt.Printf("Valid: %v\n", validCount)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
