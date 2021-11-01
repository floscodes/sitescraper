package sitescraper

import (
	"strings"
)

func getInnerHTML(tagname, html string) string {

	if !strings.Contains(html, "</"+tagname) {
		return html
	}

	closingtag := html[strings.Index(html, "</"+tagname):]
	closingtag = closingtag[:strings.Index(closingtag, ">")+1]

	firstpart := html[:strings.Index(html, closingtag)+len(closingtag)]

	if checkBreak(tagname, firstpart) {
		firstpart = strings.ReplaceAll(firstpart, "<br>", breakSubst)
	}

	appearance := strings.Count(firstpart, "<"+tagname)

	if appearance < 1 {
		if strings.LastIndex(firstpart, closingtag) != -1 {
			firstpart = firstpart[:strings.LastIndex(firstpart, closingtag)]
		}
		return firstpart
	}

	//Cut away firstpart
	html = html[strings.Index(html, closingtag)+len(closingtag):]

	var secondparts []string

	x := 0
	for {
		if x > appearance {
			break
		}

		if len(html) < 1 {
			break
		}

		if !strings.Contains(html, "</"+tagname) {
			break
		}

		closingtag = html[strings.Index(html, "</"+tagname):]
		closingtag = closingtag[:strings.Index(closingtag, ">")+1]
		secondparts = append(secondparts, html[:strings.Index(html, closingtag)+len(closingtag)])
		html = html[strings.Index(html, closingtag)+len(closingtag):]

		x = x + 1

	}
	//firstpart added
	out := firstpart + strings.Join(secondparts, "")

	if strings.LastIndex(out, closingtag) != -1 {
		closingtag = out[strings.Index(out, "</"+tagname):]
		closingtag = closingtag[:strings.Index(closingtag, ">")+1]
		out = out[:strings.LastIndex(out, closingtag)]
	}

	return out

}

func checkBreak(tagname, firstpart string) bool {
	if tagname == "b" {
		if strings.Contains(firstpart, "<br") {
			return true
		}
	}
	return false
}

const breakSubst string = "{|!!45sdhj26768892347|}"
