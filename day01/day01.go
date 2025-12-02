package day01

import (
	_ "embed"
	"fmt"
)

func Part1(input []string, dial int) int {
	numberOfTimesHitZero := 0
	for _, line := range input {
		var d string
		var n int
		_, _ = fmt.Sscanf(line, "%1s%d", &d, &n)

		n = n % 100 // since the dial is circular from 0 to 99

		switch d {
		case "L":
			dial = dial - n
		case "R":
			dial = dial + n
		}

		// dial is a circular dial from 0 to 99
		// dial can be bigger than 99 or smaller than 0
		if dial > 99 {
			dial = dial - 100
		}
		if dial < 0 {
			dial = dial + 100
		}

		if dial == 0 {
			numberOfTimesHitZero++
		}

		fmt.Print("After ", line, " dial is at ", dial, "\n")
	}
	return numberOfTimesHitZero
}

func Part2(input []string, dial int) int {
	fmt.Println("Starting with ", dial)
	numberOfTimesHitZero := 0
	for _, line := range input {
		var d string
		var n int
		_, _ = fmt.Sscanf(line, "%1s%d", &d, &n)

		var moves int
		switch d {
		case "L":
			moves = -n
		case "R":
			moves = n
		}

		var crossings int
		dial, crossings = Dial(moves, dial)
		numberOfTimesHitZero += crossings

		if crossings == 0 && dial == 0 {
			numberOfTimesHitZero++
		}

		fmt.Printf("After %s dial is at %d. Number of times passed zero: %d\n", line, dial, numberOfTimesHitZero)
	}
	return numberOfTimesHitZero
}

// dial left or right n steps
// return new dial position and number of times passed zero
func Dial(n int, dial int) (int, int) {
	numberOfTimesPassedZero := 0
	steps := n
	if n < 0 {
		steps = -n
	}

	for i := 0; i < steps; i++ {
		if n > 0 {
			dial++
			if dial > 99 {
				dial = 0
			}
		} else {
			dial--
			if dial < 0 {
				dial = 99
			}
		}

		if dial == 0 {
			numberOfTimesPassedZero++
		}
	}

	return dial, numberOfTimesPassedZero
}
