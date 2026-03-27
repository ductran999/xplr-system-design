package master

import (
	pb "master-slave/api/gen/pb/agent/v1"
	"sync"
)

type ConnectionManager struct {
	mu      sync.RWMutex
	tunnels map[string]chan *pb.ConnectTunnelResponse
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		tunnels: make(map[string]chan *pb.ConnectTunnelResponse),
	}
}

func (cm *ConnectionManager) Register(clusterID string, cmdChan chan *pb.ConnectTunnelResponse) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.tunnels[clusterID] = cmdChan
}

func (cm *ConnectionManager) Unregister(clusterID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if ch, ok := cm.tunnels[clusterID]; ok {
		close(ch)
	}
	delete(cm.tunnels, clusterID)
}

func (cm *ConnectionManager) SendCommand(clusterID string, cmd *pb.ConnectTunnelResponse) bool {
	cm.mu.RLock()
	ch, ok := cm.tunnels[clusterID]
	cm.mu.RUnlock()

	if !ok {
		return false
	}

	select {
	case ch <- cmd:
		return true
	default:
		return false
	}
}
