package client

import "testing"

func TestHttpRequest(t *testing.T) {
	url := ""
	if _, err := MakeHTTPRequest[any](url, "GET", nil, nil, nil); err != nil {
		t.Errorf("%v", err.ResponseData)
	}
}
