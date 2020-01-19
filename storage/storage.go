package storage

import (
	"sort"
	"sync"
)

var instance *storage
var once sync.Once

type TAnagram string

func (a TAnagram) Key() string {
	ab := []byte(a)
	sort.Slice(ab, func(i, j int) bool { return ab[i] < ab[j] })
	return string(ab)
}

type storage struct {
	MTX sync.RWMutex

	AnagramsR map[string][]TAnagram
	AnagramsW map[string]map[TAnagram]bool
}

func Get() *storage {
	once.Do(func() {
		instance = &storage{
			AnagramsW: make(map[string]map[TAnagram]bool, 0),
			AnagramsR: make(map[string][]TAnagram, 0),
		}
	})
	return instance
}
