package sources

import (
    "github.com/astaxie/beego/orm"
    "errors"
)

type Model struct {
}

func (model *Model) GetAll() []Source {

    orm := orm.NewOrm()

    var list []Source

    orm.Raw("SELECT id, name, url, description FROM source").QueryRows(&list)

    return list
}

func (model *Model) GetSourceById(rawIndex string) (*Source, error) {

    var isValid bool;

    orm := orm.NewOrm()
	source := new(Source)

    isValid = ValidateIndex(rawIndex)

    if !isValid {
        return source, errors.New("Validation invalid")
    }
	error := orm.Raw("SELECT name, url, description FROM source WHERE id = ?", rawIndex).QueryRow(&source)
	return source, error;
}

func (model *Model) ValidateSource(rawData map[string]interface{}) (map[string][]string, error ){

    validationResult := ValidateSourceDetails(rawData)

    if !validationResult.IsValid {
        return validationResult.Errors, errors.New("Validation invalid")
    }
    return nil, nil
}

func (model *Model) UpdateSourceById(sourceId string, rawData map[string]interface{}) error {

    orm := orm.NewOrm()

    stringElements := []string{rawData["name"].(string),
                            rawData["description"].(string),
                            rawData["url"].(string),
							sourceId}

   _, err := orm.Raw("UPDATE `source` SET name=?, description=?, url=? WHERE id=?", stringElements).Exec()

    return err
}

func (model *Model) DeleteSourceById(id string) error {

    orm := orm.NewOrm()

    elememts := []string{id}

    _, err := orm.Raw("DELETE from `source` WHERE id=?", elememts).Exec()

    return err
}



func (model *Model) InsertNewSource(rawData map[string]interface{}) error {

    orm := orm.NewOrm()

    stringElements := []string{rawData["name"].(string),
                            rawData["description"].(string),
                            rawData["url"].(string)}

    _, err := orm.Raw("INSERT INTO `source` (`name`, `description`, `url`) VALUES (?, ?, ?)", stringElements).Exec()

    return err
}
