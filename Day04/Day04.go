package Day04

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//Sample:
//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
//Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
//Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
//Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
//Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
//Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11

func getNumberListFromString(s string) []int {
	re := regexp.MustCompile(" ")
	split := re.Split(s, -1)
	returnList := make([]int, 0)
	for _, val := range split {
		if val != "" {
			num, _ := strconv.Atoi(val)
			returnList = append(returnList, num)
		}
	}
	return returnList
}

func getIntersection(nums []int, nums2 []int) []int {
	returnList := make([]int, 0)
	leftSet := make(map[int]bool)
	for _, val := range nums {
		leftSet[val] = true
	}
	for _, val := range nums2 {
		if _, ok := leftSet[val]; ok {
			returnList = append(returnList, val)
		}
	}
	return returnList
}

func getScore(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	} else {
		return int(math.Pow(2, float64(len(nums)-1)))
	}
}

func solvePt1(inputLines []string) {

	sum := 0

	for _, line := range inputLines {
		numsOnly := strings.Split(line, ": ")[1]
		numsSplit := strings.Split(numsOnly, " |")
		leftNums := getNumberListFromString(numsSplit[0])
		rightNums := getNumberListFromString(numsSplit[1])
		listIntersection := getIntersection(leftNums, rightNums)
		sum += getScore(listIntersection)
	}

	fmt.Printf("%v\n", sum)
}

func solvePt2(inputLines []string) {
	cardScores := make([]int, len(inputLines))
	cardAmounts := make([]int, len(inputLines))
	for ptr, line := range inputLines {
		numsOnly := strings.Split(line, ": ")[1]
		numsSplit := strings.Split(numsOnly, " |")
		leftNums := getNumberListFromString(numsSplit[0])
		rightNums := getNumberListFromString(numsSplit[1])
		listIntersection := getIntersection(leftNums, rightNums)
		cardScores[ptr] = len(listIntersection)
		cardAmounts[ptr] = 1
	}
	sum := 0
	for ptr, val := range cardScores {
		for x := 1; x <= val; x++ {
			cardAmounts[ptr+x] += cardAmounts[ptr]
		}
		sum += cardAmounts[ptr]
	}
	fmt.Printf("%v\n", cardAmounts)
	fmt.Printf("%v\n", sum)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
