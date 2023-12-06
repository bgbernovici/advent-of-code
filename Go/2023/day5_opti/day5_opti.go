package day5_opti

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
	from              string
	to                string
	ranges            []Range
	transformedRanges []TransformedRange
}

type TransformedRange struct {
	from   int
	to     int
	offset int
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

	// Transform
	// Work only with source range, set as offset the diff between dest and source
	for i, d := range almanac.mappings {
		for _, r := range d.ranges {
			tm := TransformedRange{from: r.srcStart, to: r.srcStart + r.length - 1, offset: r.destStart - r.srcStart}
			almanac.mappings[i].transformedRanges = append(almanac.mappings[i].transformedRanges, tm)

		}
	}

	// Part 1
	start := time.Now()
	minLoc := math.MaxInt
	for _, s := range almanac.seeds {
		loc := s
		for _, d := range almanac.mappings {
			for _, tm := range d.transformedRanges {
				if tm.from <= loc && loc <= tm.to {
					loc += tm.offset
					break
				}
			}
		}
		if loc < minLoc {
			minLoc = loc
		}
	}

	fmt.Println("Lowest location number: ", minLoc)
	elapsed := time.Since(start)
	fmt.Println("Part 1 execution took ", elapsed)

	// Part 2
	start = time.Now()
	minLoc = math.MaxInt
	// Split seeds
	newSeeds := []int{}
	for i := 0; i < len(almanac.seeds)-1; i += 2 {
		for j := almanac.seeds[i]; j < almanac.seeds[i]+almanac.seeds[i+1]; j++ {
			newSeeds = append(newSeeds, j)
		}
		for _, s := range newSeeds {
			loc := s
			for _, d := range almanac.mappings {
				for _, tm := range d.transformedRanges {
					if tm.from <= loc && loc <= tm.to {
						loc += tm.offset
						break
					}
				}
			}
			if loc < minLoc {
				minLoc = loc
			}
		}
		fmt.Println("Finished ranges: ", almanac.seeds[i], almanac.seeds[i+1])
		newSeeds = []int{}
	}

	fmt.Println("Lowest location number: ", minLoc)
	elapsed = time.Since(start)
	fmt.Println("Part 2 execution took ", elapsed)
}
