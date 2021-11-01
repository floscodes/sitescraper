# sitescraper
Scraping Websites in Go!

# Examples:

### Get InnerHTML:

```
html := "<html><body><div id="hello">Hello <u>World!</u></div></body></html>"

dom := sitescraper.ParseHTML(html)

innerHTML := dom.Filter("body").GetInnerHTML()

fmt.Println(innerHTML)

//Output: <div id="hello">Hello <u>World!</u></div>

```

### Get Text:
```
html := "<html><body><div id="hello">Hello World!</div></body></html>"

dom := sitescraper.ParseHTML(html)

text := dom.Filter("div", "id", "hello").GetText()

fmt.Println(text)

//Output: Hello World!

```

### Get Text from single Tags:

```
html := "<html><body><div>Hello World!</div><div>My name is Sam!</div></body></html>"

dom := sitescraper.ParseHTML(html)

dom = dom.Filter("div")

fmt.Println(dom.Tag[0].GetText())  //Output: Hello World!

fmt.Println(dom.Tag[1].GetText())  //Output: My name is Sam!

```

**Works also with GetInnerHTML()**



### Get Website-Content:

```
html, err := sitescraper.Get("http://example.com/")

if err != nil {
    log.Fatal(err)
}

dom := sitescraper.ParseHTML(html)

dom = dom.Filter("div")

fmt.Println(dom.GetInnerHTML())

```

