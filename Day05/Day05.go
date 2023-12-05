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

type transformationMap struct {
	sourceType string
	destType   string
	mapList    *[]transformPoints
}

type transformationDictionary struct {
	dictionaries *map[string]*transformationMap
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

func solvePt1(inputLines []string) {
	seedList := makeSeedList(inputLines[0])
	transDict := makeTransformationDictionary(inputLines[2:])
	fmt.Printf("%v\n", seedList)
	fmt.Printf("%v\n", transDict)
}

func solvePt2(inputLines []string) {
	fmt.Printf("%v\n", 1)
}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
