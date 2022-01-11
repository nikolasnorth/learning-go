package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid input or missing filepath. Example usage: go run sum.go /path/to/file")
	}

	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += num
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
