package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	var url = "http://lroolle.com"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozila")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}

	resp2, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()
	s2, err := httputil.DumpResponse(resp2, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s2)
}
