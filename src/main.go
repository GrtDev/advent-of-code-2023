package main

import (
	day01 "advent-of-code-2023/day/01"
	day02 "advent-of-code-2023/day/02"
	day03 "advent-of-code-2023/day/03"
	day04 "advent-of-code-2023/day/04"
	day05 "advent-of-code-2023/day/05"
	day06 "advent-of-code-2023/day/06"
	day07 "advent-of-code-2023/day/07"
	day08 "advent-of-code-2023/day/08"
	day09 "advent-of-code-2023/day/09"
	day10 "advent-of-code-2023/day/10"
	day11 "advent-of-code-2023/day/11"
	day12 "advent-of-code-2023/day/12"
	day13 "advent-of-code-2023/day/13"
	day14 "advent-of-code-2023/day/14"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
)

var solutions = map[string]func(string) (int, error){
	"1A":  day01.RunA,
	"1B":  day01.RunB,
	"2A":  day02.RunA,
	"2B":  day02.RunB,
	"3A":  day03.RunA,
	"3B":  day03.RunB,
	"4A":  day04.RunA,
	"4B":  day04.RunB,
	"5A":  day05.RunA,
	"5B":  day05.RunB,
	"6A":  day06.RunA,
	"6B":  day06.RunB,
	"7A":  day07.RunA,
	"7B":  day07.RunB,
	"8A":  day08.RunA,
	"8B":  day08.RunB,
	"9A":  day09.RunA,
	"9B":  day09.RunB,
	"10A": day10.RunA,
	"10B": day10.RunB,
	"11A": day11.RunA,
	"11B": day11.RunB,
	"12A": day12.RunA,
	"12B": day12.RunB,
	"13A": day13.RunA,
	"13B": day13.RunB,
	"14A": day14.RunA,
	"14B": day14.RunB,
}

func main() {
	day := os.Args[2]
	fmt.Println("Running solution for day: \"" + day + "\"")

	solution, ok := solutions[day]

	if !ok {
		log.Fatal("No day specified - or - day specified has not been implemented")
	}

	start := time.Now()
	result, err := solution("")
	elapsed := time.Since(start)

	if err != nil {
		log.Fatal(err)
	}

	copiedMessage := "  -  ✗ Could not copy to clipboard"
	if clipboard.Unsupported == false {
		clipboard.WriteAll(strconv.Itoa(result))
		copiedMessage = "  -  ✓ Copied to clipboard!"
	}

	fmt.Printf("\n> Finished in: %v", elapsed.String())
	fmt.Printf("\n> The answer is: %d%v\n\n", result, copiedMessage)
}
