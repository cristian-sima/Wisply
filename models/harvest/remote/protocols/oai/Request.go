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

// Verify sends a Indentify request and returns the response
func (request *Request) Verify() ([]byte, error) {
	request.verb = "Identify"
	return request.get()
}

// Get an HTTP request, reads the body of the request and returns it
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
func (request *Request) Parse(body []byte) (*Response, error) {
	response := &Response{}
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("There was a problem while trying parsing the XML:<br /> " + err.Error())
	}
	return response, nil
}

// ---
//
// // HarvestIdentifiers arvest the identifiers of a complete OAI set
// // call the identifier callback function for each Header
// func (request *Request) HarvestIdentifiers(callback func(*Header)) {
// 	request.Verb = "ListIdentifiers"
// 	request.Harvest(func(resp *Response) {
// 		headers := resp.ListIdentifiers.Headers
// 		for _, header := range headers {
// 			callback(&header)
// 		}
// 	}, func(resp *Response) {
// 		fmt.Println(" I finished with all!")
// 	})
// }
//
// // HarvestRecords harvest the identifiers of a complete OAI set
// // call the identifier callback function for each Header
// func (request *Request) HarvestRecords(callbackRecord func(*Record), callbackFinish func(*Response)) {
// 	request.Verb = "ListRecords"
// 	request.Harvest(func(resp *Response) {
// 		records := resp.ListRecords.Records
// 		for _, record := range records {
// 			callbackRecord(&record)
// 		}
// 	}, callbackFinish)
// }
//
// // ChannelHarvestIdentifiers harvest the identifiers of a complete OAI set
// // send a reference of each Header to a channel
// func (request *Request) ChannelHarvestIdentifiers(channels []chan *Header) {
// 	request.Verb = "ListIdentifiers"
// 	request.Harvest(func(resp *Response) {
// 		headers := resp.ListIdentifiers.Headers
// 		i := 0
// 		for _, header := range headers {
// 			channels[i] <- &header
// 			i++
// 			if i == len(channels) {
// 				i = 0
// 			}
// 		}
//
// 		// If there is no more resumption token, send nil to all
// 		// the channels to signal the harvest is done
// 		hasResumptionToken := resp.HasResumptionToken()
// 		if !hasResumptionToken {
// 			for _, channel := range channels {
// 				channel <- nil
// 			}
// 		}
// 	}, func(resp *Response) {
// 	})
// }
//
// // Harvest perform a harvest of a complete OAI set, or simply one request
// // call the batchCallback function argument with the OAI responses
// // The finish callback is called when there is no more things
// func (request *Request) Harvest(batchCallback, finishCallback func(*Response)) {
//
// 	fmt.Println("------------Starting---------------")
// 	fmt.Println("<--> OAI: Start harvesting request...")
//
// 	// Use Perform to get the OAI response
//
// 	response, _ := request.Perform()
//
// 	fmt.Println(" --> OAI Request: The request has been loaded.")
//
// 	// Execute the callback function with the response
// 	batchCallback(response)
//
// 	// Check for a resumptionToken
// 	hasResumptionToken, resumptionToken := response.ObtainResumptionToken()
//
// 	// Harvest further if there is a resumption token
// 	if hasResumptionToken {
// 		fmt.Println(" --> OAI Request: Has resumption")
// 		request.Set = ""
// 		request.MetadataPrefix = ""
// 		request.From = ""
// 		request.ResumptionToken = resumptionToken
//
// 		fmt.Println("------------Finished---------------")
//
// 		request.Harvest(batchCallback, finishCallback)
// 	} else {
// 		fmt.Println(" <--> OAI: Request: Does not have resumption!")
// 		fmt.Println("<--> OAI: Finished request")
// 		fmt.Println("-------------Finished--------------")
// 		finishCallback(response)
// 	}
// }
