package htmllinkparser

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var (
	htmlFile = flag.String("html", "html/ex1.html", "the route for the html file")
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

func main() {
	flag.Parse()
	file, err := os.Open(*htmlFile)
	if err != nil {
		exit("Error opening the file")
	}

	doc, err := html.Parse(file)
	if err != nil {
		exit("Error opening the file")
	}
	var arr []Link
	f(doc, &arr)
	fmt.Println(arr)
}
