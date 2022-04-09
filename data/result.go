package data

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Url string
}

func (t *Result) Get() string {
	resp, err := http.Get(t.Url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
