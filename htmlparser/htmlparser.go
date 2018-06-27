package htmlparser

import (
	"golang.org/x/net/html"
)

// GetElementByID returns the HTML Node for a given ID
func GetElementByID(n *html.Node, ID string) *html.Node {
	return traverse(n, ID)
}

// getAttribute returns the HTML Attribute for a given HTML Node
func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

// checkID ensures that a given ID exists in a Node
func checkID(n *html.Node, ID string) bool {
	if n.Type == html.ElementNode {
		s, ok := getAttribute(n, "id")
		if ok && s == ID {
			return true
		}
	}
	return false
}

// traverse iterates over all nodes in an HTML document
func traverse(n *html.Node, ID string) *html.Node {
	if checkID(n, ID) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverse(c, ID)
		if result != nil {
			return result
		}
	}

	return nil
}
