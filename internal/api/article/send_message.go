package article

import (
	"encoding/json"
	"go-gin-frame/internal/api/socket"
	"go-gin-frame/pkg/timeutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sendMessageRequest struct {
	Message string `json:"message"` // 消息内容
}

type sendMessageResponse struct {
	Status string `json:"status"` // 状态
}

type messageBody struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func SendMessage(c *gin.Context) {
	req := new(sendMessageRequest)
	res := new(sendMessageResponse)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithError(
			http.StatusBadRequest,
			err,
		)
		return
	}

	messageData := new(messageBody)
	messageData.Message = req.Message
	messageData.Time = timeutil.CSTLayoutString()

	messageJsonData, err := json.Marshal(messageData)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 广播消息
	socket.Broadcast(messageJsonData)

	res.Status = "OK"
	c.JSON(200, res)
}
