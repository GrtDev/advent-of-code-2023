package day07

import (
	"advent-of-code-2023/utils"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type HandType int16

const (
	Unknown HandType = iota
	HighCard
	OnePair
	TwoPair
	TheeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var orderA = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var orderB = [...]string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

type hand struct {
	handType HandType
	cards    []string
	bid      int
}

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadLines(inputFile)
	} else {
		return utils.ReadLines("./day/07/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	hands := parseHands(getInput(inputFile), false)
	sortedHands := sortHandsLowToHigh(hands, orderA[:])

	totalWinnings := 0

	for i, hand := range sortedHands {
		totalWinnings += hand.bid * (i + 1)
	}

	return totalWinnings, nil
}

func RunB(inputFile string) (int, error) {
	hands := parseHands(getInput(inputFile), true)
	sortedHands := sortHandsLowToHigh(hands, orderB[:])

	totalWinnings := 0

	for i, hand := range sortedHands {
		totalWinnings += hand.bid * (i + 1)
	}

	return totalWinnings, nil

}

func parseHands(input []string, useJoker bool) []hand {
	hands := []hand{}
	for _, line := range input {
		handString := strings.Split(line, " ")
		bid, _ := strconv.Atoi(handString[1])
		cards := strings.Split(handString[0], "")
		hands = append(hands, hand{
			cards:    cards,
			bid:      bid,
			handType: getHandType(cards, useJoker),
		})
	}
	return hands
}

func getHandType(cards []string, useJoker bool) HandType {
	cardCount := make(map[string]int)

	jokers := 0

	for _, card := range cards {
		if useJoker && card == "J" {
			jokers++
			continue
		}

		if _, ok := cardCount[card]; ok {
			cardCount[card]++
		} else {
			cardCount[card] = 1
		}
	}

	values := utils.Values(cardCount)

	if useJoker && jokers > 0 {
		if jokers >= 5 {
			return FiveOfAKind
		}
		sort.Ints(values)
		values[len(values)-1] += jokers
	}

	switch len(values) {
	case 1:
		return FiveOfAKind
	case 2:
		if funk.MaxInt(values) == 4 {
			return FourOfAKind
		}
		return FullHouse
	case 3:
		if funk.MaxInt(values) == 3 {
			return TheeOfAKind
		}
		return TwoPair
	case 4:
		return OnePair
	case 5:
		return HighCard
	}
	return Unknown
}

func sortHandsLowToHigh(hands []hand, order []string) []hand {
	sort.Slice(hands, func(i, j int) bool {
		a, b := hands[i], hands[j]

		if a.handType == b.handType {
			for i, cardA := range a.cards {
				if cardA == b.cards[i] {
					continue
				}
				indexA := funk.IndexOfString(order, cardA)
				indexB := funk.IndexOfString(order, b.cards[i])

				if indexA < 0 || indexB < 0 {
					log.Fatalf("Failed to find the index of %v or %v", cardA, b.cards[i])
				}
				return indexA < indexB
			}
		}

		return a.handType < b.handType
	})

	return hands
}
