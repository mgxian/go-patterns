package singleton

import "sync"

// Singleton describes a struct of only a single instance can exist.
type Singleton struct {
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns the single instance.
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
