package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func processInput1(input string) [][]int{
	input_process := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`);
	i := input_process.FindAllStringSubmatch(input, -1);
	var nums [][]int;
	for _, match := range i {
		firstNum, err:= strconv.Atoi(match[1]);
		if err != nil {
			panic("conversion problem for "+ match[1]);
		}
		secondNum, err := strconv.Atoi(match[2]);
		if err != nil {
			panic("conversion problem for "+ match[2]);
		}
		row := []int{firstNum, secondNum};
		nums = append(nums, row);
	}
	return nums;
}

func calc(input [][]int) int{
	sum := 0;
	for _, values := range input {
		sum += values[0] * values[1]
	}
	return sum;
}

func part2(input string) {
	
}

func processInput2(input string) [][]int{
	considerInput := true; //true when do() is seen, false when dont() is seen
	input_process := regexp.MustCompile(`(?:do\(\)|don't\(\))|mul\(([0-9]{1,3}),([0-9]{1,3})\)`); 
	i := input_process.FindAllStringSubmatch(input, -1);
	var nums [][]int;
	for _, match := range i {
		if match[0] == "don't()" {
			considerInput = false;
		} else if match[0] == "do()" {
			considerInput = true;
		}
		if !considerInput {
			continue;
		}
		
		if strings.Contains(match[0], "mul"){
			firstNum, err:= strconv.Atoi(match[1]);
			if err != nil {
				panic("conversion problem for "+ match[1]);
			}
			secondNum, err := strconv.Atoi(match[2]);
			if err != nil {
				panic("conversion problem for "+ match[2]);
			}
			row := []int{firstNum, secondNum};
			nums = append(nums, row);
		}
	}
	return nums;
}


func main(){
	input := readInput();

	//part1
	fmt.Println("Part 1: ", calc(processInput1(input)));
	//part2
	fmt.Println("Part 2: ", calc(processInput2(input)));
}