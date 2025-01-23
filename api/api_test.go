package api

import (
	"github.com/Yui100901/MyGo/network/http_utils"
	"testing"
)

//
// @Author yfy2001
// @Date 2025/1/21 13 48
//

func TestDeviceMessage(t *testing.T) {
	hc := http_utils.NewHTTPClient()
	res, _ := hc.SendRequest(
		http_utils.NewHTTPRequest(
			"GetByQuery",
			"http://127.0.0.1:27777/api/v3/telemetry/dataInfo",
			nil,
			nil))
	t.Log(string(res))
}
