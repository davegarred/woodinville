package storage

import (
	"sync"
	"fmt"
)

type IdentifierType string

type IdentifierFactory struct {
	sync.Mutex
	idType IdentifierType
	last uint64
}

func NewIdentifierFactory(idType IdentifierType) *IdentifierFactory {
	return &IdentifierFactory{
		idType: idType,
		last:   10000,
	}
}

func (f *IdentifierFactory) Next() string {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()
	f.last += 1
	return fmt.Sprintf("%s%d", f.idType, f.last)
}
