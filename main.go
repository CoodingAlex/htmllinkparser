package main

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
		var attr string
		for _, att := range n.Attr {
			if att.Key == "href" {
				attr = att.Val
			}
		}
		fmt.Println(n.Attr[0].Key)
		newLink := Link{
			Href: attr,
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

func main() {
	file, err := os.Open("html/ex2.html")
	if err != nil {
		exit("err opening the file")
	}
	links := ParseLinks(file)
	fmt.Println(links)
}
