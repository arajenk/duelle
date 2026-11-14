package main

import "bytes"

type Result int

const (
	Absent Result = iota
	Present
	Correct
)

func Score(guess, answer string) [5]Result {
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
	return result

}
