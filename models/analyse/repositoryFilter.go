package analyse

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/analyse/word"
	"github.com/cristian-sima/Wisply/models/repository"
)

// RepositoryAnalyseFilter is the filter which removes words when they are analysed
// for EPrints repositories
type RepositoryAnalyseFilter struct {
	word.Filter
	repository *repository.Repository
}

func (filter *RepositoryAnalyseFilter) process() {
	list := filter.repository.GetFilter().GetStructure().Analyse.Reject

	original := filter.Filter.GetOriginal()
	data := original.GetData()

	for _, occurence := range data {
		if strings.HasPrefix(occurence.GetWord(), "keywords:") {
			words := strings.Split(occurence.GetWord(), "keywords:")
			if len(words) > 1 {
				occurence.Word = words[1]
			}
		}
	}
	filter.Filter.Process(list)
}

// NewRepositoryAnalyseFilter creates a filter for the repository
func NewRepositoryAnalyseFilter(repository *repository.Repository, digester *word.Digester) *RepositoryAnalyseFilter {

	filter := &RepositoryAnalyseFilter{
		Filter:     word.NewFilter(digester),
		repository: repository,
	}
	filter.process()
	return filter
}
