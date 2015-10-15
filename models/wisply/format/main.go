package format

// MIME holds the translation of MIME types to a Wisply ones
var MIME = map[string]string{
	"application/pdf":               "PDF",
	"application/vnd.ms-powerpoint": "Power Point",
	"application/vnd.ms-excel":      "Excel Sheet",
	"application/msword":            "Word Document",
	"application/vnd.edshare-link":  "Internal link",
	"application/x-shockwave-flash": "Flash",
	"application/zip":               "Zip archive",
	"application/rdf+xml ":          "RDF and XML file",
	"application/rtf":               "Rich Text Format",

	"audio/mpeg":  "Recording",
	"audio/mp3":   "Recording",
	"audio/x-wav": "Recording",

	"image/png":     "Picture",
	"image/jpg":     "Picture",
	"image/jpeg":    "Picture",
	"image/gif":     "Picture",
	"image/bmp":     "Picture",
	"image/svg+xml": "Vector Picture",

	"text/html":     "HTML page",
	"text/plain":    "Text file",
	"text/xml":      "XML file",
	"text/x-matlab": "Matlab file",
	"text/x-tex":    "TeX file",

	"video/quicktime": "Video",
	"video/mp4":       "Video",
	"video/mpeg":      "Video",
	"video/x-ms-wmv":  "Video",
	"video/x-msvideo": "Video",
	"video/x-m4v":     "Video",
	"video/x-flv":     "Video",
	"video/x-ms-asx":  "Video",

	"other": "Other file",
}

// ConvertFormats converts a list of formats
func ConvertFormats(values []string) []string {
	newFormats := []string{}
	for _, format := range values {
		newFormat := ConvertFormat(format)
		newFormats = append(newFormats, newFormat)
	}
	return newFormats
}

// ConvertFormat converts a raw format to a nice one
func ConvertFormat(formatRaw string) string {
	ret := formatRaw
	wisplyFormat, exists := MIME[formatRaw]
	if exists {
		ret = wisplyFormat
	}
	return ret
}
