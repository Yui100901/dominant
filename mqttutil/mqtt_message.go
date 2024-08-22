package mqttutil

import (
	"github.com/google/uuid"
	"time"
)

//
// @Author yfy2001
// @Date 2024/8/18 10 40
//

type MQTTMessage struct {
	ID        string `json:"id"`
	NodeId    string `json:"nodeId"`
	Time      string `json:"time"`
	Telemetry any    `json:"telemetry"`
}

func NewMQTTMessage(nodeId string, telemetry any) *MQTTMessage {
	return &MQTTMessage{
		ID:        uuid.NewString(),
		NodeId:    nodeId,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		Telemetry: telemetry,
	}
}
