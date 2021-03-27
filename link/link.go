package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

//Link from a html doc  (<a href="...."></a>)
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}
func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if "href" == attr.Key {
			link.Href = attr.Val
		}
	}
	link.Text = text(n)
	return link
}

func text(n *html.Node) string {

	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type == html.CommentNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = text(c)
	}
	return strings.Join(strings.Fields(ret), " ")

}
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
