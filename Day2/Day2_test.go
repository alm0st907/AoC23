package Day2

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 2285
	actual := Day2Part1Driver("input.txt")
	if actual != expected {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 77021
	actual := Day2Part2Driver("input.txt")
	if actual != expected {
		t.Errorf("Part2 was incorrect, got: %d, want: %d.", actual, expected)
	}
}
