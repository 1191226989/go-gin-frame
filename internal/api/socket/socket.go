package socket

import (
	"errors"
	"go-gin-frame/internal/socket"
	"math/rand"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	serverMap sync.Map
)

// 获取连接
func GetConn(userId int) (socket.Server, error) {
	conn, ok := serverMap.Load(userId)
	if !ok {
		return nil, errors.New("conn is nil")
	}
	s := conn.(socket.Server)
	return s, nil
}

// 广播
func Broadcast(message []byte) {
	serverMap.Range(func(key, value interface{}) bool {
		s := value.(socket.Server)
		s.OnSend(message)
		return true
	})
}

// 创建链接
func Connect(c *gin.Context) {
	// token 验证，用户数据
	server, err := socket.New(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// 保存客户端链接
	userId := rand.Intn(100)
	serverMap.Store(userId, server)

	go server.OnMessage()
}
