package main

import (
	"fmt"

	"scraping/data"
)

func main() {
	res := new(data.Result)
	res.Url = "https://twitter.com/"
	fmt.Println(res.Get())
}
