package day1

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("../../../Inputs/2023/Day1_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	expected := 56049
	sum := solve(file, false)

	if sum != expected {
		t.Errorf("Solution for the part part 1 returned %d, expected %d", sum, expected)
	}
}

func TestPart2(t *testing.T) {
	file, err := os.Open("../../../Inputs/2023/Day1_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	expected := 54530
	sum := solve(file, true)

	if sum != expected {
		t.Errorf("Solution for the part part 2 returned %d, expected %d", sum, expected)
	}
}
