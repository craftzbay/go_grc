package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestError struct {
	StatusCode   int
	ResponseData map[string]interface{}
	Err          error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

// HTTP хүсэлтийн тохиргоо
//
// Url: хүсэлтийн url заавал байх ёстой
//
// Method: хүсэлтийн төрөл явуулаагүй үед GET байна
//
// Headers: default-р Content-Type нь application/json байна
//
// Parameters: query parameters хүсэлтийн төрөл GET үед ашиглана
//
// Body: хүсэлтийн бие
//
// Timeout: хүсэлтийн timeout явуулаагүй үед 30сек байна
type RequestConfig struct {
	Url        string
	Method     string
	Headers    *map[string]string
	Parameters *url.Values
	Body       interface{}
	Timeout    uint
}

// HTTP хүсэлт илгээгч
//
// *RequestConfig: хүсэлтийн тохиргоо
func MakeHTTPRequest[T any](config *RequestConfig) (*T, *RequestError) {

	if config.Method == "" {
		config.Method = "GET"
	}

	if config.Timeout == 0 {
		config.Timeout = 30
	}

	client := http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
	}
	u, err := url.Parse(config.Url)
	if err != nil {
		return nil, &RequestError{Err: err}
	}

	config.Method = strings.ToUpper(config.Method)

	// if it's a GET, we need to append the query parameters.
	if config.Method == "GET" {
		q := u.Query()
		if config.Parameters != nil {
			for k, v := range *config.Parameters {
				// this depends on the type of api, you may need to do it for each of v
				q.Set(k, strings.Join(v, ","))
			}
		}
		// set the query to the encoded parameters
		u.RawQuery = q.Encode()
	}

	// regardless of GET or POST, we can safely add the body
	jsonStrBytes, err := json.Marshal(config.Body)
	if err != nil {
		return nil, &RequestError{Err: err}
	}
	req, err := http.NewRequest(config.Method, u.String(), bytes.NewBuffer(jsonStrBytes))
	if err != nil {
		return nil, &RequestError{Err: err}
	}

	// for each header passed, add the header value to the request
	req.Header.Add("Content-Type", "application/json")
	if config.Headers != nil {
		for k, v := range *config.Headers {
			req.Header.Set(k, v)
		}
	}
	// optional: log the request for easier stack tracing
	log.Printf("%s %s\n", config.Method, req.URL.String())

	// finally, do the request
	res, err := client.Do(req)
	if err != nil {
		return nil, &RequestError{Err: err}
	}

	if res == nil {
		return nil, &RequestError{Err: fmt.Errorf("error: calling %s returned empty response", u.String())}
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &RequestError{Err: err}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		httpErr := &RequestError{StatusCode: res.StatusCode, Err: errors.New("http request failed")}
		json.Unmarshal(responseData, &httpErr.ResponseData)
		return nil, httpErr
	}

	responseObject := new(T)
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Printf("error unmarshaling response: %+v", err)
		return nil, &RequestError{Err: err}
	}

	return responseObject, nil
}
