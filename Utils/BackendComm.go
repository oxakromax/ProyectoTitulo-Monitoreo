package Utils

import (
	"golang.org/x/exp/slices"
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

type ProcesosBDD struct {
	Nombre           string `json:"Nombre"`
	Folderid         int    `json:"Folderid"`
	WarningTolerance int    `json:"WarningTolerance"`
	ErrorTolerance   int    `json:"ErrorTolerance"`
	FatalTolerance   int    `json:"FatalTolerance"`
}

type ProcessBDDArray struct {
	Processes []ProcesosBDD `json:"Processes"`
	mu        sync.Mutex
}

func (p *ProcessBDDArray) Set(value []ProcesosBDD) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Processes = value
}

func (p *ProcessBDDArray) Get() []ProcesosBDD {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.Processes
}

func (p *ProcessBDDArray) Add(value ProcesosBDD) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Processes = append(p.Processes, value)
}

func (p *ProcessBDDArray) Delete(value ProcesosBDD) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for i, v := range p.Processes {
		if v == value {
			p.Processes = append(p.Processes[:i], p.Processes[i+1:]...)
		}
	}
}

func (p *ProcessBDDArray) FilterUniqueFoldersID() []int {
	p.mu.Lock()
	defer p.mu.Unlock()
	var foldersID []int
	for _, v := range p.Processes {
		if !slices.Contains(foldersID, v.Folderid) {
			foldersID = append(foldersID, v.Folderid)
		}
	}
	return foldersID
}
