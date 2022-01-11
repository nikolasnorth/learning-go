package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ParseInputToFloat(a string, b string) (float64, float64, error) {
	x, err1 := strconv.ParseFloat(a, 64)
	y, err2 := strconv.ParseFloat(b, 64)
	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("could not convert at least one input value to valid float")
	}
	return x, y, nil
}

func Calc(op string, a float64, b float64) (float64, error) {
	switch op {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "divide":
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
		errMsg := `invalid number of arguments (found %d, need 3)
format should follow: [operator] [x] [y]
example:
> add 1 2
> subtract 2 1
> multiply 2 2
> divide 4 2`
		err := fmt.Errorf(errMsg, len(args))
		fmt.Println(err.Error())
		os.Exit(1)
	}

	a, b, err := ParseInputToFloat(args[1], args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result, err := Calc(args[0], a, b)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(result)
}
