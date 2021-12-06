package Day03

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type empty struct {
}

type coord struct {
	x int
	y int
}

type wireObj struct {
	wirePath        map[coord]int
	currentEnd      coord
	currentDistance int
}

func (w *wireObj) processStep(step string) {
	xInc := 0
	yInc := 0

	switch dir := step[0]; dir {
	case 'U':
		yInc = 1
	case 'D':
		yInc = -1
	case 'L':
		xInc = -1
	case 'R':
		xInc = 1
	default:
		fmt.Printf("incorrect direction")
	}
	distance, _ := strconv.Atoi(step[1:])

	for i := 0; i < distance; i++ {
		w.currentEnd.x += xInc
		w.currentEnd.y += yInc
		w.currentDistance += 1
		if _, ok := w.wirePath[w.currentEnd]; ok {
			//w.wirePath[w.currentEnd] += 1
		} else {
			w.wirePath[w.currentEnd] = w.currentDistance
		}
	}

}

func getWire(s string) wireObj {

	wire := wireObj{
		wirePath:        make(map[coord]int),
		currentEnd:      coord{0, 0},
		currentDistance: 0,
	}
	wire.wirePath[wire.currentEnd] = 1

	wireSteps := strings.Split(s, ",")

	for _, step := range wireSteps {
		wire.processStep(step)
	}

	return wire
}

func (c coord) getManhattanDistanceFromOrigin() int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))

}

func solvePt1(inputLines []string) {

	wire1 := getWire(inputLines[0])
	wire2 := getWire(inputLines[1])

	commonPts := make([]coord, 0)

	originCoord := coord{0, 0}

	for k, _ := range wire1.wirePath {
		if k != originCoord {
			if _, ok := wire2.wirePath[k]; ok {
				commonPts = append(commonPts, k)
			}
		}
	}

	closest := commonPts[0].getManhattanDistanceFromOrigin()

	for _, val := range commonPts {
		if val.getManhattanDistanceFromOrigin() < closest {
			closest = val.getManhattanDistanceFromOrigin()
		}
	}

	fmt.Printf("Closest: %v\n", closest)

}

func solvePt2(inputLines []string) {

	wire1 := getWire(inputLines[0])
	wire2 := getWire(inputLines[1])

	commonPts := make([]coord, 0)

	originCoord := coord{0, 0}

	for k, _ := range wire1.wirePath {
		if k != originCoord {
			if _, ok := wire2.wirePath[k]; ok {
				commonPts = append(commonPts, k)
			}
		}
	}

	closest := wire1.wirePath[commonPts[0]]+wire2.wirePath[commonPts[0]]

	for _, val := range commonPts {
		thisDistance:=wire1.wirePath[val]+wire2.wirePath[val]
		if thisDistance < closest {
			closest = thisDistance
		}
	}

	fmt.Printf("Closest: %v\n", closest)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
