package main

import (
	"os"
	"log"
	"time"
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"sort"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func (p IssuesSearchResult) Len() int {
	return len(p.Items)
}

func (p IssuesSearchResult) Less(i, j int) bool {
	return p.Items[i].CreatedAt.After(p.Items[j].CreatedAt)
}

func (p IssuesSearchResult) Swap(i, j int) {
	p.Items[i], p.Items[j] = p.Items[j], p.Items[i]
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d items:\n", result.TotalCount)
	sort.Sort(result)
	var m, y, my bool
	for _, item := range result.Items {
		if curType := (time.Since(item.CreatedAt)).Hours()/24/30; curType < 1 {
			if m != true {
				fmt.Println("=========================== Less than 1 month ==========================")
				m = true
			}
		} else if curType >= 1 && curType < 12 {
			if y != true {
				fmt.Println("=========================== Less than 1 year  ==========================")
				y = true
			}
		} else {
			if my != true {
				fmt.Println("=========================== More than year    ==========================")
				my = true
			}
		}
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}