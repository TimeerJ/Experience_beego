package session

import "go/ast"

// ManagerConfig define the session config
type ManagerConfig struct {
	CookieName			string  `json:"cookieName"`
	EnableSetCookie		bool	`json:"enableSetCookie, omitempty"`

}

func NewManager(provideName string, cf *ManagerConfig) (*Manager, error) {

}

// it can do gc in times after gc lifetime.
func (manager *Manager) GC() {

}
