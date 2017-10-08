package crawler

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asciimoo/colly"

	"github.com/goware/urlx"
	"github.com/joeguo/tldextract"
)

const (
	domainSuffixListFile string = "public_suffix_list.dat"
)

// Crawler struct
type Crawler struct {
	SeedRawURL string
	SeedURL    *url.URL
	tldExtract *tldextract.TLDExtract
	collector  *colly.Collector
}

//New Crawler instance
func New(u string) (*Crawler, error) {

	rawURL, err := url.Parse(u)
	if err != nil {
		return &Crawler{}, err
	}

	c := &Crawler{
		SeedRawURL: u,
		SeedURL:    rawURL,
	}

	c.collector = colly.NewCollector()
	c.collector.MaxDepth = 2

	tldExtract, err := tldextract.New(domainSuffixListFile, false)
	if err != nil {
		return c, err
	}

	c.tldExtract = tldExtract

	return c, nil
}

//Run starts crawlering the URL
func (c *Crawler) Run() {

	result := Links{
		Links: []LinkMetadata{},
		Page:  c.SeedRawURL,
	}

	c.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "#") {
			return
		}

		normalized, err := c.normalizeURL(link)
		if err != nil {
			return
		}

		linkDetails, _ := c.getLinkDetails(normalized)
		result.Links = append(result.Links, linkDetails)
		//fmt.Println(normalized)
		//c.collector.Visit(e.Request.AbsoluteURL(link))
	})
	c.collector.Visit(c.SeedRawURL)

	result.Count = len(result.Links)

	b, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func (c *Crawler) getLinkDetails(s string) (LinkMetadata, error) {
	result := LinkMetadata{}

	u, err := urlx.Parse(s)
	if err != nil {
		return result, err
	}

	normalizedURL, err := c.normalizeURL(s)
	if err != nil {
		return result, err
	}

	domain := c.tldExtract.Extract(u.Host)

	if domain.Flag == tldextract.Domain {

		destinationURL, err := c.followURLRedirects(normalizedURL)
		if err != nil {
			return result, err
		}

		originalURLHash := c.stringHash(normalizedURL)
		destinationURLHash := c.stringHash(destinationURL)

		var redirects bool

		if originalURLHash != destinationURLHash {
			redirects = true
		}

		result.Domain = domain.Root + "." + domain.Tld
		result.TLD = domain.Tld
		result.OriginalURL = normalizedURL
		result.OriginalURLHash = originalURLHash
		result.DestinationURL = destinationURL
		result.DestinationURLHash = destinationURLHash
		result.Redirects = redirects
	}

	return result, nil
}

func (c *Crawler) followURLRedirects(link string) (string, error) {

	normalized, err := c.normalizeURL(link)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(normalized)
	if err != nil {
		return "", err
	}

	return resp.Request.URL.String(), nil
}

func (c *Crawler) normalizeURL(link string) (string, error) {

	u, err := urlx.Parse(link)
	if err != nil {
		return "", err
	}

	if c.SeedURL.Host == u.Host {
		u.Scheme = c.SeedURL.Scheme
	}

	normalized, err := urlx.Normalize(u)
	if err != nil {
		return "", err
	}

	return normalized, nil
}

func (c *Crawler) stringHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
