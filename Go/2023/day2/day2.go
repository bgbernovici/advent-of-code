package day2

/*
   https://adventofcode.com/2023/day/2
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/bgbernovici/advent-of-code/common"
)

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	r int
	g int
	b int
}

type Expectation struct {
	r int
	g int
	b int
}

func Execute() {
	file, err := os.Open("../Inputs/2023/Day2_.txt")
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

	reGroup := regexp.MustCompile(`(\d+\s*(?:blue|red|green))?(,\s*(\d+\s*(?:blue|red|green)))?(,\s*(\d+\s*(?:blue|red|green)))?\s*(;|\n|$)`)
	reId := regexp.MustCompile(`(\d+):`)
	games := []Game{}
	for _, line := range lines {
		game := Game{}
		game.sets = []Set{}
		matchesId := reId.FindAllStringSubmatch(line, -1)
		for _, match := range matchesId {
			game.id, err = strconv.Atoi(match[1])
			if err != nil {
				fmt.Printf("Could not convert string digit to int")
			}
		}

		matchesGroup := reGroup.FindAllString(line, -1)

		for _, match := range matchesGroup {
			generateSet(&game, match)
		}

		games = append(games, game)
	}

	// Part 1
	expected := Expectation{12, 13, 14}
	sum := 0
	for _, game := range games {
		possible := true
		for _, set := range game.sets {
			if !possible {
				break
			}
			if !isSetSmallerOrEqualThanExpectated(set, expected) {
				possible = false
			}
		}
		if possible {
			sum += game.id
		}
	}
	fmt.Println("Part 1: ", sum)

	//Part 2
	sum = 0
	for _, game := range games {
		maxR, maxG, maxB := 0, 0, 0
		for _, set := range game.sets {
			maxR = common.MaxInt(set.r, maxR)
			maxG = common.MaxInt(set.g, maxG)
			maxB = common.MaxInt(set.b, maxB)
		}
		powerOfFewest := maxR * maxG * maxB
		sum += powerOfFewest
	}
	fmt.Println("Part 2: ", sum)
}

var reRed = regexp.MustCompile(`(\d+)\s+red`)
var reGreen = regexp.MustCompile(`(\d+)\s+green`)
var rBlue = regexp.MustCompile(`(\d+)\s+blue`)
var reDigits = regexp.MustCompile(`(\d+)`)

func generateSet(game *Game, match string) {
	set := Set{}

	mRed := reRed.FindAllStringSubmatch(match, -1)
	mGreen := reGreen.FindAllStringSubmatch(match, -1)
	mBlue := rBlue.FindAllStringSubmatch(match, -1)
	for _, match := range mRed {
		mDigits := reDigits.FindAllStringSubmatch(match[1], -1)
		for _, match := range mDigits {
			set.r, _ = strconv.Atoi(match[1])
		}
	}
	for _, match := range mGreen {
		mDigits := reDigits.FindAllStringSubmatch(match[1], -1)
		for _, match := range mDigits {
			set.g, _ = strconv.Atoi(match[1])
		}
	}
	for _, match := range mBlue {
		mDigits := reDigits.FindAllStringSubmatch(match[1], -1)
		for _, match := range mDigits {
			set.b, _ = strconv.Atoi(match[1])
		}
	}
	game.sets = append(game.sets, set)
}

func isSetSmallerOrEqualThanExpectated(set Set, expected Expectation) bool {
	return set.r <= expected.r && set.b <= expected.b && set.g <= expected.g
}
