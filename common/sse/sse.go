package sse

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrClientNotFound = errors.New("client not found")
	EventHeartbeat    = "heartbeat"
)

type Message struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

type client struct {
	id          string
	messageChan chan Message
	closeChan   chan struct{}
}

type SSE interface {
	Handler() gin.HandlerFunc
	Broadcast(msg Message)
	Send(clientID string, msg Message) error
	Close()
	ClientCount() int
}

type sseImpl struct {
	clients map[string]*client
	mu      sync.RWMutex
}

func NewSSE() SSE {
	return &sseImpl{
		clients: make(map[string]*client),
	}
}

func (s *sseImpl) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建客户端连接
		clientID := uuid.New().String()
		newClient := &client{
			id:          clientID,
			messageChan: make(chan Message, 10), // 带缓冲的通道
			closeChan:   make(chan struct{}),
		}

		// 注册客户端
		s.mu.Lock()
		s.clients[clientID] = newClient
		s.mu.Unlock()

		// 设置SSE响应头
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 立即发送clientID给客户端
		c.SSEvent("init", clientID)
		c.Writer.Flush()

		// 心跳定时器
		heartbeat := time.NewTicker(10 * time.Second)

		defer func() {
			heartbeat.Stop()
			s.removeClient(clientID)
			close(newClient.messageChan)
			close(newClient.closeChan)
		}()

		// 处理连接关闭
		go func() {
			<-c.Request.Context().Done()
			s.removeClient(clientID)
			newClient.closeChan <- struct{}{}
		}()

		// 主循环
		for {
			select {
			case <-heartbeat.C:
				// 发送心跳保持连接
				c.SSEvent(EventHeartbeat, "ping")
				c.Writer.Flush()

			case msg := <-newClient.messageChan:
				// 发送消息给客户端
				c.SSEvent(msg.Event, msg.Data)
				c.Writer.Flush()

			case <-newClient.closeChan:
				// 关闭连接
				return
			}
		}
	}
}

func (s *sseImpl) Broadcast(msg Message) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, client := range s.clients {
		select {
		case client.messageChan <- msg:
		default:
			// 防止消息阻塞，丢弃无法发送的消息
			log.Printf("client %s message channel full, message dropped", client.id)
		}
	}
}

func (s *sseImpl) Send(clientID string, msg Message) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	client, exists := s.clients[clientID]
	if !exists {
		return ErrClientNotFound
	}

	select {
	case client.messageChan <- msg:
	default:
		log.Printf("client %s message channel full, message dropped", clientID)
	}
	return nil
}

func (s *sseImpl) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, client := range s.clients {
		client.closeChan <- struct{}{}
		delete(s.clients, id)
	}
}

func (s *sseImpl) ClientCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.clients)
}

func (s *sseImpl) removeClient(clientID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, clientID)
}
