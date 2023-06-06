package client

import "testing"

func TestHttpRequest(t *testing.T) {

	config := &RequestConfig{
		Url:    "",
		Method: "",
	}

	if _, err := MakeHTTPRequest[any](config); err != nil {
		t.Errorf("%v", err.ResponseData)
	}
}
