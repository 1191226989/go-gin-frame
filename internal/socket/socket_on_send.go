package socket

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func (s *server) OnSend(message []byte) error {
	err := s.socket.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		s.OnClose()
		logrus.Error("socket on send error", err)
	}
	return err
}
