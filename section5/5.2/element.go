package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find element: %v\n", err)
		os.Exit(1)
	}
	elements := map[string]int{}
	visit(elements, doc)
	for elem, count := range elements {
		fmt.Printf("%s\t%d\n", elem, count)
	}
}

func visit(e map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		e[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(e, c)
	}
}
