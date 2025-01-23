package access

import (
	"github.com/Yui100901/MyGo/network/http_utils"
	"time"
)

//
// @Author yfy2001
// @Date 2025/1/14 10 24
//

type HTTPDemoClient struct {
	ID         string
	HttpClient *http_utils.HTTPClient
}

func NewHTTPDemoClient(id string) *HTTPDemoClient {
	return &HTTPDemoClient{
		ID:         id,
		HttpClient: http_utils.NewHTTPClient(),
	}
}

func (c *HTTPDemoClient) Receive() {
	for {
		c.HttpClient.SendRequest(http_utils.NewHTTPRequest("a", "http://www.example.com", nil, nil))
		time.Sleep(1 * time.Second)
	}
}

func (c *HTTPDemoClient) HTTPSend(r *http_utils.HTTPRequest) ([]byte, error) {
	return c.HttpClient.SendRequest(r)
}
