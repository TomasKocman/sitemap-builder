package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercies.com", "the url to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
}
