package kutex

import (
	"hash/crc32"
	"sync"
)

type Kutex struct {
	mus  []sync.Mutex
	size uint
}

// New creates a Kutex instance.
func New(size uint) *Kutex {
	return &Kutex{size: size, mus: make([]sync.Mutex, size, size)}
}

// Lock waits to acquire lock.
func (me *Kutex) Lock(key string) {
	i := uint(crc32.ChecksumIEEE([]byte(key)))
	me.mus[i%me.size].Lock()
}

// Unlock release the lock.
func (me *Kutex) Unlock(key string) {
	i := uint(crc32.ChecksumIEEE([]byte(key)))
	me.mus[i%me.size].Unlock()
}
