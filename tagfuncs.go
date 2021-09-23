package sitescraper

import "strings"

//Returns InnerHTML inside a Tag as string
func (t Tag) GetInnerHTML() string {
	return t.innerHTML
}

//Returns pure Text inside a Tag leaving out all nested HTML-Tags and their content
func (t Tag) GetText() string {
	return getText(t.tagname, t.innerHTML)
}

//Returns the name of the Tag as string
func (t Tag) GetTagName() string {
	return t.tagname
}

//Returns the Value of the given Attribute as string
func (t Tag) GetAttrValue(attr string) string {
	var out string

	if strings.Contains(t.tagcontent, attr+"=") {
		out = t.tagcontent[strings.Index(t.tagcontent, attr+`=`)+len(attr+`=`):]

	}

	if strings.Contains(out, ">") {
		out = strings.ReplaceAll(out, ">", "")
	}

	if strings.Contains(out, `"`) {
		out = strings.ReplaceAll(out, `"`, "")
	}

	if strings.Contains(out, " ") {
		out = out[:strings.Index(out, " ")]
	}

	return out
}
