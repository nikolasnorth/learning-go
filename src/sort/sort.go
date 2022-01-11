package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid input or missing filepath. Example usage: go run sort.go /path/to/file")
	}

	filepath := os.Args[1]
	infile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer infile.Close()

	var nums []float64
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		nums = append(nums, num)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	sort.Float64s(nums)

	outfile, err := os.Create("sortedNumbers.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outfile.Close()

	for _, num := range nums {
		_, err := outfile.WriteString(fmt.Sprintf("%.1f\n", num))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
