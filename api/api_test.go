package api

import (
	"github.com/Yui100901/MyGo/network/http_utils"
	"testing"
)

//
// @Author yfy2001
// @Date 2025/1/21 13 48
//

const (
	IP      = "127.0.0.1"
	PORT    = "28888"
	ADDR    = IP + ":" + PORT
	BaseUrl = "http://" + IP + ":" + PORT
)

func TestDeviceMessage(t *testing.T) {
	apiUrl := BaseUrl + UrlPrefix + "/telemetry/dataInfo"
	hc := http_utils.NewHTTPClient()
	res, _ := hc.SendRequest(
		http_utils.NewHTTPRequest(
			"GET",
			apiUrl,
			nil,
			nil,
			"",
			nil))
	t.Log(string(res))
}
