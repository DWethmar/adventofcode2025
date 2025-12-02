package day01

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

//go:embed input_test.txt
var inputTest string

//go:embed input_1.txt
var input1 string

func Test_Part1(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		r := Part1(strings.Split(input1, "\n"), 50)
		fmt.Printf("Answer: %d\n", r)
	})

	t.Run("solve test", func(t *testing.T) {
		r := Part1(strings.Split(inputTest, "\n"), 50)
		fmt.Printf("Answer: %d\n", r)
	})
}

func Test_Part2(t *testing.T) {
	t.Run("solve", func(t *testing.T) {
		r := Part2(strings.Split(input1, "\n"), 50)
		fmt.Printf("Answer: %d\n", r)
	})

	t.Run("solve test", func(t *testing.T) {
		r := Part2(strings.Split(inputTest, "\n"), 50)
		fmt.Printf("Answer: %d\n", r)
	})
}

func TestDial(t *testing.T) {
	t.Run("add r1000", func(t *testing.T) {
		newDial, crossings := Dial(1000, 50)
		if newDial != 50 {
			t.Errorf("Dial(1000, 50) = %d, want %d", newDial, 50)
		}
		if crossings != 10 {
			t.Errorf("Dial(1000, 50) crossings = %d, want %d", crossings, 10)
		}
	})
}
