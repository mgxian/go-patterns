package singleton

import (
	"testing"

	"unsafe"

	"sync"

	"github.com/stretchr/testify/assert"
)

func TestGetLoadBalancer(t *testing.T) {
	var lba *loadBalancer
	var lbb *loadBalancer
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		lba = getLoadBalancer()
	}()

	go func() {
		defer wg.Done()
		lbb = getLoadBalancer()
	}()

	wg.Wait()
	assert.Equal(t, unsafe.Pointer(lba), unsafe.Pointer(lbb))
}

func TestUseLoadBalancer(t *testing.T) {
	var lba *loadBalancer
	var lbb *loadBalancer
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		lba = getLoadBalancer()
		lba.addServer("server1")
		lba.addServer("server2")
		lba.addServer("server3")
	}()

	go func() {
		defer wg.Done()
		lbb = getLoadBalancer()
		lbb.addServer("server4")
		lbb.addServer("server5")
		lbb.addServer("server6")
	}()

	wg.Wait()
	assert.Equal(t, lba, lbb)

	t.Log(lba.getServer())
	t.Log(lba.getServer())
	t.Log(lbb.getServer())
	t.Log(lbb.getServer())
}
