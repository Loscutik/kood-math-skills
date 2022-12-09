package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"01.kood.tech/git/obudarah/math-skills/statistics"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Wrong numbers of arguments. Needs 2 files")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	strs, err := ReadLines(file)
	if err != nil {
		log.Fatal(err)
	}

	nmbs, err := StrsToNmbrs(strs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Average: %d\nMedian: %d\nVariance: %d\nStandard Deviation: %d\n",
	 statistics.Average(nmbs), statistics.Median(nmbs), statistics.Variance(nmbs), statistics.StandardDeviation(nmbs))
}

/*
	reads file line by line and returns slice of the lines
*/
func ReadLines(file *os.File) (strs []string, err error) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		strs = append(strs, fileScanner.Text())
	}
	err = fileScanner.Err()

	return
}

/*
	converts slice of strings to int. 
	The returned slice of values will contain only successefly convetrted strings from strs.
	If there are invalid digits in the strs, StrsToNmbrs will return an error, which occurred last. 
*/
func StrsToNmbrs(strs []string) (nmbs []int, err error) {
	nmbs = make([]int, len(strs))
	for i, s := range strs {
		nmbs[i], err = strconv.Atoi(s)
	}
	return
}
