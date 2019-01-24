// Copyright (c), Ruslan Sendecky. All rights reserved
// Use of this source code is governed by the MIT license
// See the LICENSE file in the project root for more information

package gsession

import (
	"sync"
	"time"
)

// MemoryStore struct
type MemoryStore struct {
	sync.RWMutex
	shelf map[string]*Session
}

// NewMemoryStore creates a new memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		shelf: make(map[string]*Session),
	}
}

// Create adds a new session entry to the store
// Takes a session ID
func (s *MemoryStore) Create(id string) error {
	s.Lock()
	defer s.Unlock()
	s.shelf[id] = &Session{
		Origin: time.Now(),
		Tstamp: time.Now(),
		Token:  "",
		Data:   make(map[string]interface{}),
	}
	return nil
}

// Read retrieves Session from store
// Takes session ID
// If session not found returns ErrSessionNoRecord error
func (s *MemoryStore) Read(id string) (*Session, error) {
	s.RLock()
	defer s.RUnlock()
	if ses, ok := s.shelf[id]; ok {
		scp := *ses
		return &scp, nil
	}
	return nil, ErrSessionNoRecord
}

// Update runs a function on Session
// Takes session ID and a function with Session as parameter
// If session not found returns ErrSessionNoRecord error
func (s *MemoryStore) Update(id string, fn func(*Session)) (err error) {
	s.Lock()
	defer s.Unlock()
	if ses, ok := s.shelf[id]; ok {
		fn(ses)
		return nil
	}
	return ErrSessionNoRecord
}

// Delete removes Session from the store
// Takes session ID
func (s *MemoryStore) Delete(id string) error {
	s.Lock()
	defer s.Unlock()
	delete(s.shelf, id)
	return nil
}
