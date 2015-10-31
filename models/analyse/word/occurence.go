package word

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bradfitz/slice"
)

// Occurence maps the number of times a word appears in a text
type Occurence struct {
	word    string
	counter int
}

func (occurence *Occurence) increaseCounter() {
	occurence.counter++
}

// GetWord returns the word
func (occurence Occurence) GetWord() string {
	return occurence.word
}

// GetCounter returns the number of times
func (occurence Occurence) GetCounter() int {
	return occurence.counter
}

// OccurenceList transforms a string to a list of occurences
type OccurenceList struct {
	originalText string
	data         []*Occurence
}

// GetData returns the processed data
func (occurences OccurenceList) GetData() []*Occurence {
	return occurences.data
}

// GetCounter returns the sum of all the words counters
func (occurences *OccurenceList) GetCounter() int {
	var counter int
	for _, occurence := range occurences.data {
		counter += occurence.GetCounter()
	}
	return counter
}

// GetOriginalText returns the original text of the occurence
func (occurences OccurenceList) GetOriginalText() string {
	return occurences.originalText
}

// Describe shows a short description of the list
func (occurences OccurenceList) Describe() {
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
func (occurences *OccurenceList) GetNumberOfWords() int {
	return len(occurences.data)
}

func (occurences *OccurenceList) process() {

	var processWord = func(toProcess string) string {
		// in case the last character is '.' we remove it
		sz := len(toProcess)
		if sz > 0 && toProcess[sz-1] == '.' {
			toProcess = toProcess[:sz-1]
		}
		return strings.TrimSpace(strings.ToLower(toProcess))
	}

	words := strings.Split(occurences.originalText, " ")
	occurences.data = []*Occurence{}
	for _, word := range words {
		var exists = false
		for _, occurence := range occurences.data {
			if occurence.word == processWord(word) {
				exists = true
				occurence.increaseCounter()
			}
		}
		wordToStore := processWord(word)
		if !exists && len(wordToStore) != 0 {
			item := Occurence{
				word:    wordToStore,
				counter: 1,
			}
			occurences.data = append(occurences.data, &item)
		}
	}
}

// SortByCounter sorts the list by the counter of the word
func (occurences *OccurenceList) SortByCounter(order string) {
	data := occurences.GetData()
	slice.Sort(data, func(i, j int) bool {
		if order == "ASC" {
			return data[i].GetCounter() < data[j].GetCounter()
		}
		return data[i].GetCounter() > data[j].GetCounter()
	})
}

// NewOccurencesList creates a new list of occurences for words
func NewOccurencesList(text string) OccurenceList {
	list := OccurenceList{
		originalText: text,
	}
	list.process()
	return list
}
