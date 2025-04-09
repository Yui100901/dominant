package global

import (
	"dominant/domain/impl/data/common"
	"github.com/Yui100901/MyGo/concurrency"
)

//
// @Author yfy2001
// @Date 2025/1/8 20 43
//

var DeviceMap *concurrency.SafeMap[string, *common.Device]

func init() {
	DeviceMap = concurrency.NewSafeMap[string, *common.Device](32)
}
