package main

import (
	day01 "advent-of-code-2023/day/01"
	day02 "advent-of-code-2023/day/02"
	day03 "advent-of-code-2023/day/03"
	day04 "advent-of-code-2023/day/04"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

var solutions = map[string]func([]string) (int, error){
	"1A": day01.RunA,
	"1B": day01.RunB,
	"2A": day02.RunA,
	"2B": day02.RunB,
	"3A": day03.RunA,
	"3B": day03.RunB,
	"4A": day04.RunA,
	"4B": day04.RunB,
}

func main() {
	day := os.Args[2]
	fmt.Println("Running solution for day: \"" + day + "\"")

	solution, ok := solutions[day]

	if !ok {
		log.Fatal("No day specified - or - day specified has not been implemented")
	}

	result, err := solution(nil)

	if err != nil {
		log.Fatal(err)
	}

	copiedMessage := "  -  ✗ Could not copy to clipboard"
	if clipboard.Unsupported == false {
		clipboard.WriteAll(strconv.Itoa(result))
		copiedMessage = "  -  ✓ Copied to clipboard!"
	}

	fmt.Printf("\n> The answer is: %d%v\n\n", result, copiedMessage)
}
