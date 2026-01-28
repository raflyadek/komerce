package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// sortCharacters()
	psbb()
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

func psbb() {
	reader := bufio.NewReader(os.Stdin)
	//variable to store input 
	var n int
	fmt.Print("Input the number of families: ")
	fmt.Scan(&n)

	fmt.Print("Input the number of members in the family (separated by a space): ")
	//read
	line, _ := reader.ReadString('\n')
	//convert it to array string
	part := strings.Fields(line)
	//validation 
	if len(part) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}
	
	families := make([]int, n)
	//scan through families 
	for i := 0; i < n; i ++ {
		families[i], _ = strconv.Atoi(part[i])
	}


	//sort the families array from low to high
	sort.Ints(families)

	left, right := 0, n-1
	buses := 0

	for left <= right {
		if families[left]+families[right] <= 4 {
			left++
		}
		right--
		buses++
	}

	fmt.Printf("Minimum bus required is: %d\n", buses)
}