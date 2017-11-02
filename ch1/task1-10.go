package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	var file *os.File
	if _, err := os.Stat("task1-10.txt"); os.IsNotExist(err) {
		file, err = os.Create("task1-10.txt")
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		file, err = os.OpenFile("task1-10.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	defer file.Close()
	for range os.Args[1:] {
		if _, err := file.WriteString("\r\n"+<-ch); err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}