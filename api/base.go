package api

import (
	"dominant/server"
	"github.com/gin-gonic/gin"
)

//
// @Author yfy2001
// @Date 2025/1/21 15 55
//

var WebEngine *gin.Engine

func init() {
	WebEngine = server.NewServer()
}
