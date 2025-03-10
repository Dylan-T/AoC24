package main

import (
	"reflect"
	"testing"
)

func TestLoadsDataFromFile(t *testing.T) {
	allLevels := loadAllLevelsFromFile("testinput.txt")

	expectedLevels := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	if !reflect.DeepEqual(allLevels, expectedLevels) {
		t.Errorf("got %v, expected %v", allLevels, expectedLevels)
	}
}

func TestDescending(t *testing.T) {
	result := areLevelsSafe([]int{30, 28, 25, 24, 21, 19, 18, 17})

	if result != true {
		t.Errorf("got %v, expected %v", result, true)
	}
}

func TestRemoveOneWorks(t *testing.T) {

	result := areLevelsSafeWithRemoval([]int{7, 9, 4, 2, 1})

	if result != true {
		t.Errorf("got %v, expected %v", result, true)
	}
}

func TestRemoveLastIndexWorks(t *testing.T) {

	result := areLevelsSafeWithRemoval([]int{7, 5, 4, 2, 8})

	if result != true {
		t.Errorf("got %v, expected %v", result, true)
	}
}

func TestRemoveFirstIndexWorks(t *testing.T) {

	result := areLevelsSafeWithRemoval([]int{2, 5, 4, 2, 1})

	if result != true {
		t.Errorf("got %v, expected %v", result, true)
	}
}
