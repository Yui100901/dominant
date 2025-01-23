package access

import "dominant/domain/impl/access"

//
// @Author yfy2001
// @Date 2025/1/15 09 53
//

var NC *NodeClient

var ShipC *access.ShipClient

func init() {
	NC = NewNodeClient("123")
	NC.Receive()

	ShipC = access.NewShipClient("ship")
	ShipC.Receive()
	//time.Sleep(3000 * time.Millisecond)
	//s.Disconnect()
}
