package singleton

import (
	"math/rand"
	"sync"
)

type loadBalancer struct {
	serverList []string
}

func newLoadBalancer() *loadBalancer {
	return &loadBalancer{
		serverList: make([]string, 0),
	}
}

var (
	loadBalancerInstance   *loadBalancer
	createLoadBalancerLock sync.Mutex
)

// double check locking
func getLoadBalancer() *loadBalancer {
	if loadBalancerInstance != nil {
		return loadBalancerInstance
	}

	createLoadBalancerLock.Lock()
	defer createLoadBalancerLock.Unlock()

	if loadBalancerInstance != nil {
		return loadBalancerInstance
	}

	loadBalancerInstance = newLoadBalancer()
	return loadBalancerInstance
}

func (lb *loadBalancer) addServer(server string) {
	lb.serverList = append(lb.serverList, server)
}

func (lb *loadBalancer) removeServer(server string) {
}

func (lb *loadBalancer) getServer() string {
	index := rand.Intn(len(lb.serverList))
	return lb.serverList[index]
}
