package ws

import "sync"

type HubManager struct {
	Rooms map[string]*Hub
	Mu    sync.Mutex
}

var manager = &HubManager{
	Rooms: make(map[string]*Hub),
}

func GetHub(room string) *Hub {
	manager.Mu.Lock()
	defer manager.Mu.Unlock()

	if hub, ok := manager.Rooms[room]; ok {
		return hub
	}

	newHub := NewHub()
	manager.Rooms[room] = newHub
	go newHub.Run()
	return newHub
}

func GetActiveRoomsMap() map[string]int {
	manager.Mu.Lock()
	defer manager.Mu.Unlock()

	active := make(map[string]int)
	for room, hub := range manager.Rooms {
		if len(hub.Clients) > 0 {
			active[room] = len(hub.Clients)
		}
	}
	return active
}
