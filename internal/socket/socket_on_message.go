package socket

import "github.com/sirupsen/logrus"

func (s *server) OnMessage() {
	defer func() {
		s.OnClose()
	}()

	for {
		//接收消息
		_, message, err := s.socket.ReadMessage()
		if err != nil {
			logrus.Error("socket on message error: ", err)
			break
		}

		// 为了便于演示，仅输出到日志文件
		logrus.Infof("receive message: %s ; send from: %s", string(message), s.socket.RemoteAddr().String())
	}
}
