package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

func ParseHTML(filename string) func(*html.Node) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	doc, err := html.Parse(strings.NewReader(string(file)))
	if err != nil {
		return nil
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)

					break
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(doc)
	return f
}
