package htmlparser

import (
	"golang.org/x/net/html"
)

// GetElementByID returns the HTML Node for a given ID
func GetElementByID(n *html.Node, ID string) *html.Node {
	return traverse(n, ID)
}

func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func checkID(n *html.Node, ID string) bool {
	if n.Type == html.ElementNode {
		s, ok := getAttribute(n, "id")
		if ok && s == ID {
			return true
		}
	}
	return false
}

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
