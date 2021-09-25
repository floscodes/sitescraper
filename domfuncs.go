package sitescraper

import (
	"strings"
)

func (d Dom) Filter(filter ...string) Dom {

	d = ParseHTML(d.string())

	if len(filter) < 1 {
		return d
	}
	if filter[0] != "" && filter[0] != "*" {
		d = d.tags(filter[0])
	}

	if len(filter) < 2 {
		return d
	}
	if filter[1] != "" && filter[1] != "*" {
		d = d.attr(filter[1])
	}

	if len(filter) < 3 {
		return d
	}
	if filter[2] != "" && filter[2] != "*" {
		d = d.attrValue(filter[2])
	}

	return d
}

//Returns a filtered Dom containing all Tags that have the given Tag-Name(s)
func (d Dom) tags(tagname ...string) Dom {
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
	return dm
}

//Returns a filtered Dom containing all Tags that contain the given Attribute(s)
func (d Dom) attr(attr ...string) Dom {
	if len(attr) < 1 {
		return d
	}
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
	return dm

}

//Returns a filtered Dom containing all Tags that contain the given Attribute-Value(s)
func (d Dom) attrValue(attrvalue ...string) Dom {
	if len(attrvalue) < 1 {
		return d
	}
	var tags []Tag

	for _, N := range attrvalue {
		for _, n := range d.Tag {
			if strings.Contains(n.tagcontent, `="`+N+`"`) || strings.Contains(n.tagcontent, `=`+N+` `) || strings.Contains(n.tagcontent, `=`+N+`>`) || strings.Contains(n.tagcontent, `='`+N+`'>`) || strings.Contains(n.tagcontent, `='`+N+`'`) {
				tags = append(tags, n)
			}

		}
	}

	var dm Dom
	dm.Tag = append(dm.Tag, tags...)
	return dm

}

func (d Dom) string() string {
	var s string
	for _, n := range d.Tag {
		s = s + n.tagcontent + n.innerHTML + "</" + n.tagname + ">"
	}
	return s
}

//Returns whole innerHTML of all Tags of the Dom or filtered Dom as string
func (d Dom) GetInnerHTML() string {
	var s []string

	for _, tag := range d.Tag {
		s = append(s, tag.innerHTML)
	}

	var cleared []string

	for _, y := range s {
		in := false
		for _, x := range cleared {
			if y == x {
				in = true
			}
		}

		if !in {
			cleared = append(cleared, y)
		}
	}

	return strings.Join(cleared, " ")
}

//Returns the whole Text of all Tags of the Dom or filtered Dom as string
func (d Dom) GetText() string {

	var s []string

	for i, tag := range d.Tag {
		s = append(s, getText(tag.tagname, d.Tag[i].GetInnerHTML()))
	}

	var cleared []string

	for _, y := range s {
		in := false
		for _, x := range cleared {
			if y == x {
				in = true
			}
		}

		if !in {
			cleared = append(cleared, y)
		}
	}

	return strings.Join(cleared, " ")

}

//Returns the Attribute-Value of all Tags of the Dom filtered by the given Attribute-Name as string
func (d Dom) GetAttrValue(attrname string) string {

	var s []string

	for _, tag := range d.Tag {
		s = append(s, tag.GetAttrValue(attrname))
	}

	var cleared []string

	for _, y := range s {
		in := false
		for _, x := range cleared {
			if y == x {
				in = true
			}
		}

		if !in {
			cleared = append(cleared, y)
		}
	}

	return strings.Join(cleared, " ")

}
