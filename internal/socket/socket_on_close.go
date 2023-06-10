package socket

import (
	"github.com/sirupsen/logrus"
)

func (s *server) OnClose() {
	err := s.socket.Close()
	if err != nil {
		logrus.Error("socket on closed error", err)
	}
}
