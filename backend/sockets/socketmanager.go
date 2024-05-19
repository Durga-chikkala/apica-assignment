package sockets

import (
	"encoding/json"
	"github.com/Durga-chikkala/apica-assignment/models"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	connections []*websocket.Conn
	mutex       sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		connections: []*websocket.Conn{},
	}
}

// UpgradeHandler upgrades HTTP to WebSocket and stores the connection
func (m *Manager) UpgradeHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}

	m.mutex.Lock()
	m.connections = append(m.connections, conn)
	m.mutex.Unlock()

	go m.handleConnection(conn)
}

// handleConnection reads messages from a WebSocket connection
func (m *Manager) handleConnection(conn *websocket.Conn) {
	defer func() {
		m.mutex.Lock()
		for i, c := range m.connections {
			if c == conn {
				m.connections = append(m.connections[:i], m.connections[i+1:]...)
				break
			}
		}

		m.mutex.Unlock()
		conn.Close()
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}
	}
}

// Broadcast sends a message to all WebSocket connections
func (m *Manager) Broadcast(cache map[string]*models.CacheData) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	message, err := json.Marshal(cache)
	if err != nil {
		log.Println("Error marshaling data to JSON:", err)
		return
	}

	for _, conn := range m.connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error sending WebSocket message:", err)
			conn.Close()
		}
	}
}
