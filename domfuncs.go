package sitescraper

import (
	"strings"
)

//Returns a filtered Dom containing all Tags that have the given Tag-Name(s)
func (d Dom) Tags(tagname ...string) Dom {
	if len(tagname) < 1 {
		return d
	}
	var tags []Tag
	for _, N := range tagname {
		for i, n := range d.Tag {
			if n.tagname == N {
				tags = append(tags, d.Tag[i])
			}
		}
	}
	var dm Dom
	dm.Tag = append(dm.Tag, tags...)
	return ParseHTML(dm.string())
}

//Returns a filtered Dom containing all Tags that contain the given Attribute(s)
func (d Dom) Attr(attr ...string) Dom {
	var tags []Tag

	for _, N := range attr {
		for _, n := range d.Tag {
			if strings.Contains(n.tagcontent, N+`="`) || strings.Contains(n.tagcontent, N+`=`) {
				tags = append(tags, n)
			}

		}
	}
	var dm Dom
	dm.Tag = append(dm.Tag, tags...)
	return ParseHTML(dm.string())

}

//Returns a filtered Dom containing all Tags that contain the given Attribute-Value(s)
func (d Dom) AttrValue(attrvalue ...string) Dom {
	var tags []Tag

	for _, N := range attrvalue {
		for _, n := range d.Tag {
			if strings.Contains(n.tagcontent, `="`+N+`"`) || strings.Contains(n.tagcontent, `=`+N+` `) || strings.Contains(n.tagcontent, `=`+N+`>`) {
				tags = append(tags, n)
			}

		}
	}
	var dm Dom
	dm.Tag = append(dm.Tag, tags...)
	return ParseHTML(dm.string())

}

//Returns whole innerHTML of all Tags of the Dom or filtered Dom as string
func (d Dom) GetInnerHTML() string {
	var s []string
	var cleared []string

	for _, tag := range d.Tag {
		s = append(s, tag.innerHTML)
	}

	for _, x := range s {
		in := false
		for _, c := range cleared {
			if x == c {
				in = true
			}
		}

		if !in {
			cleared = append(cleared, x)
		}
	}
	return strings.Join(cleared, " ")
}

//Returns the whole Text of all Tags of the Dom or filtered Dom as string
func (d Dom) GetText() string {

	var s []string
	var cleared []string

	for _, tag := range d.Tag {
		s = append(s, getText(tag.tagname, tag.innerHTML))
	}

	for _, tag := range d.Tag {
		s = append(s, tag.innerHTML)
	}

	for _, x := range s {
		in := false
		for _, c := range cleared {
			if x == c {
				in = true
			}
		}

		if !in {
			cleared = append(cleared, x)
		}
	}
	return strings.Join(cleared, " ")
}

//Returns the Attribute-Value of all Tags of the Dom filtered by the given Attribute-Name as string
func (d Dom) GetAttrValue(attrname string) string {

	var s []string

	for _, tag := range d.Tag {
		s = append(s, tag.GetAttrValue(attrname))
		s = append(s, " ")
	}

	return strings.Join(s, "")

}

func (d Dom) string() string {
	var str string
	for _, n := range d.Tag {
		str = str + n.tagcontent + n.innerHTML + "</" + n.tagname + ">"
	}
	return str
}
