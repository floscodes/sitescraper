package sitescraper

type Dom struct {
	Tag []Tag
}

type Tag struct {
	tagname    string
	tagcontent string
	innerHTML  string
}
