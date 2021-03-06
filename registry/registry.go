package registry

import (
	"errors"
	"fmt"
)

// Result is returned by a call to Next on
// the watcher. Actions can be create, update, delete
type Result ServiceURLEvent

func (r Result) String() string {
	return fmt.Sprintf("Result{Action{%s}, Service{%#v}}", r.Action.String(), r.Service)
}

// Watcher is an interface that returns updates
// about services within the registry.
type Watcher interface {
	// Next is a blocking call
	Next() (*Result, error)
	Valid() bool // 检查watcher与registry连接是否正常
	Stop()
}

// The registry provides an interface for service discovery
// and an abstraction over varying implementations
// {consul, etcd, zookeeper, ...}
type Registry interface {
	// Register(conf ServiceConfig) error
	Register(conf interface{}) error
	GetService(string) ([]*ServiceURL, error)
	ListServices() ([]*ServiceURL, error)
	Watch() (Watcher, error)
	Close()
	String() string
}

const (
	REGISTRY_CONN_DELAY = 3 // watchDir中使用，防止不断地对zk重连
)

var (
	ErrorRegistryNotFound = errors.New("registry not found")
)
