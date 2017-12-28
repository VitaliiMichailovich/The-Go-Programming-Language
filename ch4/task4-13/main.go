package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const MainURL = "http://www.omdbapi.com/?&apikey=af103384&s="

type Movie struct {
	Title  string
	Year   string
	IMDBID string
	Type   string
	Poster string
}

type SearchResult struct {
	Search       []Movie
	totalResults string
	Response     string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("There is NO search title.")
		os.Exit(0)
	}
	searchTitle := ""
	startSpace := ""
	for _, v := range os.Args[1:] {
		searchTitle += startSpace + v
		startSpace = "+"
	}
	search, err := omdb(searchTitle)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < len(search); i++ {
		if search[i].Poster != "N/A" {
			poster, err := http.Get(search[i].Poster)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer poster.Body.Close()
			file, err := os.Create("./ch4/task4-13/" + search[i].IMDBID+".jpg")
			if err != nil {
				fmt.Println(err.Error())
			}
			// Use io.Copy to just dump the response body to the file. This supports huge files
			_, err = io.Copy(file, poster.Body)
			if err != nil {
				fmt.Println(err.Error())
			}
			file.Close()
			fmt.Printf("%v\n%s\n", search[i].Title, search[i].Poster)
		}
	}
}

func omdb(title string) ([]Movie, error) {
	resp, err := http.Get(MainURL + title)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return result.Search, nil
}
