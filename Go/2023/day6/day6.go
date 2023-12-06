package day6

/*
   https://adventofcode.com/2023/day/6
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	t   int
	rec int
}

func (r Race) holdButton(ht int) int {
	return (ht * s) * (r.t - ht)
}

/*
I assume the formula: (ht * s) * (t - ht) where:
ht = amount of time button is held
s = speed constant
t = max time of the race
Can be simplified to: ht^2 - t*ht + record > 0
Below is the quadratic formula to get the bounds
*/
func (r Race) computeBounds() (float64, float64) {
	delta := math.Pow(float64(r.t), 2) - 4*float64(r.rec)
	return (float64(r.t) - math.Sqrt(delta)) / 2, (float64(r.t) + math.Sqrt(delta)) / 2
}

const s = 1 //speed constant

func Execute() {
	file, err := os.Open("../Inputs/2023/Day6_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	reNumbers := regexp.MustCompile(`(\d+)`)

	races := []Race{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		numMatches := reNumbers.FindAllStringSubmatch(line, -1)
		for j, match := range numMatches {
			num, _ := strconv.Atoi(match[1])
			if i == 0 {
				race := Race{}
				race.t = num
				races = append(races, race)
			} else {
				races[j].rec = num
			}
		}
		i += 1
	}

	error := 1
	for _, race := range races {
		lowerBound, upperBound := race.computeBounds()
		fmt.Println(math.Ceil(upperBound)-math.Floor(lowerBound)-1, "different ways")
		error *= int(math.Ceil(upperBound) - math.Floor(lowerBound) - 1)
	}

	fmt.Println("Part 1: ", error)

	concatT := ""
	concatRec := ""
	for _, race := range races {
		concatT += strconv.Itoa(race.t)
		concatRec += strconv.Itoa(race.rec)
	}

	_concatT, _ := strconv.Atoi(concatT)
	_concatRec, _ := strconv.Atoi(concatRec)

	concatLowerBound, concatUpperBound := Race{_concatT, _concatRec}.computeBounds()

	fmt.Println("Part 2: ", int(math.Ceil(concatUpperBound)-math.Floor(concatLowerBound)-1))
}
