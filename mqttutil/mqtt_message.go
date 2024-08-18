package mqttutil

//
// @Author yfy2001
// @Date 2024/8/18 10 40
//

type MqttMessage struct {
	ID        string `json:"id"`
	NodeId    string `json:"nodeId"`
	Telemetry any    `json:"telemetry"`
}
