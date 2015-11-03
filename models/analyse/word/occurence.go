package word

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bradfitz/slice"
)

// Occurence maps the number of times a word appears in a text
type Occurence struct {
	Word    string `json:"Word"`
	Counter int    `json:"Counter"`
}

func (occurence *Occurence) increaseCounter() {
	occurence.Counter++
}

// GetWord returns the word
func (occurence Occurence) GetWord() string {
	return occurence.Word
}

// GetCounter returns the number of times
func (occurence Occurence) GetCounter() int {
	return occurence.Counter
}

// Digester transforms a string to a list of occurences
type Digester struct {
	data []*Occurence
}

// GetData returns the processed data
func (occurences Digester) GetData() []*Occurence {
	return occurences.data
}

// GetCounter returns the sum of all the words counters
func (occurences *Digester) GetCounter() int {
	var counter int
	for _, occurence := range occurences.data {
		counter += occurence.GetCounter()
	}
	return counter
}

// GetJSON transforms the list into a json
func (occurences *Digester) GetJSON() string {
	text, _ := json.MarshalIndent(occurences.data, "", "    ")
	return string(text)
}

// Describe shows a short description of the list
func (occurences Digester) Describe() {
	fmt.Println("-----")
	fmt.Println("Words: ")

	for _, occurence := range occurences.data {
		fmt.Print("[" + occurence.GetWord() + " " + strconv.Itoa(occurence.GetCounter()) + "] ")
	}
	fmt.Println("")
	fmt.Println("Number of words: " + strconv.Itoa(occurences.GetNumberOfWords()))
	fmt.Println("Total counter: " + strconv.Itoa(occurences.GetCounter()))
	fmt.Println("-----")

}

// GetNumberOfWords returns the number of words
func (occurences *Digester) GetNumberOfWords() int {
	return len(occurences.data)
}

// AnalyseText adds a text
func (occurences *Digester) AnalyseText(originalText string) {
	words := strings.Split(originalText, " ")
	occurences.AnalyseWords(words)
}

// AnalyseWords inserts an array of words
func (occurences *Digester) AnalyseWords(words []string) {
	var processWord = func(toProcess string) string {
		// in case the last character is '.' we remove it
		sz := len(toProcess)
		if sz > 0 {
			lastChar := string(toProcess[sz-1])
			firstChar := string(toProcess[0])
			rejectedChars := []string{".", ",", "'", ")", "(", ":", ";", "-", "^", "&", "*", "!", "\""}
			for _, rejectedChar := range rejectedChars {
				if lastChar == rejectedChar {
					toProcess = toProcess[:sz-1]
				}
				if firstChar == rejectedChar {
					toProcess = toProcess[0:]
				}
			}

		}
		return strings.TrimSpace(strings.ToLower(toProcess))
	}

	for _, word := range words {
		var exists = false
		for _, occurence := range occurences.data {
			if occurence.GetWord() == processWord(word) {
				exists = true
				occurence.increaseCounter()
			}
		}
		wordToStore := processWord(word)
		if !exists && len(wordToStore) != 0 {
			item := Occurence{
				Word:    wordToStore,
				Counter: 1,
			}
			occurences.data = append(occurences.data, &item)
		}
	}
}

// SortByCounter sorts the list by the counter of the word
func (occurences *Digester) SortByCounter(order string) {
	data := occurences.GetData()
	slice.Sort(data, func(i, j int) bool {
		if order == "ASC" {
			return data[i].GetCounter() < data[j].GetCounter()
		}
		return data[i].GetCounter() > data[j].GetCounter()
	})
}

// NewDigester creates a new list of occurences for words
func NewDigester(text string) Digester {
	list := Digester{}
	list.data = []*Occurence{}
	list.AnalyseText(text)
	return list
}
