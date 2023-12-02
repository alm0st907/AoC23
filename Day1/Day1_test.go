package Day1

import "testing"

func TestPart1(t *testing.T) {
	input := "1234"
	expected := 14
	actual := Part1(input)
	if actual != expected {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	input := "1234"
	expected := 14
	actual := Part2(input)
	if actual != expected {
		t.Errorf("Part2 was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPart2Words(t *testing.T) {
	input := "one two three four"
	expected := 14
	actual := Part2(input)
	if actual != expected {
		t.Errorf("Part2 was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPart1Input(t *testing.T) {
	input := "test.txt"
	expected := 142

	actual := RunDayOnePartOne(input)
	if actual != expected {
		t.Errorf("Part1Input was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPart2Input(t *testing.T) {
	input := "test2.txt"
	expected := 281

	actual := RunDayOnePartTwo(input)
	if actual != expected {
		t.Errorf("Part2Input was incorrect, got: %d, want: %d.", actual, expected)
	}
}
