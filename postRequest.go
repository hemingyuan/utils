package utils

import (
	"fmt"
	"io"
	"net/http"
)

// PostRequest 向指定url发起POST请求
func PostRequest(url string, headers map[string]string, body io.Reader) (res io.ReadCloser, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)

	if req, err = http.NewRequest("POST", url, body); err != nil {
		err = fmt.Errorf("Create New Request [%s] Error. Err: %v", url, err)
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if resp, err = http.DefaultClient.Do(req); err != nil {
		err = fmt.Errorf("Do Request [%s] Error. Err: %v", url, err)
		return
	}
	defer resp.Body.Close()

	res = resp.Body
	return
}
