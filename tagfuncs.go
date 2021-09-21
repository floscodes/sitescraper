package sitescraper

import "strings"

func (t *Tag) InnerHTML() string {
	return t.innerHTML
}

func (t *Tag) Text() string {
	return getText(t.innerHTML)
}

func (t *Tag) GetTagName() string {
	return t.tagname
}

func (t *Tag) GetAttr(attr string) string {
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
