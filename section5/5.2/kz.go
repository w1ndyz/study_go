package main

import "golang.org/x/net/html"

var filter = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
}

func visit(links []string, n *html.Node) []string {
	for k, v := range filter {
		if n.Type == html.ElementNode && n.Data == k {
			for _, a := range n.Attr {
				if a.Key == v {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
