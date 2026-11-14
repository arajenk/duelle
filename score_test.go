package main

import (
	"testing"
)

func TestScoreBasic(t *testing.T) {
	got := Score("APPLE", "EAGLE")
	want := [5]Result{Present, Absent, Absent, Correct, Correct}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestScoreDuplicates(t *testing.T) {
	got := Score("EAGLE", "APPLE")
	want := [5]Result{Absent, Present, Absent, Correct, Correct}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
