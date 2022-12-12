package Utils

import (
	"sync"
	"time"
)

type RefreshingBool struct {
	IsRefreshingToken bool
	mu                sync.Mutex
}

func (r *RefreshingBool) Set(value bool) {
	r.mu.Lock()
	r.IsRefreshingToken = value
	r.mu.Unlock()
}

func (r *RefreshingBool) Get() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.IsRefreshingToken
}

var (
	LastMonitoredTime time.Time      // LastMonitoredTime almacena el tiempo desde el último monitoreo realizado.
	IsRefreshingToken RefreshingBool // IsRefreshingToken indica si se está refrescando el token de acceso a la API de UiPath.
	// Al ser una variable global, que se consulta por diferentes gorutinas, es necesario protegerla con un mutex.
)
