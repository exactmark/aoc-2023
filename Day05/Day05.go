package Day05

import (
	"fmt"
	"strconv"
	"strings"
)

var seedToLocMemo map[int]int

type transformPoints struct {
	destination int
	source      int
	transRange  int
}

func (p transformPoints) transform(i int) int {
	tDist := p.source - p.destination
	return i - tDist
}

func (p transformPoints) tPtsContains(i int) bool {
	if i >= p.source && i < (p.source+p.transRange) {
		return true
	}
	return false
}

type transformationMap struct {
	sourceType string
	destType   string
	mapList    *[]transformPoints
}

type transformationDictionary struct {
	dictionaries *map[string]*transformationMap
}

func (d transformationDictionary) getTransformation(s string, i int) (string, int) {
	correctDict := (*d.dictionaries)[s]
	destPt := getTransformedPt(correctDict.mapList, i)
	return correctDict.destType, destPt
}

func getTransformedPt(list *[]transformPoints, i int) int {
	ptsList := *list
	for _, singlePtMap := range ptsList {
		if singlePtMap.tPtsContains(i) {
			return singlePtMap.transform(i)
		}
	}
	return i
}

func getNumList(s string) []int {
	atoiList := strings.Split(s, " ")
	rList := make([]int, 0)
	for _, val := range atoiList {
		num, _ := strconv.Atoi(val)
		rList = append(rList, num)
	}
	return rList
}

func makeSeedList(s string) []int {
	numPart := strings.Split(s, ": ")[1]
	return getNumList(numPart)
}

func makeTransformationDictionary(s []string) *transformationDictionary {
	returnDict := transformationDictionary{dictionaries: nil}
	dictionaries := make(map[string]*transformationMap, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
		} else {
			var newMap *transformationMap
			newMap, i = getMap(s, i)
			dictionaries[newMap.sourceType] = newMap
		}
	}
	returnDict.dictionaries = &dictionaries
	return &returnDict
}

func getMap(s []string, i int) (*transformationMap, int) {
	newMap := transformationMap{
		sourceType: "",
		destType:   "",
		mapList:    nil,
	}
	mapPart := strings.Split(s[i], " ")[0]
	namesPart := strings.Split(mapPart, "-")
	newMap.sourceType = namesPart[0]
	newMap.destType = namesPart[2]
	i++
	mapList := make([]transformPoints, 0)
	for ; i < len(s); i++ {
		if s[i] == "" {
			break
		}
		numList := getNumList(s[i])
		pointSet := transformPoints{
			destination: numList[0],
			source:      numList[1],
			transRange:  numList[2],
		}
		mapList = append(mapList, pointSet)
	}
	newMap.mapList = &mapList
	return &newMap, i
}

func getLocationFromSeed(tDict *transformationDictionary, i int) int {
	//if val,ok:=seedToLocMemo[i];ok{
	//	return val
	//}
	nextSource := "seed"
	currentPtr := i
	for nextSource != "location" {
		nextSource, currentPtr = tDict.getTransformation(nextSource, currentPtr)
	}
	//seedToLocMemo[i]=currentPtr
	return currentPtr
}

func solvePt1(inputLines []string) {
	seedList := makeSeedList(inputLines[0])

	transDict := makeTransformationDictionary(inputLines[2:])

	location := getLocationFromSeed(transDict, seedList[0])
	for _, seed := range seedList[1:] {
		newLoc := getLocationFromSeed(transDict, seed)
		if newLoc < location {
			location = newLoc
		}
	}

	fmt.Printf("%v\n", location)
	//for x:=0;x<=100;x++{
	//	_, ptr := transDict.getTransformation("seed", x)
	//	fmt.Printf("%v,%v\n",x,ptr)
	//}
	//destination, ptr := transDict.getTransformation("seed", 50)
	//fmt.Printf("%v\n", transDict)
	//fmt.Printf("%v\n", destination)
	//fmt.Printf("%v\n", ptr)

}

func solvePt2Single(inputLines []string) {
	seedRanges := makeSeedList(inputLines[0])
	transDict := makeTransformationDictionary(inputLines[2:])
	location := getLocationFromSeed(transDict, seedRanges[0])

	for i := 0; i < len(seedRanges); i += 2 {
		for j := seedRanges[i]; j < seedRanges[i]+seedRanges[i+1]; j++ {
			//newLoc:=getLocationFromSeed(transDict,j)
			newLoc := j
			if newLoc < location {
				location = newLoc
			}
		}
		fmt.Printf("%v of %v\n", i/2, len(seedRanges)/2)
	}

	fmt.Printf("%v\n", location)
}

func getLowestLocationInRange(c chan int, start int, seedRange int, transDict *transformationDictionary) {
	location := getLocationFromSeed(transDict, start)
	for i := start; i < start+seedRange; i++ {
		newLoc := getLocationFromSeed(transDict, i)
		if newLoc < location {
			location = newLoc
		}
	}
	c <- location
}

func solvePt2(inputLines []string) {
	seedRanges := makeSeedList(inputLines[0])
	transDict := makeTransformationDictionary(inputLines[2:])
	location := getLocationFromSeed(transDict, seedRanges[0])

	c := make(chan int)
	numChannels := 0
	for i := 0; i < len(seedRanges); i += 2 {
		go getLowestLocationInRange(c, seedRanges[i], seedRanges[i+1], transDict)
		numChannels++
	}

	var newLoc int
	for i := 0; i < numChannels; i++ {
		newLoc = <-c
		if newLoc < location {
			location = newLoc
		}
		fmt.Printf("received %v of %v\n", i, numChannels)
	}

	fmt.Printf("%v\n", location)
}

func Solve(inputLines []string) {
	//seedToLocMemo=make(map[int]int)
	//	solvePt1(inputLines)
	solvePt2(inputLines)
}
