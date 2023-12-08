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
	fmt.Printf("%v\n", currentNodeList)
	for !isAllZs(currentNodeList) {
		nextStep := directions[steps%len(directions)]
		//fmt.Printf("nextStep: %v\n", string(nextStep))
		for i, v := range currentNodeList {
			currentNodeList[i] = getNextNode(dirMap, v, nextStep)
		}
		steps++
		//fmt.Printf("%v\n", currentNodeList)
	}

	fmt.Printf("%v\n", directions)
	fmt.Printf("%v\n", dirMap)
	fmt.Printf("%v\n", steps)

	fmt.Printf("%v\n", "fin")
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
