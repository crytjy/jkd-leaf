package network

import (
	"sync"
)

var (
	requestID uint64
	mu        sync.Mutex
)

// SetRequestID 设置全局请求ID
func SetRequestID(id uint64) {
	mu.Lock()
	defer mu.Unlock()
	requestID = id
}

// GetRequestID 获取全局请求ID
func GetRequestID() uint64 {
	mu.Lock()
	defer mu.Unlock()

	return requestID
}
