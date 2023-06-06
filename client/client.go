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
)

type RequestError struct {
	StatusCode   int
	ResponseData map[string]interface{}
	Err          error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

// HTTP хүсэлт илгээгч
//
// default-р Content-Type нь application/json байгаа headers параметрээр өөрчилж болно
func MakeHTTPRequest[T any](fullUrl string, httpMethod string, headers *map[string]string, queryParameters *url.Values, body interface{}) (*T, *RequestError) {
	client := http.Client{}
	u, err := url.Parse(fullUrl)
	if err != nil {
		return nil, &RequestError{Err: err}
	}

	// if it's a GET, we need to append the query parameters.
	if httpMethod == "GET" {
		q := u.Query()
		if queryParameters != nil {
			for k, v := range *queryParameters {
				// this depends on the type of api, you may need to do it for each of v
				q.Set(k, strings.Join(v, ","))
			}
		}
		// set the query to the encoded parameters
		u.RawQuery = q.Encode()
	}

	// regardless of GET or POST, we can safely add the body
	jsonStrBytes, err := json.Marshal(body)
	if err != nil {
		return nil, &RequestError{Err: err}
	}
	req, err := http.NewRequest(httpMethod, u.String(), bytes.NewBuffer(jsonStrBytes))
	if err != nil {
		return nil, &RequestError{Err: err}
	}

	// for each header passed, add the header value to the request
	req.Header.Add("Content-Type", "application/json")
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}
	// optional: log the request for easier stack tracing
	log.Printf("%s %s\n", httpMethod, req.URL.String())

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
