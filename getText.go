package sitescraper

import (
	"strings"
)

func getText(tagname, innerhtml string) string {

	if !strings.Contains(innerhtml, ">") {
		return innerhtml
	}

	if !strings.Contains(innerhtml, "<") {
		return innerhtml
	}

	checkbreak := checkBreak(tagname, innerhtml)
	if checkbreak {
		innerhtml = strings.ReplaceAll(innerhtml, "<br>", "|{}|")
	}

	firstpart := innerhtml[:strings.Index(innerhtml, "<")]
	middle := innerhtml[strings.Index(innerhtml, "<") : strings.LastIndex(innerhtml, ">")+1]
	lastpart := innerhtml[strings.LastIndex(innerhtml, ">")+1:]

	appearance := strings.Count(middle, "<")
	var parts1 []string
	for x := 0; x < appearance; x++ {
		tagname, _ := getTagnameAndContent(middle)
		if checkTagname(tagname) {
			break
		} else {
			middle = middle[1:]
			if !strings.Contains(middle, "<") {
				break
			}
			parts1 = append(parts1, "<"+middle[:strings.Index(middle, "<")])
			middle = middle[strings.Index(middle, "<"):]
		}
	}

	var parts2 []string
	for {

		if !strings.Contains(middle, "<") {
			break
		}

		part := middle[strings.LastIndex(middle, "<"):]

		tagname, _ := getTagnameAndContent(part)
		if strings.Contains(tagname, "/") {
			tagname = strings.ReplaceAll(tagname, "/", "")
		}
		if checkTagname(tagname) {
			break
		}

		parts2 = append(parts2, part)
		middle = middle[:strings.LastIndex(middle, part)]

	}

	var parts2sorted []string
	x := len(parts2) - 1
	for {

		if x < 0 {
			break
		}

		parts2sorted = append(parts2sorted, parts2[x])

		x = x - 1

	}

	out := firstpart + " " + strings.Join(parts1, "") + " " + strings.Join(parts2sorted, "") + " " + lastpart

	if checkbreak {
		out = strings.ReplaceAll(out, "|{}|", "<br>")
		out = strings.TrimRight(out, "<br>")
		out = strings.TrimLeft(out, "<br>")
	}
	return strings.Trim(out, " ")
}
