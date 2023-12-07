package Day07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type hand struct {
	handRaw       string
	convertedHand *[]int
	bid           int
	score         int
	handMap       *map[int]int
	pt2score      int
}

func convertHand(s string) *[]int {
	convertedHand := make([]int, 0)
	for _, r := range s {
		cardScore := getCardScore(r)
		convertedHand = append(convertedHand, cardScore)
	}
	return &convertedHand
}

func getCardScore(r int32) int {
	switch r {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	val, _ := strconv.Atoi(string(r))
	return val
}

func (h *hand) populateScore() {
	scoreString := ""
	typeScore := getType(h.handMap)
	scoreString = fmt.Sprintf("%v%v", scoreString, typeScore)

	for _, val := range *(h.convertedHand) {
		if val < 10 {
			scoreString = fmt.Sprintf("%v%v", scoreString, 0)
		}
		scoreString = fmt.Sprintf("%v%v", scoreString, val)
	}
	score, _ := strconv.Atoi(scoreString)
	h.score = score
}

func (h *hand) populatePt2Score() {
	scoreString := ""
	typeScore := getPt2Type(h.handMap)
	scoreString = fmt.Sprintf("%v%v", scoreString, typeScore)

	for _, val := range *(h.convertedHand) {
		if val == 11 {
			val = 1
		}
		if val < 10 {
			scoreString = fmt.Sprintf("%v%v", scoreString, 0)
		}
		scoreString = fmt.Sprintf("%v%v", scoreString, val)
	}
	score, _ := strconv.Atoi(scoreString)
	h.pt2score = score
}

func getPt2Type(handMapPtr *map[int]int) int {
	handMap := *handMapPtr
	if _, ok := handMap[11]; !ok {
		return getType(handMapPtr)
	}
	if len(handMap) == 1 {
		return getType(handMapPtr)
	}
	// lets throw all jacks onto highest nonjack
	hMapCopy := make(map[int]int)
	highestValKey := 0
	highestValVal := 0
	for k, v := range handMap {
		if k != 11 {
			hMapCopy[k] = v
			if v > highestValVal {
				highestValVal = v
				highestValKey = k
			}
		}
	}
	hMapCopy[highestValKey] += handMap[11]
	return getType(&hMapCopy)
}

func getType(handMapPtr *map[int]int) int {
	handMap := *handMapPtr
	switch len(handMap) {
	case 1:
		// 5 of a kind
		return 7
	case 2:
		//four of a kind or full house
		singleKey := 0
		for k := range handMap {
			singleKey = k
			break
		}
		if handMap[singleKey] == 1 || handMap[singleKey] == 4 {
			// four of a kind
			return 6
		} else {
			// full house
			return 5
		}
	case 3:
		//	three of a kind or two pair
		singlesCount := 0
		for k := range handMap {
			if handMap[k] == 1 {
				singlesCount++
			}
		}
		if singlesCount == 2 {
			//three of a kind
			return 4
		} else {
			//two pair
			return 3
		}
	case 4:
		//one pair
		return 2
	case 5:
		//high card
		return 1
	}
	return -1
}

func makeHand(l string) *hand {
	newHand := hand{
		handRaw:       "",
		convertedHand: nil,
		bid:           0,
		score:         0,
		handMap:       nil,
	}
	lineSplit := strings.Split(l, " ")
	newHand.handRaw = lineSplit[0]
	bid, _ := strconv.Atoi(lineSplit[1])
	newHand.bid = bid
	convertedHand := convertHand(newHand.handRaw)
	newHand.convertedHand = convertedHand
	handMap := make(map[int]int)
	for _, val := range *convertedHand {
		if _, ok := handMap[val]; ok {
			handMap[val]++
		} else {
			handMap[val] = 1
		}
	}
	newHand.handMap = &handMap
	newHand.populateScore()
	newHand.populatePt2Score()

	return &newHand
}

func solvePt1(inputLines []string) {

	handList := make([]*hand, 0)

	for _, l := range inputLines {
		thisHand := makeHand(l)
		handList = append(handList, thisHand)
		//fmt.Printf("%v\n", thisHand)
		//fmt.Printf("%v\n", thisHand.convertedHand)
	}
	slices.SortFunc(handList, func(a, b *hand) int {
		return a.score - b.score
	})
	//slices.Reverse(handList)
	for _, val := range handList {
		fmt.Printf("%v\n", val)
	}
	score := 0
	for i, val := range handList {
		score += (i + 1) * (val.bid)
	}
	fmt.Printf("%v\n", score)
}

func solvePt2(inputLines []string) {

	handList := make([]*hand, 0)

	for _, l := range inputLines {
		thisHand := makeHand(l)
		handList = append(handList, thisHand)
		//fmt.Printf("%v\n", thisHand)
		//fmt.Printf("%v\n", thisHand.convertedHand)
	}
	slices.SortFunc(handList, func(a, b *hand) int {
		return a.pt2score - b.pt2score
	})
	//slices.Reverse(handList)
	for _, val := range handList {
		fmt.Printf("%v\n", val)
	}
	score := 0
	for i, val := range handList {
		score += (i + 1) * (val.bid)
	}
	fmt.Printf("%v\n", score)
	fmt.Printf("%v\n", 2)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)

}
