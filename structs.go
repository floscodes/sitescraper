package sitescraper

type Dom struct {
	tags []Tag
}

type Tag struct {
	tagname    string
	tagcontent string
	innerHTML  string
}
