package sitescraper

import "strings"

func getInnerHTML(tagname, html string) string {

	firstpart := html[:strings.Index(html, "</"+tagname+">")+len("</"+tagname+">")]

	checkbreak := checkBreak(tagname, firstpart)
	if checkbreak {
		firstpart = strings.ReplaceAll(firstpart, "<br", "{||}")
	}

	appearance := strings.Count(firstpart, "<"+tagname) + 1

	if appearance < 1 {
		return firstpart
	}

	var secondparts []string

	x := 0
	for {
		if x == appearance {
			break
		}

		secondparts = append(secondparts, html[:strings.Index(html, "</"+tagname+">")+len("</"+tagname+">")])
		html = html[strings.Index(html, "</"+tagname+">")+len("</"+tagname+">"):]

		x = x + 1

	}

	out := strings.Join(secondparts, "")

	if checkbreak {
		out = strings.ReplaceAll(out, "{||}", "<br>")
	}

	if !strings.Contains(out, "</"+tagname+">") {
		return out
	}

	return out[:strings.LastIndex(out, "</"+tagname+">")]

}

func checkBreak(tagname, firstpart string) bool {
	if tagname == "b" {
		if strings.Contains(firstpart, "<br") {
			return true
		}
	}
	return false
}
