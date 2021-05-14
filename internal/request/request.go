package request

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Request struct {
    Url     string `json:"url"`
    Headers string `json:"headers"`
    Body    string `json:"body"`
    Method  string `json:"method"`
}

func (r *Request) Send() (string, string, error) {
    log.Println("sendRequest: " + r.Url)
    body := strings.NewReader(r.Body)
    request, err := http.NewRequest(r.Method, r.Url, body)
    if err != nil {
        return "", "", err
    }
    headers := strings.Split(r.Headers, "\n")
    for _, header := range headers {
        if len(header) == 0 {
            continue
        }
        header := strings.SplitN(header, ":", 2)
        if len(header) < 2 || len(header[1]) == 0 {
            continue
        }
        request.Header.Add(header[0], header[1])
    }
    response, err := http.DefaultClient.Do(request)
    if err != nil {
        return "", "", err
    }
    defer response.Body.Close()
    content, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return "", "", err
    }
    h := ""
    for key, elements := range response.Header {
        for _, element := range elements {
            h += key + ":" + element + "\n"
        }
    }
    return string(content), h, nil
}
