package day5

/*
   https://adventofcode.com/2023/day/5
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Almanac struct {
	seeds    []int
	mappings []Domain
}

type Range struct {
	destStart int
	srcStart  int
	length    int
}

type Domain struct {
	from   string
	to     string
	ranges []Range
}

func Execute() {
	file, err := os.Open("../Inputs/2023/Day5_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	almanac := Almanac{[]int{}, []Domain{}}

	seedsStr := strings.Replace(lines[0], "seeds: ", "", -1)
	seedsStrVec := strings.Split(seedsStr, " ")
	for _, s := range seedsStrVec {
		n, _ := strconv.Atoi(s)
		almanac.seeds = append(almanac.seeds, n)
	}

	i := 1
	for i < len(lines) {
		if strings.Contains(lines[i], "to") {
			fp := strings.Split(lines[i], "-to-")
			sp := strings.Split(fp[1], " ")
			domain := Domain{from: fp[0], to: sp[0], ranges: []Range{}}
			almanac.mappings = append(almanac.mappings, domain)
		} else {
			rangesStr := strings.Split(lines[i], " ")
			if len(rangesStr) == 3 {
				r := Range{}
				for i, s := range rangesStr {
					n, _ := strconv.Atoi(s)
					if i == 0 {
						r.destStart = n
					} else if i == 1 {
						r.srcStart = n
					} else if i == 2 {
						r.length = n
					}
				}
				almanac.mappings[len(almanac.mappings)-1].ranges = append(almanac.mappings[len(almanac.mappings)-1].ranges, r)
			}
		}
		i += 1
	}

	start := time.Now()
	minLocGlobal := computeMinimumLocation(&almanac)
	fmt.Println("Part 1: ", minLocGlobal)
	elapsed := time.Since(start)
	fmt.Println("Part 1 execution took ", elapsed)

	start = time.Now()
	minLocGlobal = math.MaxInt
	newSeeds := []int{}
	for i := 0; i < len(almanac.seeds)-1; i += 2 {
		for j := almanac.seeds[i]; j < almanac.seeds[i]+almanac.seeds[i+1]; j++ {
			newSeeds = append(newSeeds, j)
		}
		tempAlmanac := Almanac{newSeeds, almanac.mappings}
		minLoc := computeMinimumLocation(&tempAlmanac)
		if minLoc < minLocGlobal {
			minLocGlobal = minLoc
		}
		fmt.Println("Finished ranges: ", almanac.seeds[i], almanac.seeds[i+1])
		newSeeds = []int{}
	}
	almanac.seeds = newSeeds

	fmt.Println("Part 2: ", minLocGlobal)
	elapsed = time.Since(start)
	fmt.Println("Part 2 execution took ", elapsed)
}

func computeMinimumLocation(almanac *Almanac) int {
	minLocGlobal := math.MaxInt
	for _, seed := range almanac.seeds {
		next := []int{seed}
		for _, domain := range almanac.mappings {
			temp := []int{}
			for _, n := range next {
				for _, r := range domain.ranges {
					candidate := r.destStart + r.length - ((r.srcStart + r.length) - n)
					if n >= r.srcStart && n < r.srcStart+r.length {
						temp = append(temp, candidate)
					}
				}
				if len(temp) == 0 {
					temp = append(temp, n)
				}
			}
			next = temp
		}
		minLoc := math.MaxInt
		for _, n := range next {
			if n < minLoc {
				minLoc = n
			}
		}
		if minLoc < minLocGlobal {
			minLocGlobal = minLoc
		}
	}
	return minLocGlobal
}
