package htmllinkparser

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func f(n *html.Node, a *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		newLink := Link{
			Href: n.Attr[0].Val,
			Text: n.FirstChild.Data,
		}
		*a = append(*a, newLink)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c, a)
	}
}

func ParseLinks(r io.Reader) []Link {
	doc, err := html.Parse(r)
	if err != nil {
		exit("Error parsing the file")
	}
	var arr []Link
	f(doc, &arr)
	return arr
}
