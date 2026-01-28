package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sortCharacters()
}

func sortCharacters() {
	//variable to store word
	var resultConstant string
	var resultVowels string

	//variable to store vowel
	vowel := "aiueo"

	fmt.Print("Input one line of words (S): ")
	//scan input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	//sanitize input
	inputTrim := strings.Join(strings.Fields(input), "")
	inputLower := strings.ToLower(inputTrim)
	
	//loop the input and search if its in the vowel
	for i := 0; i < len(inputLower); i++ {
		char := string(inputLower[i])

		if strings.Contains(vowel, char) {
			resultVowels += char
		} else {
			resultConstant += char
		}
		
	}

	fmt.Println("\nVowel Characters:", resultVowels)
	fmt.Println("Constant Characters:", resultConstant)
}