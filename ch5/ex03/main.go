package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func getText(links []string, n *html.Node) []string {
	if n.Type == html.TextNode && (n.Parent).Data != "script" && (n.Parent).Data != "style" {
		links = append(links, n.Data)
	}
	if n.FirstChild != nil {
		links = getText(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = getText(links, n.NextSibling)
	}
	return links
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./ex03 http://example.com")
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return getText(nil, doc), nil
}
