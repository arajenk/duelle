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
	
	n := rand.IntN(len(words))
	answer := words[n]
	scanner := bufio.NewScanner(os.Stdin)
	
	count := 0
	for count < 6 {
		fmt.Print("Enter your guess: ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error reading input:", err)
				os.Exit(1)
			}
			fmt.Printf("\nNo more input. The word was: %s\n", answer)
			return
		}
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
