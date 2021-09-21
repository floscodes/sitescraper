package sitescraper

import (
	"strings"
)

func (d *Dom) GetTags(tagname ...string) []Tag {
	if len(tagname) < 1 {
		return d.tags
	}
	var tags []Tag
	for _, N := range tagname {
		for i, n := range d.tags {
			if n.tagname == N {
				tags = append(tags, d.tags[i])
			}
		}
	}

	return tags
}

func (d *Dom) GetByAttr(attr, value string) []Tag {
	var tags []Tag

	for _, n := range d.tags {
		if strings.Contains(n.tagcontent, attr+`="`+value+`"`) || strings.Contains(n.tagcontent, attr+`=`+value+` `) {
			tags = append(tags, n)
		}

	}
	return tags
}
