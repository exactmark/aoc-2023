package Day03

import (
	"fmt"
)

type part struct {
	x       int
	y       int
	id      int
	partLen int
}

var symbolMap map[string]string
var partMap map[string]part

func populateMaps(inputLines []string) {
	symbolMap = make(map[string]string)
	partMap = make(map[string]part)
	for y := 0; y < len(inputLines); y++ {
		for x := 0; x < len(inputLines[y]); x++ {
			if isSymbol(inputLines[y][x]) {
				symbolMap[fmt.Sprintf("%v,%v", x, y)] = string(inputLines[y][x])
			}
			if isNumber(inputLines[y][x]) {
				numStart := x
				morePart := true
				workingId := int(inputLines[y][x] - 48)

				for morePart {
					if x+1 >= len(inputLines[y]) {
						morePart = false
					} else if !isNumber(inputLines[y][x+1]) {
						morePart = false
					} else {
						workingId *= 10
						workingId += int(inputLines[y][x+1] - 48)
						x++
					}
				}
				newPart := part{
					x:       numStart,
					y:       y,
					id:      workingId,
					partLen: len(fmt.Sprintf("%v", workingId)),
				}
				partMap[fmt.Sprintf("%v,%v", numStart, y)] = newPart
			}
		}
	}
}

func isNumber(u uint8) bool {
	if u >= 48 && u <= 57 {
		return true
	}
	return false
}

func isSymbol(u uint8) bool {
	if isNumber(u) {
		return false
	}
	if u == 46 {
		return false
	}
	return true

}

func (p part) hasAdjacentSymbol() bool {
	//Can check out of bounds and full number list! :D
	for y := p.y - 1; y <= p.y+1; y++ {
		for x := p.x - 1; x <= p.x+p.partLen; x++ {
			//fmt.Printf("checking: %v,%v\n",x,y)
			if _, ok := symbolMap[fmt.Sprintf("%v,%v", x, y)]; ok {
				return true
			}
		}
	}
	return false
}

func solvePt1(inputLines []string) {
	populateMaps(inputLines)
	sum := 0

	for _, singlePart := range partMap {
		if singlePart.hasAdjacentSymbol() {
			fmt.Printf("adding: %v\n", singlePart.id)
			sum += singlePart.id
		} else {
			fmt.Printf("skipping: %v\n", singlePart.id)
		}
	}
	partMap["0,0"].hasAdjacentSymbol()
	fmt.Printf("%v\n", sum)
}

func solvePt2(inputLines []string) {

	fmt.Printf("%v\n", 2)
}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
