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
func Crawl(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	locker := sync.Mutex{}
	var crawl func(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool)
	crawl = func(url string, depth int, fetcher Fetcher, ch chan string, history map[string]bool) {
		if depth <= 0 {
			return
		}

		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			fmt.Println(err)
			return
		}

		locker.Lock()
		history[url] = true
		ch <- fmt.Sprintf("Found %s %q %v", url, body, urls)
		locker.Unlock()

		for _, u := range urls {
			locker.Lock()
			if !history[u] {
				history[u] = true
				locker.Unlock()

				crawl(u, depth-1, fetcher, ch, history)
			} else {
				locker.Unlock()
				fmt.Printf("    %s is crawled\n", u)
			}
		}
	}

	crawl(url, depth, fetcher, ch, history)

	close(ch)
	return
}

func main() {
	ch := make(chan string)
	history := make(map[string]bool)
	go Crawl("http://golang.org/", 4, fetcher, ch, history)

	for url := range ch {
		fmt.Println(url)
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
