package sitescraper

import (
	"strings"
)

//Public-Function to Call the HTML-Parser
func ParseHTML(html string) Dom {
	return fetch(html)
}

//The private functions handle the parsing process
func fetch(html string) Dom {
	var dm Dom
	for {
		if !strings.Contains(html, "<") {
			break
		}

		tagname, tagcontent := getTagnameAndContent(html)

		html = html[strings.Index(html, tagcontent)+len(tagcontent):]

		//Check if string behing "<" is a valid tagname
		if checkTagname(tagname) {
			dm.tags = append(dm.tags, Tag{tagname, tagcontent, getInnerHTML(tagname, html)})
		}
	}

	dm = clearClosingTags(dm)
	return dm
}

func clearClosingTags(d Dom) Dom {
	var out Dom

	for _, n := range d.tags {
		if !strings.Contains(n.tagcontent, "</") {
			out.tags = append(out.tags, n)
		}
	}

	return out
}

func getTagnameAndContent(html string) (string, string) {
	tagcontent := html[strings.Index(html, "<"):]
	tagcontent = tagcontent[:strings.Index(tagcontent, ">")+1]
	var tagname string

	if strings.Contains(tagcontent, " ") {
		tagname = tagcontent[1:strings.Index(tagcontent, " ")]
	} else {
		tagname = tagcontent[1:strings.Index(tagcontent, ">")]
	}

	return tagname, tagcontent
}

func checkTagname(tagname string) bool {
	for _, n := range tagnames {
		if tagname == n {
			return true
		}
	}

	return false
}
