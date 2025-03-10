package main

import "testing"

const memoryContents = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const memoryContentsDay2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestLoadMemoryToString(t *testing.T) {

	expected := memoryContents

	result := loadMemoryToString("testinput.txt")

	if result != expected {
		t.Errorf("Unexpected result: %s", result)
	}
}

func TestMultiplyNumbers(t *testing.T) {

	expected := 161

	result := multiplyNumbers(memoryContents)

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}

func TestMultiplyNumbersDay2(t *testing.T) {
	expected := 88

	result := multiplyNumbers(memoryContentsDay2)

	if result != expected {
		t.Errorf("Unexpected result: %d", result)
	}
}
