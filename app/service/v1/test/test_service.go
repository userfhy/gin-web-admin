package testService

import (
	"sync"
	"time"

	"gin-web-admin/common/sse"
	"gin-web-admin/utils/logging"
	"gin-web-admin/utils/system_monitor"

	"github.com/gin-gonic/gin"
)

type Service struct {
	server      sse.SSE
	monitorOnce sync.Once
}

func NewService() *Service {
	return &Service{
		server: sse.NewSSE(
			sse.WithConfig(sse.Config{
				HeartbeatInterval: 10 * time.Second,
				ChannelBufferSize: 50,
				WriteTimeout:      30 * time.Second,
				MaxClientMessages: 10000,
			}),
			sse.WithSendFailHandler(func(clientID string, msg sse.Message) {
				logging.Printf("重要消息发送失败: client=%s event=%s", clientID, msg.Event)
			}),
			sse.WithLogger(&sse.LogrusAdapter{}),
		),
	}
}

func (s *Service) SSEHandler() gin.HandlerFunc {
	return s.server.Handler()
}

func (s *Service) Broadcast(msg sse.Message) {
	s.server.Broadcast(msg)
}

func (s *Service) Send(clientID string, msg sse.Message) error {
	return s.server.Send(clientID, msg)
}

func (s *Service) ClientCount() int {
	return s.server.ClientCount()
}

func (s *Service) StartSystemStatsBroadcast() {
	s.monitorOnce.Do(func() {
		go func() {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()
			for range ticker.C {
				stats, err := system_monitor.GetSystemStats()
				if err != nil {
					logging.Warn("获取系统状态失败: ", err)
					continue
				}
				s.Broadcast(sse.Message{
					Event: "system_status",
					Data: gin.H{
						"systemStats": stats.String(),
						"clients":     s.ClientCount(),
					},
				})
			}
		}()
	})
}
