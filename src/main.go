package main

import (
	"fmt"
	"os"
)

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

	switch args[0] {
	case "add":
		fmt.Println("Calling add")
	case "subtract":
		fmt.Println("calling subtract")
	case "multiply":
		fmt.Println("calling multiply")
	case "divide":
		fmt.Println("calling divide")
	default:
		errMsg := `invalid operator %q
valid operators are: 'add', 'subtract', 'multiply', and 'divide'`
		err := fmt.Errorf(errMsg, args[0])
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(args)
}
