package sse

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrClientNotFound   = errors.New("client not found")
	ErrMessageDropped   = errors.New("message dropped due to full channel")
	EventHeartbeat      = "heartbeat"
	EventClientRegister = "register"
	DefaultConfig       = Config{
		HeartbeatInterval: 10 * time.Second,
		ChannelBufferSize: 50,
		WriteTimeout:      30 * time.Second,
		MaxClientMessages: 10000,
	}
)

type (
	Message struct {
		Event string `json:"event"`
		Data  any    `json:"data"`
	}

	client struct {
		sync.Mutex
		id           string
		messageChan  chan Message
		closeChan    chan struct{}
		lastActive   time.Time
		messageCount uint
	}

	SSE interface {
		Handler() gin.HandlerFunc
		Broadcast(msg Message)
		Send(clientID string, msg Message) error
		Close()
		ClientCount() int
		GetClientIDs() []string
	}

	sseImpl struct {
		clients    map[string]*client
		mu         sync.RWMutex
		config     Config
		logger     Logger
		onSendFail func(clientID string, msg Message)
	}

	Config struct {
		HeartbeatInterval time.Duration
		ChannelBufferSize int
		WriteTimeout      time.Duration
		MaxClientMessages uint
	}

	Logger interface {
		Printf(format string, v ...any)
	}

	Option func(*sseImpl)
)

// NewSSE 创建SSE实例，支持可选配置
func NewSSE(opts ...Option) SSE {
	s := &sseImpl{
		clients: make(map[string]*client),
		config:  DefaultConfig,
		logger:  defaultLogger{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithConfig 自定义配置选项
func WithConfig(cfg Config) Option {
	return func(s *sseImpl) {
		s.config = cfg
	}
}

// WithLogger 自定义日志记录器
func WithLogger(l Logger) Option {
	return func(s *sseImpl) {
		s.logger = l
	}
}

// WithSendFailHandler 设置消息发送失败回调
func WithSendFailHandler(fn func(clientID string, msg Message)) Option {
	return func(s *sseImpl) {
		s.onSendFail = fn
	}
}

func (s *sseImpl) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := uuid.NewString()
		newClient := &client{
			id:          clientID,
			messageChan: make(chan Message, s.config.ChannelBufferSize),
			closeChan:   make(chan struct{}, 1),
			lastActive:  time.Now(),
		}

		s.registerClient(clientID, newClient)

		defer s.cleanupClient(clientID, newClient)

		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		if err := s.sendInitialData(c, clientID); err != nil {
			return
		}

		heartbeat := time.NewTicker(s.config.HeartbeatInterval)
		defer heartbeat.Stop()

		for {
			select {
			case <-heartbeat.C:
				if !s.handleHeartbeat(c) {
					return
				}

			case msg := <-newClient.messageChan:
				if !s.handleClientMessage(c, newClient, msg) {
					return
				}

			case <-newClient.closeChan:
				return

			case <-c.Request.Context().Done():
				return
			}
		}
	}
}

// Broadcast 广播消息给所有客户端
func (s *sseImpl) Broadcast(msg Message) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for id, client := range s.clients {
		select {
		case client.messageChan <- msg:
			client.messageCount++
		default:
			s.handleMessageDrop(id, msg)
		}

		if client.messageCount >= s.config.MaxClientMessages {
			s.safeCloseClient(id)
		}
	}
}

// Send 发送消息到指定客户端
func (s *sseImpl) Send(clientID string, msg Message) error {
	s.mu.RLock()
	client, exists := s.clients[clientID]
	s.mu.RUnlock()

	if !exists {
		return ErrClientNotFound
	}

	// 单独对客户端加锁
	client.Lock()
	defer client.Unlock()

	select {
	case client.messageChan <- msg:
		client.messageCount++
		if client.messageCount >= s.config.MaxClientMessages {
			go s.safeCloseClient(clientID)
		}
		return nil
	default:
		s.handleMessageDrop(clientID, msg)
		return ErrMessageDropped
	}
}

// Close 安全关闭所有客户端连接
func (s *sseImpl) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id := range s.clients {
		s.closeClient(id)
	}
}

// ClientCount 获取当前活跃客户端数量
func (s *sseImpl) ClientCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.clients)
}

// GetClientIDs 获取所有客户端ID列表
func (s *sseImpl) GetClientIDs() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ids := make([]string, 0, len(s.clients))
	for id := range s.clients {
		ids = append(ids, id)
	}
	return ids
}

// 私有方法实现
func (s *sseImpl) registerClient(clientID string, c *client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[clientID] = c
}

func (s *sseImpl) cleanupClient(clientID string, c *client) {
	s.safeCloseClient(clientID)
	close(c.messageChan)
	close(c.closeChan)
}

func (s *sseImpl) safeCloseClient(clientID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closeClient(clientID)
}

func (s *sseImpl) closeClient(clientID string) {
	if client, exists := s.clients[clientID]; exists {
		select {
		case client.closeChan <- struct{}{}:
		default:
		}
		delete(s.clients, clientID)
	}
}

func (s *sseImpl) sendInitialData(c *gin.Context, clientID string) error {
	c.SSEvent(EventClientRegister, gin.H{
		"client_id": clientID,
		"clients":   s.ClientCount(),
		"timestamp": time.Now().Unix(),
	})

	if err := s.flushWithTimeout(c); err != nil {
		s.logger.Printf("Failed to send initial data: %v", err)
		return err
	}
	return nil
}

func (s *sseImpl) handleHeartbeat(c *gin.Context) bool {
	c.SSEvent(EventHeartbeat, gin.H{
		"timestamp": time.Now().Unix(),
		"clients":   s.ClientCount(),
	})

	if err := s.flushWithTimeout(c); err != nil {
		s.logger.Printf("Heartbeat failed: %v", err)
		return false
	}
	return true
}

func (s *sseImpl) handleClientMessage(c *gin.Context, client *client, msg Message) bool {
	c.SSEvent(msg.Event, msg.Data)
	client.lastActive = time.Now()

	if err := s.flushWithTimeout(c); err != nil {
		s.logger.Printf("Message send failed: %v", err)
		return false
	}
	return true
}

func (s *sseImpl) flushWithTimeout(c *gin.Context) error {
	done := make(chan struct{})
	errChan := make(chan error)

	go func() {
		defer close(done)
		if _, err := c.Writer.Write([]byte("\n")); err != nil {
			errChan <- err
		}
		c.Writer.Flush()
	}()

	select {
	case <-done:
		return nil
	case err := <-errChan:
		return err
	case <-time.After(s.config.WriteTimeout):
		return fmt.Errorf("write timeout")
	}
}

func (s *sseImpl) handleMessageDrop(clientID string, msg Message) {
	s.logger.Printf("Message dropped for client %s (event: %s)", clientID, msg.Event)
	if s.onSendFail != nil {
		go s.onSendFail(clientID, msg)
	}
}

// 默认日志实现
type defaultLogger struct{}

func (d defaultLogger) Printf(format string, v ...any) {
	log.Printf(format, v...)
}
