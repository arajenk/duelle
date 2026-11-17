package main

import (
	"testing"
	"errors"
)

func TestScoreBasic(t *testing.T) {
	got, err := Score("APPLE", "EAGLE")
	if err != nil{
		t.Fatalf("unexpected error: %v", err)
	}
	want := [5]Result{Present, Absent, Absent, Correct, Correct}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestScoreDuplicates(t *testing.T) {
	got, err := Score("EAGLE", "APPLE")
	if err != nil{
		t.Fatalf("unexpected err: %v", err)
	}
	want := [5]Result{Absent, Present, Absent, Correct, Correct}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestScoreShortGuess(t *testing.T) {
	_, err := Score("APP", "EAGLE")
	if !errors.Is(err, ErrGuessLength){
		t.Errorf("got %v, wanted %v", err, ErrGuessLength)
	}
}