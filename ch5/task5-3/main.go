package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	task5_2("http://dreamer.pp.ua/")
}

func task5_2(url string)  {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	finder(doc)
	return
}

func finder(n *html.Node) {
	if n.Type == html.ElementNode && n.Data != "script" &&  n.Data != "style"{
		fmt.Printf("%v - %v\n", n.Data, n.Attr)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		finder(c)
	}
}
