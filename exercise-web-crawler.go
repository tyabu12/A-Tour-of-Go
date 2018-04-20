// +build ignore

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup
	msg := make(chan string)

	var mux sync.Mutex
	memo := map[string]bool{}

	var crawler func(string, int)

	crawler = func(url string, depth int) {
		defer wg.Done()
		if depth <= 0 {
			return
		}
		mux.Lock()
		if memo[url] {
			mux.Unlock()
			return
		}
		memo[url] = true
		mux.Unlock()
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			msg <- fmt.Sprintln(err)
			return
		}
		msg <- fmt.Sprintf("found: %s %q\n", url, body)
		wg.Add(len(urls))
		for _, u := range urls {
			go crawler(u, depth-1)
		}
	}

	wg.Add(1)
	go crawler(url, depth)

	go func() {
		wg.Wait()
		close(msg)
	}()

	for m := range msg {
		fmt.Print(m)
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
