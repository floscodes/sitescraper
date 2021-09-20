package sitescraper

type dom struct {
	tags []tag
}

type tag struct {
	tagname    string
	tagcontent string
	innerHTML  string
}
