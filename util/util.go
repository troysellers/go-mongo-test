package util

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func DoHttp(url string, body []byte, method string, headers map[string]string) (http.Header, []byte, error) {
	return doHttp(url, body, method, headers)
}

// returns the response body bytes if we had a 200 response.
// errors for all others.
func doHttp(url string, body []byte, method string, headers map[string]string) (http.Header, []byte, error) {

	defer trackTime(time.Now(), fmt.Sprintf("%v:%v", method, url))

	client := &http.Client{}
	var r *bytes.Reader
	if body != nil {
		r = bytes.NewReader(body)
	} else {
		r = bytes.NewReader(nil)
	}
	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, nil, err
	}
	for header, value := range headers {
		req.Header.Add(header, value)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return res.Header, nil, fmt.Errorf("unsuccesful attempt to call endpoint %v\n%v", res.Status, string(bytes))
	}

	log.Printf("Received %d bytes with http response %v\n", len(bytes), res.Status)
	return res.Header, bytes, nil
}

func trackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s\n", name, elapsed)
}
