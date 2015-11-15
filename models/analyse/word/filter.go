package word

import "strings"

// Filter returns the filter
type Filter struct {
	original *Digester
	data     *Digester
}

// GetOriginal returns the original digester
func (filter *Filter) GetOriginal() *Digester {
	return filter.original
}

// GetData returns the processed digester
func (filter *Filter) GetData() *Digester {
	return filter.data
}

// Process takes a list of strings and removes the words from the original digester
// it creates a new digester with the new set of words
func (filter *Filter) Process(list []string) {
	allowedOccurences := []*Occurence{}
	for _, originalOccurence := range (*filter.original).GetData() {
		reject := false
		currentWord := originalOccurence.GetWord()
		for _, filterWord := range list {
			if strings.ToLower(currentWord) == strings.ToLower(filterWord) {
				reject = true
			}
		}
		if !reject {
			allowedOccurences = append(allowedOccurences, originalOccurence)
		}
	}
	item := Digester{
		data: allowedOccurences,
	}
	filter.data = &item
}

// PrepositionFilter is a Filter which rejects all the prepositions
type PrepositionFilter struct {
	Filter
}

func (filter *PrepositionFilter) process() {
	filter.Filter.Process(prepositions)
}

// NewPrepositionsFilter creates a Filter which removes all the prepositions
func NewPrepositionsFilter(list *Digester) *PrepositionFilter {
	filter := PrepositionFilter{
		Filter: NewFilter(list),
	}
	filter.process()
	return &filter
}

// ConjunctionsFilter is a Filter which removes all the conjunctions
type ConjunctionsFilter struct {
	Filter
}

func (filter *ConjunctionsFilter) process() {
	filter.Filter.Process(GetFilterList("conjunctions"))
}

// NewConjunctionsFilter creates a Filter which removes all the conjuections
func NewConjunctionsFilter(list *Digester) *ConjunctionsFilter {
	filter := ConjunctionsFilter{
		Filter: NewFilter(list),
	}
	filter.process()
	return &filter
}

// PronounsFilter is a Filter which removes all the pronouns
type PronounsFilter struct {
	Filter
}

func (filter *PronounsFilter) process() {
	filter.Filter.Process(GetFilterList("pronouns"))
}

// NewPronounsFilter creates a Filter which removes all the pronouns
func NewPronounsFilter(list *Digester) *PronounsFilter {
	Filter := PronounsFilter{
		Filter: NewFilter(list),
	}
	Filter.process()
	return &Filter
}

// ArticleFilter removes the occurences which contain articles
type ArticleFilter struct {
	Filter
}

func (filter *ArticleFilter) process() {
	filter.Filter.Process(GetFilterList("articles"))
}

// NewArticleFilter creates a new article Filter.
// This Filter removes all the occures which are articles
func NewArticleFilter(list *Digester) *ArticleFilter {
	Filter := ArticleFilter{
		Filter: NewFilter(list),
	}
	Filter.process()
	return &Filter
}

// GrammarFilter removes the occurences which contain articles
type GrammarFilter struct {
	Filter
}

func (grammar *GrammarFilter) process() {
	list := []string{}
	list = append(list, articles...)
	list = append(list, conjunctions...)
	list = append(list, redundantWords...)
	list = append(list, pronouns...)
	list = append(list, prepositions...)
	list = append(list, education...)
	grammar.Filter.Process(list)
}

// NewGrammarFilter creates a Filter which removes: articles, conjuctions, prepositions,
// redundant words and pronouns
func NewGrammarFilter(list *Digester) *GrammarFilter {
	Filter := GrammarFilter{
		Filter: NewFilter(list),
	}
	Filter.process()
	return &Filter
}

// NewFilter creates a new Filter
func NewFilter(list *Digester) Filter {
	return Filter{
		original: list,
	}
}
