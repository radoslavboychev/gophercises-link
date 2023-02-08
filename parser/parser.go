package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/radoslavboychev/gophercises-link/models"
	"github.com/radoslavboychev/gophercises-link/parsingerrors"
	"golang.org/x/net/html"
)

func ParseHTML(filename string) ([]models.Link, error) {

	// read the html file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return []models.Link{}, parsingerrors.ErrOpeningFile
	}

	// create new reader
	r := strings.NewReader(string(file))

	// create a new tokenizer
	tokenizer := html.NewTokenizer(r)

	// create a struct to hold all links
	res := []models.Link{}

	// loop through page tokens
	for {
		// scans the next html token and returns its type
		tokenType := tokenizer.Next()

		// if the token is an error token, return and break the loop
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				if len(res) != 0 {
					return res, nil
				} else {
					return nil, parsingerrors.ErrEOF
				}
			}
			break
		}

		// if the token type is a start tag token
		// return the current token
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			// create a new link object to map the token to
			l := models.Link{}
			// if the tokan is an anchor element
			if token.Data == "a" {
				// check the next token type
				tokenType = tokenizer.Next()
				// if its an html text token
				if tokenType == html.TextToken {
					for _, a := range token.Attr {
						l.Text = strings.TrimSpace(tokenizer.Token().Data)
						l.Href = strings.TrimSpace(a.Val)
						res = append(res, l)
						fmt.Printf("Href: %v \n Text: %v\n", l.Href, l.Text)
					}
				}
			}
		}
	}

	return res, nil
}
