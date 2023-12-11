package day10

/*
   https://adventofcode.com/2023/day/10
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

/*
	0,1 (north)

-1,0	1,0 (east)

	0,-1
*/
var pipeCons = map[rune][2][2]int{
	'|': {{0, -1}, {0, 1}},
	'-': {{1, 0}, {-1, 0}},
	'L': {{0, -1}, {1, 0}},
	'J': {{0, -1}, {-1, 0}},
	'7': {{0, 1}, {-1, 0}},
	'F': {{0, 1}, {1, 0}},
}

func Execute() {
	start := time.Now()
	file, err := os.Open("../Inputs/2023/Day10_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	m := [][]rune{}

	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)

		l := []rune{}
		for _, tile := range line {
			l = append(l, tile)
		}
		m = append(m, l)
	}

	loopPath := make([][]rune, len(m))
	for i := range loopPath {
		loopPath[i] = make([]rune, len(m[0]))
	}

	var startX int
	var startY int
	for i := 0; i < len(m)-1; i++ {
		for j := 0; j < len(m[i])-1; j++ {
			if m[i][j] == 'S' {
				startX = i
				startY = j
			}
		}
	}

	// Rune 'P' marks the path
	loopPath[startX][startY] = 'P'

	var a int
	if startX+1 < len(m[0]) {
		a = walk(m, startX, startY, startX+1, startY, &loopPath)
	} else {
		a = 0
	}
	var b int
	if startX-1 >= 0 {
		b = walk(m, startX, startY, startX-1, startY, &loopPath)
	} else {
		b = 0
	}
	var c int
	if startY+1 < len(m) {
		c = walk(m, startX, startY, startX, startY+1, &loopPath)
	} else {
		c = 0
	}
	var d int
	if startY-1 >= 0 {
		d = walk(m, startX, startY, startX, startY-1, &loopPath)
	} else {
		d = 0
	}

	dist := max(math.Ceil(float64(a)/2), math.Ceil(float64(b)/2), math.Ceil(float64(c)/2), math.Ceil(float64(d)/2))
	fmt.Println("Part 1: ", dist)

	insideSum := 0
	for i := range m {
		for j := range m[i] {
			inside := false

			i2 := i
			j2 := j

			for i2 < len(m) && j2 < len(m[i2]) {
				if (m[i2][j2] != 'L' && m[i2][j2] != '7') && loopPath[i2][j2] == 'P' {
					inside = !inside
				}
				i2 += 1
				j2 += 1
			}

			if inside && loopPath[i][j] != 'P' {
				insideSum += 1
			}
		}
	}
	fmt.Println("Part 2: ", insideSum)
	elapsed := time.Since(start)
	fmt.Println("Total execution took ", elapsed)
}

func walk(m [][]rune, srcX, srcY, destX int, destY int, loopPath *[][]rune) int {
	if pipeCons[m[destX][destY]][0][0] == srcY-destY && pipeCons[m[destX][destY]][0][1] == srcX-destX {
		(*loopPath)[destX][destY] = 'P'
		return 1 + walk(m, destX, destY, destX+pipeCons[m[destX][destY]][1][1], destY+pipeCons[m[destX][destY]][1][0], loopPath)
	}
	if pipeCons[m[destX][destY]][1][0] == srcY-destY && pipeCons[m[destX][destY]][1][1] == srcX-destX {
		(*loopPath)[destX][destY] = 'P'
		return 1 + walk(m, destX, destY, destX+pipeCons[m[destX][destY]][0][1], destY+pipeCons[m[destX][destY]][0][0], loopPath)
	}
	return 0
}
