package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type RequestSchema struct {
    Method  string `json:"method"`
    Url     string `json:"url"`
    Body    string `json:"body"`
    Headers string `json:"headers"`
}

func sendRequestFromRequestSchema(rs RequestSchema) {
    body := strings.NewReader(rs.Body)
    request, err := http.NewRequest(rs.Method, rs.Url, body)
    if err != nil {
        panic(err)
    }
    headers := strings.Split(rs.Headers, "\n")
    for i := range headers {
        header := strings.SplitN(headers[i], ":", 2)
        request.Header.Add(header[0], header[1])
    }
    response, err := http.DefaultClient.Do(request)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    content, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }
    log.Print(string(content))
}
