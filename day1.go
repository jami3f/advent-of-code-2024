package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getIntPairs() ([]int, []int) {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	firsts := make([]int, 0)
	seconds := make([]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numPair := scanner.Text()
		splitPair := strings.Split(numPair, "   ")
		first, second := splitPair[0], splitPair[1]
		first = strings.TrimSpace(first)
		second = strings.TrimSpace(second)
		firstInt, err := strconv.Atoi(first)
		check(err)
		secondInt, err := strconv.Atoi(second)
		check(err)
		firsts = append(firsts, firstInt)
		seconds = append(seconds, secondInt)
	}
	return firsts, seconds
}

func calculateDistance() {
	firsts, seconds := getIntPairs()
	sort.Ints(firsts)
	sort.Ints(seconds)
	distance := 0
	for i := range firsts {
		distance += int(math.Abs(float64(firsts[i] - seconds[i])))
	}
	fmt.Print(distance)
}

func calculateSimilarityScore() {
	firsts, seconds := getIntPairs()
	similarityScore := 0
	for i := range firsts {
		value := firsts[i]
		count := 0
		for j := range seconds {
			if value == seconds[j] {
				count++
			}
		}
		similarityScore += value * count
	}
	println(similarityScore)
}

func main() {
	calculateSimilarityScore()
}