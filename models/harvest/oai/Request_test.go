package oai

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRequestURL(t *testing.T) {
	request := Request{
		BaseURL: "http://eprints.uwe.ac.uk/",
	}
	Convey("get just address", t, func() {
		So(request.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?")
	})

	Convey("get request URL with verb", t, func() {
		request2 := Request{
			BaseURL: "http://eprints.uwe.ac.uk/",
			Verb:    "identify",
		}

		So(request2.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?verb=identify")
	})

	Convey("get request URL with set", t, func() {
		request3 := Request{
			BaseURL: "http://eprints.uwe.ac.uk/",
			Set:     "history",
		}
		So(request3.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?set=history")
	})

	Convey("get request URL with prefix", t, func() {
		request4 := Request{
			BaseURL:        "http://eprints.uwe.ac.uk/",
			MetadataPrefix: "prefix",
		}
		So(request4.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?metadataPrefix=prefix")
	})

	Convey("get request URL with prefix, verb and set", t, func() {
		request5 := Request{
			BaseURL:        "http://eprints.uwe.ac.uk/",
			Set:            "history",
			MetadataPrefix: "prefix",
			Verb:           "identify",
		}
		So(request5.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?verb=identify&set=history&metadataPrefix=prefix")
	})

	Convey("get request URL with resumptionToken", t, func() {
		request6 := Request{
			BaseURL:         "http://eprints.uwe.ac.uk/",
			ResumptionToken: "metadataPrefix%3Doai_dc%26offset%3D275",
		}
		So(request6.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?resumptionToken=metadataPrefix%3Doai_dc%26offset%3D275")
	})

	Convey("get request URL with identifier", t, func() {
		request7 := Request{
			BaseURL:    "http://eprints.uwe.ac.uk/",
			Identifier: "79",
		}
		So(request7.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?identifier=79")
	})

	Convey("get request URL with from", t, func() {
		request8 := Request{
			BaseURL: "http://eprints.uwe.ac.uk/",
			From:    "2009",
		}
		So(request8.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?from=2009")
	})

	Convey("get request URL with until", t, func() {
		request9 := Request{
			BaseURL: "http://eprints.uwe.ac.uk/",
			Until:   "2015",
		}
		So(request9.GetFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?until=2015")
	})

	Convey("get request full URL: prefix, verb, set, identifier, from, until and resumptionTokenidentifier", t, func() {
		request10 := Request{
			BaseURL:         "http://eprints.uwe.ac.uk/",
			Set:             "history",
			MetadataPrefix:  "prefix",
			Verb:            "identify",
			Identifier:      "79",
			From:            "2009",
			Until:           "2015",
			ResumptionToken: "metadataPrefix%3Doai_dc%26offset%3D275",
		}
		expectedURL := "http://eprints.uwe.ac.uk/?verb=identify&set=history&metadataPrefix=prefix&identifier=79&from=2009&until=2015&resumptionToken=metadataPrefix%3Doai_dc%26offset%3D275"
		So(request10.GetFullURL(), ShouldEqual, expectedURL)
	})
}
