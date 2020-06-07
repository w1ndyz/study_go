package main

import "golang.org/x/net/html"

func ElementByID(doc *html.Node, id string) *html.Node  {
	var elem *html.Node
	forEachNode(doc, func(n *html.Node) bool  {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.val == id {
					elem = n
					return false
				}
			}
		}
		return true
	}, nil)
	return elem
}