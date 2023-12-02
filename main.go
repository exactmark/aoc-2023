package main

import (
	"aoc-2023/Day01"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func readInputFile(filename string) []string {
	var returnStrings []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		returnStrings = append(returnStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return returnStrings
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	currentDay := "01"

	//inputLines := readInputFile("./Day" + currentDay + "/test_input.txt")
	inputLines := readInputFile("./Day" + currentDay + "/input.txt")

	start := time.Now()

	Day01.Solve(inputLines)

	elapsed := time.Since(start)
	fmt.Printf("solve took %s\n", elapsed)
}
