package sitescraper

import (
	"strings"
)

func (d *dom) GetTags(tagname ...string) []tag {
	if len(tagname) < 1 {
		return d.tags
	}
	var tags []tag
	for _, N := range tagname {
		for i, n := range d.tags {
			if n.tagname == N {
				tags = append(tags, d.tags[i])
			}
		}
	}

	return tags
}

func (d *dom) GetByAttr(attr, value string) []tag {
	var tags []tag

	for _, n := range d.tags {
		if strings.Contains(n.tagcontent, attr+`="`+value+`"`) || strings.Contains(n.tagcontent, attr+`=`+value+` `) {
			tags = append(tags, n)
		}

	}
	return tags
}
