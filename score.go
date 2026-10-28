package main

import "bytes"

func Score(guess, answer string) [5]string {
	var result [5]string
	answerSlice := []byte(answer)
	for i := range 5 {
		if guess[i] == answer[i] {
			result[i] = "correct"
			answerSlice[i] = '_'
		} else {
			result[i] = "absent"
		}
	}
	for i := range 5 {
		if result[i] == "correct" {
			continue
		}
		j := bytes.IndexByte(answerSlice, guess[i])
		if j != -1 {
			result[i] = "present"
			answerSlice[j] = '_'
		}
	}
	return result

}
