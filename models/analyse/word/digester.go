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

// GetMostProeminent returns a digester with the most prominent words
// The most prominent words are those with the counter >= (totalCount/distict_nr_of_words)
func (digester Digester) GetMostProeminent(factor int) *Digester {
	totalCount := digester.GetCounter() * factor
	totalWords := len(digester.data)
	if totalWords == 0 {
		return NewDigester("")
	}
	threshold := int(totalCount / totalWords)
	return digester.filterDataByCounter(threshold)
}

// GetMostRelevant returns a digester with the most relevant words
// The most relevant words are those with the counter >= (topCounter/2)
func (digester Digester) GetMostRelevant() *Digester {
	max := digester.GetTopOccurrence().GetCounter()
	threshold := int(max / 4)
	return digester.filterDataByCounter(threshold)
}

// GetTopOccurrence returns the occurence with the highest counter
func (digester Digester) GetTopOccurrence() *Occurence {
	max := &Occurence{
		Counter: -1,
	}
	for _, occurence := range digester.data {
		if occurence.GetCounter() > max.GetCounter() {
			max = occurence
		}
	}
	return max
}

// GetTop returns the top x words
func (digester Digester) GetTop(threshold int) *Digester {
	digester.SortByCounter("DESC")
	list := []*Occurence{}
	number := 0
	for _, occurence := range digester.data {
		if number == threshold {
			return &Digester{
				data: list,
			}
		}
		list = append(list, occurence)
		number++
	}
	return &Digester{
		data: list,
	}
}

func (digester Digester) filterDataByCounter(threshold int) *Digester {
	list := []*Occurence{}
	for _, occurence := range digester.data {
		if occurence.GetCounter() >= threshold {
			list = append(list, occurence)
		}
	}
	return &Digester{
		data: list,
	}
}

// GetData returns the processed data
func (digester Digester) GetData() []*Occurence {
	return digester.data
}

// GetCounter returns the sum of all the words counters
func (digester *Digester) GetCounter() int {
	var counter int
	for _, occurence := range digester.data {
		counter += occurence.GetCounter()
	}
	return counter
}

// GetJSON transforms the list into a json
func (digester *Digester) GetJSON() string {
	text, _ := json.MarshalIndent(digester.data, "", "    ")
	return string(text)
}

// GetPlainJSON returns a plain version of json
func (digester *Digester) GetPlainJSON() string {
	text, _ := json.Marshal(digester.data)
	return string(text)
}

// RemoveOccurence removes a word from the digester
func (digester *Digester) RemoveOccurence(word string) {
	for index, occurence := range digester.GetData() {
		if occurence.GetWord() == strings.ToLower(word) {
			digester.data = append(digester.data[:index], digester.data[index+1:]...)
		}
	}
}

// Combine combines two digesters into one
func (digester *Digester) Combine(secondDigester *Digester) *Digester {
	combination := &Digester{
		data: digester.GetData(),
	}
	for _, occurenceSecond := range secondDigester.GetData() {
		exists := false
		var theOccurence *Occurence
		for _, occurenceFirst := range combination.GetData() {
			if occurenceFirst.GetWord() == occurenceSecond.GetWord() {
				exists = true
				theOccurence = occurenceFirst
			}
		}
		if exists {
			theOccurence.Counter = theOccurence.Counter + occurenceSecond.GetCounter()
		} else {
			item := Occurence{
				Word:    occurenceSecond.GetWord(),
				Counter: occurenceSecond.GetCounter(),
			}
			combination.data = append(combination.data, &item)
		}
	}
	return combination
}

// Describe shows a short description of the list
func (digester Digester) Describe() {
	fmt.Println("-----")
	fmt.Println("Words: ")
	for _, occurence := range digester.data {
		fmt.Print("[" + occurence.GetWord() + " " + strconv.Itoa(occurence.GetCounter()) + "] ")
	}
	fmt.Println("")
	fmt.Println("Number of words: " + strconv.Itoa(digester.GetNumberOfWords()))
	fmt.Println("Total counter: " + strconv.Itoa(digester.GetCounter()))
	fmt.Println("-----")
}

// GetNumberOfWords returns the number of words
func (digester *Digester) GetNumberOfWords() int {
	return len(digester.data)
}

// AnalyseText adds a text
func (digester *Digester) AnalyseText(originalText string) {
	originalText = strings.Replace(originalText, "\n", " ", -1)
	originalText = strings.Replace(originalText, "\r", " ", -1)
	words := strings.Split(originalText, " ")
	digester.AnalyseWords(words)
}

// GetString returns the data as a string/text, separated by space
func (digester *Digester) GetString() string {
	buffer := ""
	for _, occurence := range digester.data {
		buffer += occurence.GetWord() + " "
	}
	return buffer
}

// AnalyseWords inserts an array of words
func (digester *Digester) AnalyseWords(words []string) {
	var processWord = func(toProcess string) string {
		toProcess = strings.ToLower(toProcess)
		sz := len(toProcess)
		if sz > 0 {
			rejectedChars := []string{"“", "‘", "’", "”", ".", ",", "‘", "’", "'", ")", "(", ":", ";", "-", "^", "&", "*", "!", "\""}
			tryAgain := true
			for tryAgain && (len(toProcess) > 0) {
				tryAgain = false
				toProcess = strings.TrimSpace(toProcess)
				sz = len(toProcess)
				lastChar := string(toProcess[sz-1])
				firstChar := string(toProcess[0])
				for _, rejectedChar := range rejectedChars {
					size := len(toProcess)
					if lastChar == rejectedChar {
						toProcess = toProcess[0 : size-1]
						tryAgain = true
					}
					if firstChar == rejectedChar {
						if len(toProcess) < 2 {
							toProcess = ""
						} else {
							toProcess = toProcess[1:]
						}
						tryAgain = true
					}
				}
			}
		}

		// reject numbers
		_, err := strconv.Atoi(toProcess)
		if err == nil {
			toProcess = ""
		}
		return toProcess
	}

	for _, word := range words {
		var exists = false
		for _, occurence := range digester.data {
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
			digester.data = append(digester.data, &item)
		}
	}
}

// SortByCounter sorts the list by the counter of the word
func (digester *Digester) SortByCounter(order string) {
	data := digester.GetData()
	slice.Sort(data, func(i, j int) bool {
		if order == "ASC" {
			return data[i].GetCounter() < data[j].GetCounter()
		}
		return data[i].GetCounter() > data[j].GetCounter()
	})
}

// NewDigester creates a new list of occurences for words
func NewDigester(text string) *Digester {
	digester := Digester{}
	digester.data = []*Occurence{}
	digester.AnalyseText(text)
	return &digester
}

// NewDigesterFromJSON transforms a json string to a digester
func NewDigesterFromJSON(text string) *Digester {
	var list []*Occurence
	err := json.Unmarshal([]byte(text), &list)
	if err != nil {
		fmt.Println(err)
	}
	return &Digester{
		data: list,
	}
}
