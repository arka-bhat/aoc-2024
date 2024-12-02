package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var input_file = "input.txt";

func readInput() string{
	file, err := os.Open(input_file);
	if err != nil {
		panic("Error opening input file");
	}

	defer file.Close();
	content, err := io.ReadAll(file);
	if err != nil {
		panic("Error opening input file");
	}
	return string(content)
}

func processInput(input string) [][]int{
	input_split := regexp.MustCompile("[\n]").Split(input, -1);
	split_by_space := regexp.MustCompile(" +");
	var input_nums [][]int;
	for _, line := range input_split {
		line_split := split_by_space.Split(line, -1);
		var nums []int;
		for _, n := range line_split {
			num, err := strconv.Atoi(n);
			if err != nil {
				panic("error in conversion for " + n);
			}
			nums = append(nums, num);
		}
		input_nums = append(input_nums, nums);
	}
	return input_nums;
}

func abs(a int) int {
    if a < 0 {
        return -1 * a
    }
    return a
}

func strictlyIncreasing(row []int) bool {
    for i := 1; i < len(row); i++ {
        if row[i-1] >= row[i] {
            return false
        }
    }
    return true
}
 
func strictlyDecreasing(row []int) bool {
    for i := 1; i < len(row); i++ {
        if row[i] >= row[i-1] {
            return false
        }
    }
    return true
}

func isSafe(row []int, ) bool {
	isSafe := true;

	for i := 1; i < len(row); i++ {
		diff := abs(row[i-1] - row[i]);
		if diff > 3 || diff < 1 {
            isSafe = false
        }
	}
	if !strictlyIncreasing(row) && !strictlyDecreasing(row){
		isSafe = false;
	}

	return isSafe;
}

func canBeSafe(row []int) bool {
	for i := range row {
		rowCopy := make([]int, len(row));
		copy(rowCopy, row);
		rowCopy = append(rowCopy[:i], rowCopy[i+1:]...);
		if isSafe(rowCopy) {
			return true;
		}
	}
	return false;
}

func part1(inputMatrix [][]int) {
	safeCount := 0;
	for _, row := range inputMatrix {
		if isSafe(row) {
			safeCount++;
		}
	}
	fmt.Println("Part 1: ", safeCount);
}

func part2(inputMatrix [][]int) {
	safeCount := 0;
	for _, row := range inputMatrix {
		if isSafe(row) || canBeSafe(row) {
			safeCount++;
		}
	}
	fmt.Println("Part 2: ", safeCount);
}

func main(){
	input := readInput();
	inputMatrix := processInput(input);

	//part1
	part1(inputMatrix);
	//part2
	part2(inputMatrix);
}