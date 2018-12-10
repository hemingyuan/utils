package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// PostRequest 向指定url发起POST请求
func PostRequest(url string, headers map[string]string, body io.Reader) (respData []byte, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)

	if req, err = http.NewRequest("POST", url, body); err != nil {
		err = fmt.Errorf("ERROR: Create New Request [%s] Error. %v", url, err)
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if resp, err = http.DefaultClient.Do(req); err != nil {
		err = fmt.Errorf("ERROR: Do Request [%s] Error. %v", url, err)
		return
	}
	defer resp.Body.Close()

	if respData, err = ioutil.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("ERROR: Read Response Data from [%s] Error. %v", url, err)
		return
	}

	if resp.StatusCode != 200 && resp.StatusCode != 301 && resp.StatusCode != 302 {
		err = fmt.Errorf("ERROR: %s\n%s", resp.Status, respData)
		return
	}
	return
}
