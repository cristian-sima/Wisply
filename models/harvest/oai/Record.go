package oai

// Record is metadata expressed in a single format.
//
// A record is returned in an XML-encoded byte stream in response to an OAI-PMH request for metadata from an item.
// A record is identified unambiguously by the combination of the unique identifier of the item from which the record is available, the metadataPrefix identifying the metadata format of the record, and the datestamp of the record.
// The XML-encoding of records is organized into the following parts: header, metadata, about
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#Record
type Record struct {
	Header   Header   `xml:"header"`
	Metadata Metadata `xml:"metadata"`
	About    About    `xml:"about"`
}
