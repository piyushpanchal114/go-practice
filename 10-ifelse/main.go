package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("If else in Go lang")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
	}

	fmt.Println("str", str)
	num, err := strconv.Atoi(strings.TrimSpace(str))

	fmt.Println("num, err", num, err)
	if err == nil {
		fmt.Println("Number is", num)
	} else {
		fmt.Println("Invalid number")
	}
	if num < 10 {
		fmt.Println("Num is less than 10")
	} else if num > 10 {
		fmt.Println("Num is greater than 10")
	} else {
		fmt.Println("Exactly 10")
	}

	if num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

	if x := 3; x < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("greater than 10")
	}

}
