package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Operator string

const (
	Add      Operator = "add"
	Subtract Operator = "subtract"
	Multiply Operator = "multiply"
	Divide   Operator = "divide"
)

func ParseInputToFloat(a string, b string) (float64, float64, error) {
	x, err1 := strconv.ParseFloat(a, 64)
	y, err2 := strconv.ParseFloat(b, 64)
	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("could not convert at least one input value to valid float")
	}
	return x, y, nil
}

func Calc(op Operator, a float64, b float64) (float64, error) {
	switch op {
	case Add:
		return a + b, nil
	case Subtract:
		return a - b, nil
	case Multiply:
		return a * b, nil
	case Divide:
		if b == 0 {
			return 0, errors.New("divisor cannot be zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatal("invalid number of arguments")
	}

	var op Operator
	switch args[0] {
	case "add":
		op = Add
	case "subtract":
		op = Subtract
	case "multiply":
		op = Multiply
	case "divide":
		op = Divide
	default:
		log.Fatalf("invalid operator '%s'", args[0])
	}

	a, b, err := ParseInputToFloat(args[1], args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result, err := Calc(op, a, b)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(result)
}
