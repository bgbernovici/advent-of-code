package day11_manhattan

/*
   https://adventofcode.com/2023/day/11
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

func Execute() {
	start := time.Now()
	file, err := os.Open("../Inputs/2023/Day11_.txt")
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

	tempM := [][]rune{}

	for i := range m {
		tempM = append(tempM, m[i])
		allDots := true
		for _, c := range m[i] {
			if c != '.' {
				allDots = false
			}
		}
		if allDots {
			for k, c := range m[i] {
				if c == '.' {
					m[i][k] = 'M'
				}
			}
		}
	}

	m = make([][]rune, len(tempM))
	copy(m, tempM)

	for i := range m[0] {
		allDots := true
		for j := range m {
			if m[j][i] != '.' && m[j][i] != 'M' {
				allDots = false
			}
		}
		if allDots {
			for j := range m {
				m[j][i] = 'M'
			}
		}
	}

	m = tempM
	galaxies := [][2]int{}
	for i := range m {
		for j := range m[i] {
			if tempM[i][j] == '#' {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}

	sumChPart1 := make(chan int)
	sumChPart2 := make(chan int)
	var wgPart1 sync.WaitGroup
	var wgPart2 sync.WaitGroup
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			wgPart1.Add(1)
			go manhattan(&m, [2]int{galaxies[i][0], galaxies[i][1]}, [2]int{galaxies[j][0], galaxies[j][1]}, 2, sumChPart1, &wgPart1)
			wgPart2.Add(1)
			go manhattan(&m, [2]int{galaxies[i][0], galaxies[i][1]}, [2]int{galaxies[j][0], galaxies[j][1]}, 1000000, sumChPart2, &wgPart2)
		}
	}

	go func() {
		wgPart1.Wait()
		close(sumChPart1)
		wgPart2.Wait()
		close(sumChPart2)
	}()

	sumPart1 := 0
	for s := range sumChPart1 {
		sumPart1 += s
	}

	sumPart2 := 0
	for s := range sumChPart2 {
		sumPart2 += s
	}

	fmt.Println("Part 1: ", sumPart1)
	fmt.Println("Part 2: ", sumPart2)

	elapsed := time.Since(start)
	fmt.Println("Total execution took ", elapsed)
}

func manhattan(m *[][]rune, source [2]int, dest [2]int, expandFactor int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	mX := 0
	mY := 0
	if source[0] < dest[0] {
		for i := source[0]; i <= dest[0]; i++ {
			if (*m)[i][source[1]] == 'M' {
				mX += 1
			}
		}
	} else if source[0] > dest[0] {
		for i := source[0]; i >= dest[0]; i-- {
			if (*m)[i][source[1]] == 'M' {
				mX += 1
			}
		}
	}
	if source[1] < dest[1] {
		for i := source[1]; i <= dest[1]; i++ {
			if (*m)[source[0]][i] == 'M' {
				mY += 1
			}
		}
	} else if source[1] > dest[1] {
		for i := source[1]; i >= dest[1]; i-- {
			if (*m)[source[0]][i] == 'M' {
				mY += 1
			}
		}
	}
	dist := int(math.Abs(float64(dest[0]-source[0])) + float64(mX*(expandFactor-1)) + math.Abs(float64(dest[1]-source[1])) + float64(mY*(expandFactor-1)))
	ch <- int(dist)
}
