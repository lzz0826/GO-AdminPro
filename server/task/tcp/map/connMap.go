package _map

import (
	"fmt"
	"net"
	"sync"
)

// key: userId  value:netConn
var connections = make(map[string]net.Conn)
var mutex = &sync.Mutex{}

// 保存连接
func SaveConnection(userID string, conn net.Conn) {
	mutex.Lock()
	connections[userID] = conn
	mutex.Unlock()
}

// 发送数据
func SendToUser(userID string, message []byte) error {
	mutex.Lock()
	conn, exists := connections[userID]
	mutex.Unlock()

	if exists {
		_, err := conn.Write(message)
		return err
	} else {
		return fmt.Errorf("user ID %s not found", userID)
	}
}

// 删除连接
func RemoveConnection(userID string) {
	mutex.Lock()
	delete(connections, userID)
	mutex.Unlock()
}
