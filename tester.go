package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func test() {

	url := "http://localhost:8080/videos"
	method := "GET"

	payload := strings.NewReader(`{
    "title": "title 3",
    "description": "desc 3",
    "url": "url 3"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
