package utils

import (
	"net/http"
	"net/http/httputil"
)

// DumpRequest --
func DumpRequest(req *http.Request) []byte {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil
	}
	return requestDump
}

// DumpResponse --
func DumpResponse(resp *http.Response) []byte {
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil
	}
	return responseDump
}
