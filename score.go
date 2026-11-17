package main

import (
	"bytes"
	"errors"
)

type Result int

const (
	Absent Result = iota
	Present
	Correct
)
var ErrGuessLength = errors.New("guess must be 5 characters")
var ErrAnswerLength = errors.New("answer must be 5 characters")

func Score(guess, answer string) ([5]Result, error) {
	if len(guess) != 5 {
		return [5]Result{}, ErrGuessLength
	}
	if len(answer) != 5 {
		return [5]Result{}, ErrAnswerLength
	}
	var result [5]Result
	answerSlice := []byte(answer)
	for i := range 5 {
		if guess[i] == answer[i] {
			result[i] = Correct
			answerSlice[i] = '_'
		} 
	}
	for i := range 5 {
		if result[i] == Correct {
			continue
		}
		j := bytes.IndexByte(answerSlice, guess[i])
		if j != -1 {
			result[i] = Present
			answerSlice[j] = '_'
		}
	}
	return result, nil

}
