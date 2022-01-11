package main

import (
	"bufio"
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

	var nums []int
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err.Error())
		}

		nums = append(nums, num)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(nums)

	outfile, err := os.Create("sortedNumbers.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outfile.Close()

	for _, num := range nums {
		_, err := outfile.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
