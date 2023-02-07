package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	// f := parser.ParseHTML(".././src/ex1.html")
	file, err := ioutil.ReadFile("../.././src/ex1.html")
	if err != nil {
		log.Println(err)
		return
	}

	doc, err := html.Parse(strings.NewReader(string(file)))
	if err != nil {
		log.Fatal(err)
	}

	l := Link{}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					l.Href = a.Val
					
					fmt.Printf("l: %v\n", l)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}
