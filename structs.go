package sitescraper

type Dom struct {
	Tag      []Tag
	isparsed bool
}

type Tag struct {
	tagname    string
	tagcontent string
	innerHTML  string
}
