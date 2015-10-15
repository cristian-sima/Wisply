package format

var formats = map[string]map[string]string{
	"application/pdf": map[string]string{
		"image": "img",
		"title": "PDF",
	},
}

// ShowFormat the label of a format
func ShowFormat(formatRaw string) string {
	ret := formatRaw
	format, exists := formats[formatRaw]
	if exists {
		ret = format["title"]
	}
	return ret
}
