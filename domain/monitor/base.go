package monitor

import "time"

//
// @Author yfy2001
// @Date 2024/8/1 12 41
//

// MainMonitor 全局监视器负责检测节点状态
var MainMonitor *Monitor

func init() {
	MainMonitor = NewMonitor(120 * time.Minute)
	go MainMonitor.Start()

}

func GetNodesList() ([]string, []string) {
	return MainMonitor.GetOnlineNodes(), MainMonitor.GetOfflineNodes()
}

func GetNodeStatus(id string) bool {
	return MainMonitor.GetOnlineStatusById(id)
}
