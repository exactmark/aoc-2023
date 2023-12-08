package Day08

import (
	"fmt"
	"strings"
)

func getNextNode(dirMap map[string]string, currentNode string, step uint8) string {
	if step == 'L' {
		return strings.Split(dirMap[currentNode], ", ")[0]
	}
	return strings.Split(dirMap[currentNode], ", ")[1]
}

func solvePt1(inputLines []string) {
	directions := inputLines[0]

	dirMap := make(map[string]string)
	for x := 2; x < len(inputLines); x++ {
		dirParts := strings.Split(inputLines[x], " = ")
		dirMap[dirParts[0]] = dirParts[1][1:9]
	}
	steps := 0
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		nextStep := directions[steps%len(directions)]
		fmt.Printf("nextStep: %v\n", string(nextStep))
		currentNode = getNextNode(dirMap, currentNode, nextStep)
		steps++
		fmt.Printf("%v\n", currentNode)
	}

	fmt.Printf("%v\n", directions)
	fmt.Printf("%v\n", dirMap)
	fmt.Printf("%v\n", steps)
	fmt.Printf("%v\n", "fin")
}

func solvePt2(inputLines []string) {
	directions := inputLines[0]

	dirMap := make(map[string]string)
	for x := 2; x < len(inputLines); x++ {
		dirParts := strings.Split(inputLines[x], " = ")
		dirMap[dirParts[0]] = dirParts[1][1:9]
	}
	steps := 0
	currentNodeList := make([]string, 0)
	for k, _ := range dirMap {
		if k[2] == 'A' {
			currentNodeList = append(currentNodeList, k)
		}
	}
	offsetList := make([]int, 0)
	cycleList := make([]int, 0)

	fmt.Printf("%v\n", currentNodeList)

	currentNode := currentNodeList[0]
	for i := range currentNodeList {
		steps = 0
		currentNode = currentNodeList[i]
		foundZ := make([]string, 0)
		foundZSteps := make([]int, 0)
		for len(foundZSteps) < 2 {
			nextStep := directions[steps%len(directions)]
			//fmt.Printf("nextStep: %v\n", string(nextStep))
			currentNode = getNextNode(dirMap, currentNode, nextStep)
			if currentNode[2] == 'Z' {
				foundZ = append(foundZ, currentNode)
				foundZSteps = append(foundZSteps, steps)
			}
			steps++
			//fmt.Printf("%v\n", currentNode)
		}
		offsetList = append(offsetList, foundZSteps[0])
		cycleList = append(cycleList, foundZSteps[1]-foundZSteps[0])
	}

	//hardcoded solution boo
	fmt.Printf("%v\n", LCM(cycleList[0], cycleList[1], cycleList[2], cycleList[3], cycleList[4], cycleList[5]))

	//for !isAllZs(currentNodeList) {
	//	nextStep := directions[steps%len(directions)]
	//	//fmt.Printf("nextStep: %v\n", string(nextStep))
	//	for i, v := range currentNodeList {
	//		currentNodeList[i] = getNextNode(dirMap, v, nextStep)
	//	}
	//	steps++
	//	//fmt.Printf("%v\n", currentNodeList)
	//}

	//fmt.Printf("%v\n", directions)
	//fmt.Printf("%v\n", dirMap)
	//fmt.Printf("%v\n", steps)

	fmt.Printf("%v\n", "fin")
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func isAllZs(list []string) bool {
	for _, val := range list {
		if val[2] != 'Z' {
			return false
		}
	}
	return true
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
