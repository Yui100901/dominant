package global

import (
	"dominant/domain/impl/data/common"
	"github.com/Yui100901/MyGo/concurrent"
)

//
// @Author yfy2001
// @Date 2025/1/8 20 43
//

var DeviceMap *concurrent.SafeMap[string, *common.Device]

func init() {
	DeviceMap = concurrent.NewSafeMap[string, *common.Device](32)
}
