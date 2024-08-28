package crawler

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	"codeflow.dananglin.me.uk/apollo/web-crawler/internal/util"
)

type Crawler struct {
	pages      map[string]pageStat
	baseURL    *url.URL
	mu         *sync.Mutex
	workerPool chan struct{}
	wg         *sync.WaitGroup
	maxPages   int
}

type pageStat struct {
	count    int
	internal bool
}

func NewCrawler(rawBaseURL string, maxWorkers, maxPages int) (*Crawler, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the base URL: %w", err)
	}

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	crawler := Crawler{
		pages:      make(map[string]pageStat),
		baseURL:    baseURL,
		mu:         &sync.Mutex{},
		workerPool: make(chan struct{}, maxWorkers),
		wg:         &waitGroup,
		maxPages:   maxPages,
	}

	return &crawler, nil
}

func (c *Crawler) Crawl(rawCurrentURL string) {
	// Add an empty struct to channel here
	c.workerPool <- struct{}{}

	// Decrement the wait group counter and free up the worker pool when
	// finished crawling.
	defer func() {
		<-c.workerPool
		c.wg.Done()
	}()

	if c.reachedMaxPages() {
		return
	}

	// get normalised version of rawCurrentURL
	normalisedCurrentURL, err := util.NormaliseURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("WARNING: Error normalising %q: %v.\n", rawCurrentURL, err)

		return
	}

	isInternalLink, err := c.isInternalLink(rawCurrentURL)
	if err != nil {
		fmt.Printf(
			"WARNING: Unable to determine if %q is an internal link; %v.\n",
			rawCurrentURL,
			err,
		)

		return
	}

	// Add (or update) a record of the URL in the pages map.
	// If there's already an entry of the URL in the map then return early.
	if existed := c.addPageVisit(normalisedCurrentURL, isInternalLink); existed {
		return
	}

	// if current URL is an external link then return early.
	if !isInternalLink {
		return
	}

	// Get the HTML from the current URL, print that you are getting the HTML doc from current URL.
	fmt.Printf("Crawling %q\n", rawCurrentURL)

	htmlDoc, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf(
			"WARNING: Error retrieving the HTML document from %q: %v.\n",
			rawCurrentURL,
			err,
		)

		return
	}

	// Get all the URLs from the HTML doc.
	links, err := util.GetURLsFromHTML(htmlDoc, c.baseURL.String())
	if err != nil {
		fmt.Printf(
			"WARNING: Error retrieving the links from the HTML document: %v.\n",
			err,
		)

		return
	}

	// Recursively crawl each URL on the page.
	for ind := range len(links) {
		c.wg.Add(1)
		go c.Crawl(links[ind])
	}
}

// isInternalLink evaluates whether the input URL is an internal link to the
// base URL. An internal link is determined by comparing the host names of both
// the input and base URLs.
func (c *Crawler) isInternalLink(rawURL string) (bool, error) {
	parsedRawURL, err := url.Parse(rawURL)
	if err != nil {
		return false, fmt.Errorf("error parsing the URL %q: %w", rawURL, err)
	}

	return c.baseURL.Hostname() == parsedRawURL.Hostname(), nil
}

// addPageVisit adds a record of the visited page's URL to the pages map.
// If there is already a record of the URL then it's record is updated (incremented)
// and the method returns true. If the URL is not already recorded then it is created
// and the method returns false.
func (c *Crawler) addPageVisit(normalisedURL string, internal bool) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.pages[normalisedURL]

	if exists {
		stat := c.pages[normalisedURL]
		stat.count++
		c.pages[normalisedURL] = stat
	} else {
		c.pages[normalisedURL] = pageStat{
			count:    1,
			internal: internal,
		}
	}

	return exists
}

func (c *Crawler) Wait() {
	c.wg.Wait()
}

func (c *Crawler) PrintReport() {
	c.mu.Lock()
	defer c.mu.Unlock()

	r := newReport(c.baseURL.String(), c.pages)

	fmt.Fprint(os.Stdout, r)
}

func (c *Crawler) reachedMaxPages() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return len(c.pages) >= c.maxPages
}
