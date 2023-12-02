package Day01

import (
	"fmt"
)

var numMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func convertToNumString(line string) string {
	returnLine := ""
	for i := 0; i < len(line); i++ {
		if isNum, num := ptrIsNum(line, i); isNum {
			returnLine = returnLine + fmt.Sprintf("%v", num)
		}
	}
	return returnLine
}

func ptrIsNum(line string, ptr int) (bool, int) {

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("Recovered. Error:\n", r)
		}
	}()
	if runeIsNum(line[ptr]) {
		return true, (int)(line[ptr] - 48)
	}

	if v, ok := numMap[line[ptr:ptr+3]]; ok {
		return true, v
	}
	if v, ok := numMap[line[ptr:ptr+4]]; ok {
		return true, v
	}
	if v, ok := numMap[line[ptr:ptr+5]]; ok {
		return true, v
	}
	return false, 0
}

func runeIsNum(r uint8) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}

func getNumFromLine(singleLine string) int {
	var firstNum, secondNum int
	//48-57   minus 48
	for i := 0; i < len(singleLine); i++ {
		if runeIsNum(singleLine[i]) {
			firstNum = (int)(singleLine[i] - 48)
			break
		}
	}
	for i := len(singleLine) - 1; i >= 0; i-- {
		if runeIsNum(singleLine[i]) {
			secondNum = (int)(singleLine[i] - 48)
			break
		}
	}
	return (firstNum * 10) + secondNum
}

func solvePt1(inputLines []string) {
	sum := 0
	for _, singleLine := range inputLines {
		sum += getNumFromLine(singleLine)
	}
	fmt.Printf("%v\n", sum)
}

func solvePt2(inputLines []string) {

	//testLine := "1xtwone3four"
	//newString := convertToNumString(testLine)
	//fmt.Printf("%v\n", newString)
	//fmt.Printf("expected 12134\n")
	//isNum, num := ptrIsNum(testLine, 0)
	//fmt.Printf("%v\n", isNum)
	//fmt.Printf("%v\n",num)

	sum := 0
	for _, line := range inputLines {
		newLine := convertToNumString(line)
		sum += getNumFromLine(newLine)
	}
	fmt.Printf("%v\n", sum)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
