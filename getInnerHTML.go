package sitescraper

import "strings"

func getInnerHTML(tagname, html string) string {

	if !strings.Contains(html, "</"+tagname) {
		return ""
	}

	innerHTML := html[:strings.Index(html, "</"+tagname)]
	lastpart := html[:strings.LastIndex(html, "</"+tagname)]
	appearance := strings.Count(innerHTML, "<"+tagname)

	if appearance < 1 {
		return innerHTML
	} else {
		var parts []string

		for x := 0; x < appearance; x++ {
			parts = append(parts, html[:strings.Index(html, "</"+tagname)])
			html = html[:strings.LastIndex(html, "</"+tagname)]
		}

		return strings.Join(parts, "") + lastpart
	}

}
