package oai

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request represents a request URL and query string to an OAI-PMH service
type Request struct {
	BaseURL         string
	Set             string
	MetadataPrefix  string
	Verb            string
	Identifier      string
	ResumptionToken string
	From            string
	Until           string
}

// HarvestIdentifiers arvest the identifiers of a complete OAI set
// call the identifier callback function for each Header
func (request *Request) HarvestIdentifiers(callback func(*Header)) {
	request.Verb = "ListIdentifiers"
	request.Harvest(func(resp *Response) {
		headers := resp.ListIdentifiers.Headers
		for _, header := range headers {
			callback(&header)
		}
	}, func(resp *Response) {
		fmt.Println(" I finished with all!")
	})
}

// HarvestRecords harvest the identifiers of a complete OAI set
// call the identifier callback function for each Header
func (request *Request) HarvestRecords(callbackRecord func(*Record), callbackFinish func(*Response)) {
	request.Verb = "ListRecords"
	request.Harvest(func(resp *Response) {
		records := resp.ListRecords.Records
		for _, record := range records {
			callbackRecord(&record)
		}
	}, callbackFinish)
}

// ChannelHarvestIdentifiers harvest the identifiers of a complete OAI set
// send a reference of each Header to a channel
func (request *Request) ChannelHarvestIdentifiers(channels []chan *Header) {
	request.Verb = "ListIdentifiers"
	request.Harvest(func(resp *Response) {
		headers := resp.ListIdentifiers.Headers
		i := 0
		for _, header := range headers {
			channels[i] <- &header
			i++
			if i == len(channels) {
				i = 0
			}
		}

		// If there is no more resumption token, send nil to all
		// the channels to signal the harvest is done
		hasResumptionToken, _ := resp.ObtainResumptionToken()
		if !hasResumptionToken {
			for _, channel := range channels {
				channel <- nil
			}
		}
	}, func(resp *Response) {
	})
}

// Harvest perform a harvest of a complete OAI set, or simply one request
// call the batchCallback function argument with the OAI responses
// The finish callback is called when there is no more things
func (request *Request) Harvest(batchCallback, finishCallback func(*Response)) {

	fmt.Println("------------Starting---------------")
	fmt.Println("<--> OAI: Start harvesting request...")

	// Use Perform to get the OAI response

	response := request.Perform()

	fmt.Println(" --> OAI Request: The request has been loaded.")

	// Execute the callback function with the response
	batchCallback(response)

	// Check for a resumptionToken
	hasResumptionToken, resumptionToken := response.ObtainResumptionToken()

	// Harvest further if there is a resumption token
	if hasResumptionToken {
		fmt.Println(" --> OAI Request: Has resumption")
		request.Set = ""
		request.MetadataPrefix = ""
		request.From = ""
		request.ResumptionToken = resumptionToken

		fmt.Println("------------Finished---------------")

		request.Harvest(batchCallback, finishCallback)
	} else {
		fmt.Println(" <--> OAI: Request: Does not have resumption!")
		fmt.Println("<--> OAI: Finished request")
		fmt.Println("-------------Finished--------------")
		finishCallback(response)
	}
}

// Perform an HTTP GET request using the OAI Requests fields
// and return an OAI Response reference
func (request *Request) Perform() (response *Response) {
	url := request.GetFullURL()

	fmt.Println("<--> OAI: Performing request from:")
	fmt.Println(url)
	fmt.Println("<--> OAI: Please wait...")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read all the data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshall all the data
	err = xml.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(" <--> Problem when unmarshal XML")
		fmt.Println(err)
		panic(err)
	}

	return response
}

// GetFullURL returns the full URL address for the request
// scheme:[BaseURL][?][parameter=value[&]]
// see http://www.w3.org/Addressing/URL/url-spec.txt
func (request *Request) GetFullURL() string {
	parameters := []string{}

	addParameter := func(name, value string) {
		if value != "" {
			parameters = append(parameters, name+"="+value)
		}
	}

	addParameter("verb", request.Verb)
	addParameter("set", request.Set)
	addParameter("metadataPrefix", request.MetadataPrefix)
	addParameter("resumptionToken", request.ResumptionToken)
	addParameter("identifier", request.Identifier)
	addParameter("from", request.From)
	addParameter("until", request.Until)

	query := "?" + strings.Join(parameters, "&")
	URL := request.BaseURL + query

	return URL
}

// ObtainResumptionToken determine the resumption token in this Response
func (resp *Response) ObtainResumptionToken() (hasResumptionToken bool, resumptionToken string) {
	hasResumptionToken = false
	resumptionToken = ""
	if resp == nil {
		return
	}

	// First attempt to obtain a resumption token from a ListIdentifiers response
	resumptionToken = resp.ListIdentifiers.ResumptionToken

	// Then attempt to obtain a resumption token from a ListRecords response
	if resumptionToken == "" {
		resumptionToken = resp.ListRecords.ResumptionToken
	}

	// If a non-empty resumption token turned up it can safely inferred that...
	if resumptionToken != "" {
		hasResumptionToken = true
	}

	return
}
