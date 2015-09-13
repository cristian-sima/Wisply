package sources

import (
	. "github.com/cristian-sima/Wisply/models/wisply"
	"strconv"
)

type Source struct {
	Id          int
	Name        string
	Url         string
	Description string
}

func (source *Source) Delete() error {
	elememts := []string{
		strconv.Itoa(source.Id),
	}
	_, err := Database.Raw("DELETE from `source` WHERE id=?", elememts).Exec()
	return err
}
