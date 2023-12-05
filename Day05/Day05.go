package Day05

import (
	"fmt"
	"strconv"
	"strings"
)

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
	nextSource := "seed"
	currentPtr := i
	for nextSource != "location" {
		nextSource, currentPtr = tDict.getTransformation(nextSource, currentPtr)
	}
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

func solvePt2(inputLines []string) {
	fmt.Printf("%v\n", 1)
}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
