// Package main provides a simple text analysis program.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// countWordsAndCharacters takes a text input and returns the count of characters and words.
func countWordsAndCharacters(text string) (int, int) {
	characterCount := len(text)
	wordsCount := 0
	inWord := false

	for _, char := range text {
		if char == ' ' {
			if inWord {
				wordsCount++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	// If there is only 1 word or to count the last word
	if inWord {
		wordsCount++
	}

	return characterCount, wordsCount
}

// typePercentage takes a text input and returns the percentage of vowels, consonants, and special characters.
func typePercentage(text string) (vowelCount, consonantCount, specialCount int) {
	vowelCount = 0
	consonantCount = 0
	specialCount = 0
	characterCount := len(text)

	for _, char := range text {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
				vowelCount++
			} else {
				consonantCount++
			}
		} else {
			specialCount++
		}
	}

	return (vowelCount * 100) / characterCount, consonantCount * 100 / characterCount, specialCount * 100 / characterCount
}

func main() {
	// Make a scanner using bufio, create a new scanner object from os.Stdin for input
	scanner := bufio.NewScanner(os.Stdin)

	// Ask the user for input text
	fmt.Printf("\nEnter the text: ")
	scanner.Scan()
	var inputText string
	inputText = scanner.Text()

	fmt.Println("----------------------------------------------------------")
	fmt.Println("Word and Character count")

	// Call the function to count characters and words
	characterCount, wordsCount := countWordsAndCharacters(inputText)

	// Call the function to return the percentage of vowels, consonants, special characters
	vowelCount, consonantCount, specialCount := typePercentage(inputText)

	// Print results
	fmt.Printf("\nCharacter Count: %d\n", characterCount)
	fmt.Printf("Word Count: %d\n", wordsCount)
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Vowels and Consonants percentage")

	fmt.Printf("\nVowels: %d%%", vowelCount)
	fmt.Printf("\nConsonants: %d%%", consonantCount)
	fmt.Printf("\nOther characters: %d%%\n", specialCount)
	fmt.Println("----------------------------------------------------------")
}
