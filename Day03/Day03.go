package Day03

import (
	"fmt"
	"strconv"
	"strings"
)

type part struct {
	x       int
	y       int
	id      int
	partLen int
}

var symbolMap map[string]string
var partMap map[string]part
var partPtrMap map[string]*part

func populateMaps(inputLines []string) {
	symbolMap = make(map[string]string)
	partMap = make(map[string]part)
	partPtrMap = make(map[string]*part)
	for y := 0; y < len(inputLines); y++ {
		for x := 0; x < len(inputLines[y]); x++ {
			if isSymbol(inputLines[y][x]) {
				symbolMap[coordToString(x, y)] = string(inputLines[y][x])
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
				partMap[coordToString(numStart, y)] = newPart
				for ptr := numStart; ptr < numStart+newPart.partLen; ptr++ {
					partPtrMap[coordToString(ptr, y)] = &newPart
				}
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
			if _, ok := symbolMap[coordToString(x, y)]; ok {
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
	//partMap["0,0"].hasAdjacentSymbol()
	fmt.Printf("%v\n", sum)
}

func coordToString(x int, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func stringToCoord(s string) (int, int) {
	x, _ := strconv.Atoi(strings.Split(s, ",")[0])
	y, _ := strconv.Atoi(strings.Split(s, ",")[1])
	return x, y
}

func solvePt2(inputLines []string) {
	populateMaps(inputLines)
	sum := 0
	for coord, singleSymbol := range symbolMap {
		if singleSymbol == "*" {
			fmt.Printf("%v,%v\n", coord, singleSymbol)
			xGear, yGear := stringToCoord(coord)
			adjPartMap := make(map[*part]bool)
			for y := yGear - 1; y <= yGear+1; y++ {
				for x := xGear - 1; x <= xGear+1; x++ {
					if thisPart, ok := partPtrMap[coordToString(x, y)]; ok {
						adjPartMap[thisPart] = true
					}
				}
			}
			if len(adjPartMap) == 2 {
				ratio := 1
				for singlePart, _ := range adjPartMap {
					ratio *= singlePart.id
				}
				sum += ratio
			}

		}
	}

	fmt.Printf("%v\n", sum)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
