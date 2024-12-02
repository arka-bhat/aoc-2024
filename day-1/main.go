package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readInput() string{
	file, err := os.Open("input.txt");
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

func main(){
	input := readInput();

	split_input := regexp.MustCompile(`[\n]`).Split(input, -1) ;
	pairsplitter := regexp.MustCompile(" +");

	var numsleft []int;
	var numsright []int;
	for a, n := range split_input{
		var numpair = pairsplitter.Split(n, -1);

		left, err := strconv.Atoi(strings.TrimSpace(numpair[0]));
		if err != nil {
			panic("conversion problem at " + strconv.Itoa(a) + " num: " + n);
		}
		numsleft = append(numsleft, left);

		right, err := strconv.Atoi(strings.TrimSpace(numpair[1]));
		if err != nil {
			panic("conversion problem at " + strconv.Itoa(a) + " num: " + n);
		}
		numsright = append(numsright, right);
	}

	//part 1
	part1(numsleft, numsright);

	//part 2
	part2(numsleft, numsright);
}

func part1(numsleft []int, numsright []int) {
	sort.Ints(numsleft);
	sort.Ints(numsright);
	sum := 0.0;
	for i := 0; i < len(numsleft); i++ {
		diff := math.Abs(float64(numsleft[i]) - float64(numsright[i]));
		sum += diff;
	}
	fmt.Printf("%.0f\n", sum);
}

func part2(numsleft []int, numsright []int) {
	freqMap := make(map[int]int);
	for _, num := range numsright{
		freqMap[num]++
	}
	var ans uint64 = 0;
	for _, num := range numsleft{
		freq := freqMap[num];
		if freq != 0 {
			ans += uint64(freq * num);
		}
	}
	fmt.Print(ans);
}