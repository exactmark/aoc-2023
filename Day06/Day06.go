package Day06

import (
	"fmt"
	"strconv"
	"strings"
)

func bruteGetWinningTimes(time, dist int) int {
	count := 0
	for x := 0; x <= time; x++ {
		calc := (time - x) * x
		if calc > dist {
			count++
		}
	}
	return count
}

func getNumList(s string) []int {
	numPart := strings.Split(s, ":")[1]
	numPartSplit := strings.Split(numPart, " ")
	returnList := make([]int, 0)
	for _, val := range numPartSplit {
		if val != "" {
			num, _ := strconv.Atoi(val)
			returnList = append(returnList, num)
		}
	}
	return returnList
}

func solvePt1(inputLines []string) {

	times := getNumList(inputLines[0])
	distances := getNumList(inputLines[1])

	fmt.Printf("%v\n", times)
	fmt.Printf("%v\n", distances)
	prod := 1
	for i, _ := range times {
		thisVal := bruteGetWinningTimes(times[i], distances[i])
		fmt.Printf("%v\n", thisVal)
		prod *= thisVal
	}
	fmt.Printf("%v\n", prod)
}

func getNumListComb(s string) int {
	numList := getNumList(s)
	appendString := ""
	for _, val := range numList {
		appendString = fmt.Sprintf("%v%v", appendString, val)
	}
	num, _ := strconv.Atoi(appendString)
	return num

}

func solvePt2(inputLines []string) {

	time := getNumListComb(inputLines[0])
	distance := getNumListComb(inputLines[1])

	fmt.Printf("%v\n", time)
	fmt.Printf("%v\n", distance)

	//(time - x) * x

	lowPoint := 0
	for x := 0; x <= time; x++ {
		if (time-x)*x > distance {
			lowPoint = x
			break
		}
	}
	highPoint := 0

	for x := time; x > 0; x-- {
		if (time-x)*x > distance {
			highPoint = x + 1
			break
		}
	}

	fmt.Printf("%v\n", highPoint-lowPoint)
	fmt.Printf("%v\n", 1)
}

func solvePt2Brute(inputLines []string) {

	time := getNumListComb(inputLines[0])
	distance := getNumListComb(inputLines[1])

	fmt.Printf("%v\n", time)
	fmt.Printf("%v\n", distance)

	solves := bruteGetWinningTimes(time, distance)

	fmt.Printf("%v\n", solves)
	fmt.Printf("%v\n", 1)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
	//solvePt2Brute(inputLines)
}
