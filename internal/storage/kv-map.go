package storage

import (
	"errors"
	"sync"
)

var ErrorNoSuchKey = errors.New("No such key")

type lockableMap struct {
	m map[string]string
	sync.RWMutex
}

func NewLockableMap() *lockableMap {
	return &lockableMap{
		m: make(map[string]string),
	}
}

func (lm *lockableMap) Get(key string) (string, error) {
	lm.RLock()
	defer lm.Unlock()

	value, ok := lm.m[key]
	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func (lm *lockableMap) Put(key string, value string) error {
	lm.Lock()
	defer lm.Unlock()

	lm.m[key] = value
	return nil
}

func (lm *lockableMap) Delete(key string) error {
	lm.Lock()
	defer lm.Unlock()

	delete(lm.m, key)
	return nil
}
