package word

import "strings"

// Occurencer ... defines the basic sets for an occurence list
type Occurencer interface {
	GetData() []*Occurence
	GetOriginalText() string
	Describe()
}

type filter struct {
	original *OccurenceList
	data     *OccurenceList
}

func (filter *filter) GetOriginal() *OccurenceList {
	return filter.original
}

func (filter *filter) GetData() *OccurenceList {
	return filter.data
}

func (filter *filter) process(list []string) {
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
	item := OccurenceList{
		data: allowedOccurences,
	}
	filter.data = &item
}

// PrepositionFilter is a filter which rejects all the prepositions
type PrepositionFilter struct {
	filter
}

func (filter *PrepositionFilter) process() {
	filter.filter.process(prepositions)
}

// NewPrepositionsFilter creates a filter which removes all the prepositions
func NewPrepositionsFilter(list *OccurenceList) *PrepositionFilter {
	filter := PrepositionFilter{
		filter: newFilter(list),
	}
	filter.process()
	return &filter
}

// ConjunctionsFilter is a filter which removes all the conjunctions
type ConjunctionsFilter struct {
	filter
}

func (filter *ConjunctionsFilter) process() {
	filter.filter.process(conjunctions)
}

// NewConjunctionsFilter creates a filter which removes all the conjuections
func NewConjunctionsFilter(list *OccurenceList) *ConjunctionsFilter {
	filter := ConjunctionsFilter{
		filter: newFilter(list),
	}
	filter.process()
	return &filter
}

// PronounsFilter is a filter which removes all the pronouns
type PronounsFilter struct {
	filter
}

func (filter *PronounsFilter) process() {
	filter.filter.process(pronouns)
}

// NewPronounsFilter creates a filter which removes all the pronouns
func NewPronounsFilter(list *OccurenceList) *PronounsFilter {
	filter := PronounsFilter{
		filter: newFilter(list),
	}
	filter.process()
	return &filter
}

// ArticleFilter removes the occurences which contain articles
type ArticleFilter struct {
	filter
}

func (filter *ArticleFilter) process() {
	filter.filter.process(articles)
}

// NewArticleFilter creates a new article filter.
// This filter removes all the occures which are articles
func NewArticleFilter(list *OccurenceList) *ArticleFilter {
	filter := ArticleFilter{
		filter: newFilter(list),
	}
	filter.process()
	return &filter
}

// GrammarFilter removes the occurences which contain articles
type GrammarFilter struct {
	filter
}

func (grammar *GrammarFilter) process() {
	list := []string{}
	list = append(list, articles...)
	list = append(list, conjunctions...)
	list = append(list, redundantWords...)
	list = append(list, pronouns...)
	list = append(list, prepositions...)
	list = append(list, education...)
	grammar.filter.process(list)
}

// NewGrammarFilter creates a filter which removes: articles, conjuctions, prepositions,
// redundant words and pronouns
func NewGrammarFilter(list *OccurenceList) *GrammarFilter {
	filter := GrammarFilter{
		filter: newFilter(list),
	}
	filter.process()
	return &filter
}

func newFilter(list *OccurenceList) filter {
	return filter{
		original: list,
	}
}
