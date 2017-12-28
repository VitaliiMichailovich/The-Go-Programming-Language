package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"sort"
)

type Pair struct {
	Key string
	Val int
}

type PairList []Pair

func main() {
	task, _ := task5_2("http://dreamer.pp.ua/")
	pl := make(PairList, len(task))
	i := 0
	for k, v := range task {
		pl[i] = Pair{k,v}
		i++
	}
	sort.Sort(pl)
	fmt.Println(pl)
}

func task5_2(url string) (map[string]int, error) {
	var out map[string]int
	resp, err := http.Get(url)
	if err != nil {
		return out, fmt.Errorf("%v", err.Error())
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	out = finder(make(map[string]int), doc)
	return out, nil
}

func finder(stack map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		stack[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		finder(stack, c)
	}
	return stack
}

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Val < p[j].Val }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
