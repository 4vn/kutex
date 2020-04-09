package kutex

import (
	"hash/crc32"
	"sync"
)

type Kutex struct {
	mus  []sync.Mutex
	size uint
}

func New(size uint) *Kutex {
	return &Kutex{size: size, mus: make([]sync.Mutex, size, size)}
}

func (me *Kutex) Lock(key string) {
	i := uint(crc32.ChecksumIEEE([]byte(key)))
	me.mus[i%me.size].Lock()
}

func (me *Kutex) Unlock(key string) {
	i := uint(crc32.ChecksumIEEE([]byte(key)))
	me.mus[i%me.size].Unlock()
}
