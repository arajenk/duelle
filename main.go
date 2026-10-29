package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

var words = []string{"crane", "apple", "eagle", "plant", "ghost", "brick", "flame", "sword", "crown", "storm"}

func main() {
	// pick a random answer from the word_list
	n := rand.IntN(len(words))
	answer := words[n]
	scanner := bufio.NewScanner(os.Stdin)
	// take a guess from the terminal
	count := 0
	for count < 6 {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.ToLower(guess)
		guess = strings.TrimSpace(guess)
		if !isValid(guess) {
			fmt.Println("Invalid word.")
			continue
		}
		count++
		if guess == answer {
			fmt.Println("Correct.")
			fmt.Println(Score(guess,answer))
			break
		} else {
			fmt.Printf("Wrong! %d out of 6 attempts!\n", count)
			fmt.Println(Score(guess,answer))
			if count == 6 {
				fmt.Println("Out of guesses!")
				fmt.Printf("The word was: %s\n", answer)
			}
		}
	}


}
func isValid(word string) bool {
	if !slices.Contains(words, word) {
		return false
	} else if len(word) != 5 {
		return false
	}
	return true
}
