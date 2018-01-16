package cloudfinder

import (
	"sort"
	"sync"
	
	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
)

var (
	registryMu sync.Mutex
	registry   = make(map[string]provider.Provider)
)

// Register a provider. Code is adapted from:
// https://golang.org/src/database/sql/sql.go?s=1028:1076#L33
func Register(name string, p provider.Provider) {
	registryMu.Lock()
	defer registryMu.Unlock()
	if p == nil {
		panic("cloudfinder provider: Register provider is nil")
	}
	if _, dup := registry[name]; dup {
		panic("cloudfinder provider: Register called twice for provider " + name)
	}
	registry[name] = p
}

// Providers returns a list of string
func Providers() []string {
	registryMu.Lock()
	defer registryMu.Unlock()

	var list []string
	for name := range registry {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}
