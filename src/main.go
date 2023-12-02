package main

import (
	"advent-of-code-2023/days"
	"fmt"
	"log"
	"os"
)

var solutions = map[string]func()(int, error){
	"1A": days.Day01A,
	"1B": days.Day01B,
}

func main() {
	day := os.Args[2]
	fmt.Println("Running solution for day: \"" + day + "\"")

	solution, ok := solutions[day];

	if !ok { log.Fatal("No day specified - or - day specified has not been implemented") } 

	result, err := solution()

	if err != nil { log.Fatal(err) }

	fmt.Printf("The answer is: %d\n", result)
}
