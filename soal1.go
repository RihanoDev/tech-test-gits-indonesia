package main

import "fmt"

func generateSequence(num int) []int {
	if num < 1 {
		return nil
	}

	sequence := make([]int, num)
	for i := 0; i < num; i++ {
		sequence[i] = (i * (i + 1) / 2) + 1
	}

	return sequence
}

func printSequence(seq []int) {
	if seq == nil {
		fmt.Println("Invalid input")
		return
	}

	for i, val := range seq {
		if i > 0 {
			fmt.Print("-")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

func main() {
	var input int
	fmt.Print("Type input: ")
	fmt.Scanln(&input)
	printSequence(generateSequence(input))
}
