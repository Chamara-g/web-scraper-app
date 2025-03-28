package services

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/gihanc.dev/web-scraper-app/internal/models"
	"github.com/gihanc.dev/web-scraper-app/internal/utils"
	"github.com/gocolly/colly"
)

func GetSiteDataByURL(url string) (*models.SiteData, error) {

	siteData := &models.SiteData{}
	siteData.URL = url

	if url == "" {
		return nil, errors.New("URL not found")
	}

	// Scrape data
	c := colly.NewCollector()

	// get HTML version
	c.OnResponse(func(r *colly.Response) {

		htmlContent := string(r.Body)
		siteData.HTML_VERSION = DetectHTMLVersion(htmlContent)
	})

	// Extract Page Title data
	c.OnHTML("head title", func(e *colly.HTMLElement) {
		siteData.TITLE = e.Text
	})

	// Heading count
	siteData.HEADING_LEVELS = CountHeadings(c)

	// Check login form
	siteData.HAVE_LOGIN_FORM = ContainsLoginForm(c)

	// check links
	internal, external, inaccessible := CountPageLinks(url)

	links := map[string][]string{
		"internal": {},
		"external": {},
		"inaccessible": {},
	}

	links["internal"] = internal
	links["external"] = external
	links["inaccessible"] = inaccessible

	siteData.LINKS = links

	// Handle errors
	c.OnError(func(_ *colly.Response, err error) {
		// TODO
		println(err.Error())
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return siteData, nil
}

// find HTML version
func DetectHTMLVersion(htmlContent string) string {

	var htmlVersion string
	
	content := strings.ToLower(htmlContent)
	
	doctypeRegex := regexp.MustCompile(`(?i)<!DOCTYPE\s+([^>]+)>`)
	matchesDocTypes := doctypeRegex.FindStringSubmatch(content)
	
	if len(matchesDocTypes) > 1 {
		doctype := strings.ToLower(matchesDocTypes[1]) 
	
		switch {
			case strings.Contains(doctype, "html") && !strings.Contains(doctype, "public"):
				htmlVersion = "HTML5"
	
			case strings.Contains(doctype, "xhtml"):
				htmlVersion = "XHTML"
	
			case strings.Contains(doctype, "html 4.01"):
				htmlVersion = "HTML 4.01"
	
			case strings.Contains(doctype, "html 3.2"):
				htmlVersion = "HTML 3.2"
	
			case strings.Contains(doctype, "html 2.0"):
				htmlVersion = "HTML 2.0"
	
			default:
				htmlVersion = "Unknown HTML Version"
		}
	} else {
		htmlVersion = "Unknown HTML Version"
	}

	return htmlVersion

}

// calculate heading levels
func CountHeadings(c *colly.Collector) map[string]int {

	headingCount := map[string]int{
		"h1": 0,
		"h2": 0,
		"h3": 0,
		"h4": 0,
		"h5": 0,
		"h6": 0,
	}

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		headingCount["h1"]++
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		headingCount["h2"]++
	})

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		headingCount["h3"]++
	})

	c.OnHTML("h4", func(e *colly.HTMLElement) {
		headingCount["h4"]++
	})

	c.OnHTML("h5", func(e *colly.HTMLElement) {
		headingCount["h5"]++
	})

	c.OnHTML("h6", func(e *colly.HTMLElement) {
		headingCount["h6"]++
	})

	return headingCount
}

// check if login form found
func ContainsLoginForm(c *colly.Collector) bool {

	loginFormFound := false

	c.OnHTML("form", func(e *colly.HTMLElement) {

		// check input field names
		e.ForEach("input", func(i int, el *colly.HTMLElement) {

			inputName := el.Attr("name")
			
			if strings.Contains(inputName, "user") || strings.Contains(inputName, "name") || 
				strings.Contains(inputName, "pass") || strings.Contains(inputName, "login") ||
				strings.Contains(inputName, "email") {
				loginFormFound = true
			}
		})

		// check button text
		e.ForEach("button", func(i int, el *colly.HTMLElement) {

			buttonText := el.Text
			
			if strings.Contains(strings.ToLower(buttonText), "login") || strings.Contains(strings.ToLower(buttonText), "signin") {
				loginFormFound = true
			}
		})

		// check if used method is post
		method := e.Attr("method")
		if method == "post" && loginFormFound {
			loginFormFound = true
		}

	})

	return loginFormFound
}

// check external and internal links
func CountPageLinks(siteURL string) ([]string, []string, []string) {

	var internalLinks = []string{}
	var externalLinks = []string{}
	var inaccessibleLinks = []string{}

	baseURL, err := url.Parse(siteURL)
	if err != nil {
		return internalLinks, externalLinks, inaccessibleLinks
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	c := colly.NewCollector()

	// check a tags
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link == "" || !utils.IsValidURL(link){
			return
		}

		parsedLink, err := url.Parse(link)
		if err != nil {
			return
		}

		if parsedLink.Host == baseURL.Host {

			mu.Lock()
			internalLinks = utils.AppendIfNotPresent(internalLinks, link)
			mu.Unlock()
		} else {

			mu.Lock()
			externalLinks =  utils.AppendIfNotPresent(externalLinks, link)
			mu.Unlock()
		}

		wg.Add(1)

		// check accessibility of link
		go func(link string) {

			defer wg.Done()

			resp, err := http.Head(link)
			if err != nil || resp.StatusCode >= 400 {
				
				mu.Lock()
				inaccessibleLinks =  utils.AppendIfNotPresent(inaccessibleLinks, link)
				mu.Unlock()
			}
		}(link)

	})

	err = c.Visit(siteURL)
	if err != nil {
		return internalLinks, externalLinks, inaccessibleLinks
	}

	wg.Wait()

	return internalLinks, externalLinks, inaccessibleLinks
}