package oai

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request encapsulates the information to get information using OAI-PMH protocol
type Request struct {
	BaseURL         string
	Set             string
	MetadataPrefix  string
	verb            string
	Identifier      string
	ResumptionToken string
	From            string
	Until           string
}

// Identify returns the identification details of the repository
func (request *Request) Identify() ([]byte, error) {
	return request.getVerb("Identify")
}

func (request *Request) getVerb(verb string) ([]byte, error) {
	request.verb = verb
	return request.get()
}

// IsValidResponse parses a content and returns the error
func (request *Request) IsValidResponse(content []byte) error {
	_, err := request.Parse(content)
	if err != nil {
		return err
	}
	return nil
}

// Parse returns the content of the remote repository
func (request *Request) Parse(content []byte) (*Response, error) {
	response, err := request.getResponse(content)
	if err != nil {
		return response, err
	}
	return response, err
}

// It performs an HTTP request, reads the body of the request and returns it
func (request *Request) get() ([]byte, error) {
	var (
		content []byte
	)
	url := request.getFullURL()

	resp, err := http.Get(url)
	if err != nil {
		return content, errors.New("There was a problem while trying to connect to that address: <br />" + err.Error())
	}
	defer resp.Body.Close()

	// Read all the data
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return content, errors.New("There was a problem while trying to read the content of the page: <br />" + err.Error())
	}
	return content, nil
}

// GetFullURL returns the full URL address for the request
// scheme:[BaseURL][?][parameter=value[&]]
// see http://www.w3.org/Addressing/URL/url-spec.txt
func (request *Request) getFullURL() string {
	parameters := []string{}

	addParameter := func(name, value string) {
		if value != "" {
			parameters = append(parameters, name+"="+value)
		}
	}
	addParameter("verb", request.verb)
	addParameter("set", request.Set)
	addParameter("metadataPrefix", request.MetadataPrefix)
	addParameter("identifier", request.Identifier)
	addParameter("from", request.From)
	addParameter("until", request.Until)
	addParameter("resumptionToken", request.ResumptionToken)

	query := "?" + strings.Join(parameters, "&")
	URL := request.BaseURL + query

	return URL
}

// Parse an array of bytes into an OAI Response and returns the Response
func (request *Request) getResponse(body []byte) (*Response, error) {
	response := &Response{}
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("There was a problem while trying parsing the XML:<br /> " + err.Error())
	}
	return response, nil
}
