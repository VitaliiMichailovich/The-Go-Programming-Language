package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const MainURL = "https://xkcd.com/"

type comics struct {
	Month      string
	Num        int32
	Link       string
	Year       string
	News       string
	Safe_title string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func main() {
	var comicses = make([]comics, 0)
	if len(os.Args) == 1 {
		fmt.Println("First argument must be searching word or phrase.")
		os.Exit(0)
	}
	for i := 1; i <= 1917; i++ {
		res, err := xkcd(i)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			comicses = append(comicses, *res)
		}
	}
	prints := 0
	for i := 0; i <= len(comicses); i++ {
		if strings.Contains(comicses[i].Transcript, os.Args[1]) {
			fmt.Printf("----------------------------------------------------------------------------------- Comics: %d\n%s\n", comicses[i].Num, comicses[i].Transcript)
			prints++
		}
	}
	if prints == 0 {
		fmt.Println("There is NO search results.")
	}
}

func xkcd(id int) (*comics, error) {
	resp, err := http.Get(MainURL + "/" + strconv.Itoa(id) + "/info.0.json")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("ID: %d.\tSearch query failed: %s", id, resp.Status)
	}
	defer resp.Body.Close()
	var result comics
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return &result, nil
}
