package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Open(wd + "/numbers.txt")
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
