package main

import (
	"testing"
)

func TestScoreBasic(t *testing.T) {
	got := Score("APPLE", "EAGLE")
	want := [5]string{"present", "absent", "absent", "correct", "correct"}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestScoreDuplicates(t *testing.T) {
	got := Score("EAGLE", "APPLE")
	want := [5]string{"absent", "present", "absent", "correct", "correct"}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
