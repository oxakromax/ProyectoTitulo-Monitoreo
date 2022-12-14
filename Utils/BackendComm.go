package Utils

import (
	"sync"
)

type Folders struct {
	FoldersID []int
	mu        sync.Mutex
}

func (f *Folders) Set(value []int) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.FoldersID = value
}

func (f *Folders) Get() []int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.FoldersID
}

func (f *Folders) Add(value int) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.FoldersID = append(f.FoldersID, value)
}

func (f *Folders) Delete(value int) {
	f.mu.Lock()
	defer f.mu.Unlock()
	for i, v := range f.FoldersID {
		if v == value {
			f.FoldersID = append(f.FoldersID[:i], f.FoldersID[i+1:]...)
		}
	}
}
