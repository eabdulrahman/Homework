package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func countWords(words []string) map[string]int {

	wordc := make(map[string]int) //The map declaration

	for _, i := range words {
		_, ok1 := wordc[i]
		if ok1 { //if the word exists in the map, increase the value by 1
			wordc[i] += 1
		} else { //it's a new word add it to the map
			wordc[i] = 1
		}
	}
	return wordc
}

func reportResults(word_map map[string]int) {

	for k, v := range word_map {
		fmt.Printf("%v : %v\n", k, v) //print the map keys and values
	}
}

func main() {

	var fileName string
	fmt.Print("Type the name of the TXT file you would like to count the number of words in it: ")
	fmt.Scan(&fileName)                   //The name of file being entered by the user
	str, err := ioutil.ReadFile(fileName) //Reading the file
	if err != nil {
		log.Fatal("unable to read file: %v")
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	txt := strings.ToLower(strings.TrimSpace(string(str))) //convert the text to lowercase and remove whitespaces
	str2 := reg.ReplaceAllString(txt, " ")

	words := strings.Split(str2, " ") //converting the words into a slice of strings

	counted_words := countWords(words) //calling the word counter function

	reportResults(counted_words) //calling the function that prints out how often each word occurs in the file
}
