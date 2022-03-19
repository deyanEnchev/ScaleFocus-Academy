package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

func main() {
	var concTasks int
	flag.IntVar(&concTasks, "c", 2, "Type urls here")
	flag.Parse()


	var urls = make([]string, 0, 8)
	urls = append(urls, flag.Args()...)

	resultsChan := fetchURLs(urls, concTasks)

	for url := range resultsChan {
		_ = url	
	}
}

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}

func fetchURLs(urls []string, concurrency int) chan string {
	processQueue := make(chan string, concurrency)

	outChan := make(chan string)
	var wg sync.WaitGroup

	go func() {
		for _, urlToProcess := range urls {
			wg.Add(1)
			processQueue <- urlToProcess

			go func(url string) {
				defer wg.Done()
				pingURL(url)
				<-processQueue
				outChan <- url
			}(urlToProcess)
		}
		wg.Wait()
		close(outChan)
	}()

	return outChan
}

// use these commands to run on Windows:
// go build -o multiping.exe main.go
// ./multiping.exe -c 5 https://google.com https://facebook.com
// with more urls: (this one is slow, try more concurrent tasks to make it faster)
// ./multiping.exe -c 2 https://google.com https://facebook.com https://twitter.com https://instagram.com https://youtube.com https://facebook.com https://twitter.com https://instagram.com https://youtube.com https://facebook.com https://twitter.com https://instagram.com https://youtube.com https://facebook.com https://twitter.com https://instagram.com https://youtube.com