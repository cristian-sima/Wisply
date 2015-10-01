package harvest

import (
	"encoding/xml"
	"fmt"
)

// RemoteRepositoryInterface ... defines the method to be implemented by a standard (remote repository)
type RemoteRepositoryInterface interface {
	Validate()
	HarvestIdentification()
	HarvestFormats()
	HarvestCollections()
	HarvestRecords()
	SetManager(manager *Process)
}

// Message encapsulates the message to communicate with controller
type Message struct {
	Name       string
	Content    string
	Value      interface{}
	Repository int
}

// RemoteRepository represents a remote repository
type RemoteRepository struct {
	Manager *Process
}

// SetManager sets the manager of a current repository
func (repository *RemoteRepository) SetManager(manager *Process) {
	repository.Manager = manager
}

// Keys encapsulate all the dublin core keys
type Keys struct {
	Titles       []string `xml:"title"`
	Creators     []string `xml:"creator"`
	Subjects     []string `xml:"subject"`
	Descriptions []string `xml:"description"`
	Publishers   []string `xml:"publisher"`
	Contributors []string `xml:"contributor"`
	Dates        []string `xml:"date"`
	Types        []string `xml:"type"`
	Formats      []string `xml:"format"`
	Identifiers  []string `xml:"identifier"`
	Sources      []string `xml:"source"`
	Languages    []string `xml:"language"`
	Relations    []string `xml:"relation"`
	Coverages    []string `xml:"coverage"`
	Rights       []string `xml:"rights"`
}

func (repository *RemoteRepository) getKeys(plainText []byte) *Keys {

	keys := Keys{}

	// Unmarshall all the data
	err := xml.Unmarshal(plainText, &keys)
	if err != nil {
		fmt.Println("I got a problem while parsing the Dublin Core format")
		fmt.Println("This is the plain format: <<" + string(plainText) + ">>")
		return &keys
	}

	return &keys
}
