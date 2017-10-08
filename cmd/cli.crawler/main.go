package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cameronnewman/cli.crawler/crawler"
)

func main() {

	url := flag.String("url", "", "URL to crawl")

	flag.Parse()

	if *url == "" {
		fmt.Println("No URL supplied")
		return
	}

	c, err := crawler.New(*url)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	c.Run()
}
