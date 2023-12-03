package Day02

import (
	"fmt"
	"strconv"
	"strings"
)

// Sample Input
//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
//Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
//Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
//Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
//Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

type pullPart struct {
	count int
	color string
}

type gameRound struct {
	roundId      int
	maxCubeCount *map[string]int
	//pulls *[][]pullPart
}

func makeGameRound(line string) gameRound {
	roundId := getRoundIdFromLine(line)
	maxCubeCount := getMaxCubeCountFromPulls(strings.Split(line, ": ")[1])
	return gameRound{
		roundId:      roundId,
		maxCubeCount: maxCubeCount,
	}
}

func getMaxCubeCountFromPulls(s string) *map[string]int {
	//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	maxCubeCount := make(map[string]int)
	//split by pulls
	pulls := strings.Split(s, "; ")
	for _, singlePull := range pulls {
		//split by color
		pullColors := strings.Split(singlePull, ", ")
		//	add colors to max
		for _, singlePullPart := range pullColors {
			count, color := getColorCountFromDrawPart(singlePullPart)
			if currentMax, ok := maxCubeCount[color]; ok {
				if count > currentMax {
					maxCubeCount[color] = count
				}
			} else {
				maxCubeCount[color] = count
			}
		}
	}
	return &maxCubeCount
}

func getRoundIdFromLine(line string) int {
	idPart := strings.Split(line, ":")[0]
	numPart := strings.Split(idPart, " ")[1]
	idInt, _ := strconv.Atoi(numPart)
	return idInt
}

func getColorCountFromDrawPart(s string) (int, string) {
	color := strings.Split(s, " ")[1]
	numPart := strings.Split(s, " ")[0]
	count, _ := strconv.Atoi(numPart)
	return count, color
}

func roundHasMoreThanTarget(round gameRound, target string) bool {
	targetParts := strings.Split(target, ", ")
	for _, part := range targetParts {
		count, color := getColorCountFromDrawPart(part)
		if val, ok := (*round.maxCubeCount)[color]; ok {
			if val > count {
				return true
			}
		}
	}
	return false
}

func solvePt1(inputLines []string) {
	//fmt.Printf("%v\n",getRoundIdFromLine(inputLines[0]))
	//count, color := getColorCountFromDrawPart("3 blue")
	//fmt.Printf("%v\n", count)
	//fmt.Printf("%v\n", color)
	//gameRound1:=makeGameRound(inputLines[0])
	//fmt.Printf("%v\n", gameRound1)
	//fmt.Printf("%v\n",gameRound1.maxCubeCount)
	rounds := make([]gameRound, 0)
	for _, singleLine := range inputLines {
		rounds = append(rounds, makeGameRound(singleLine))
	}
	//fmt.Printf("%v\n",rounds)
	target := "12 red, 13 green, 14 blue"
	count := 0
	for _, singleRound := range rounds {
		if !roundHasMoreThanTarget(singleRound, target) {
			count += singleRound.roundId
		}
	}
	fmt.Printf("%v\n", count)
}

func solvePt2(inputLines []string) {
	rounds := make([]gameRound, 0)
	for _, singleLine := range inputLines {
		rounds = append(rounds, makeGameRound(singleLine))
	}
	powerSum := 0
	for _, singleRound := range rounds {
		thisSum := (*singleRound.maxCubeCount)["red"]
		thisSum *= (*singleRound.maxCubeCount)["blue"]
		thisSum *= (*singleRound.maxCubeCount)["green"]
		powerSum += thisSum
	}
	fmt.Printf("%v\n", powerSum)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
