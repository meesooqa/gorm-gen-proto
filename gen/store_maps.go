package gen

import (
	"encoding/json"
	"os"
	"sync"
)

// Using:
/*
	store, _ := NewStore("data.json")
	store.mu.Lock()
	store.data["key"] = "value"
	store.mu.Unlock()
	_ = store.Save("data.json")
*/

// Store stores "key => value" data
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewStore(filename string) (*Store, error) {
	s := &Store{data: make(map[string]string)}
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &s.data)
	if err != nil {
		return nil, err
	}
	return s, nil
}

//func (o *Store) Save(filename string) error {
//	o.mu.RLock()
//	defer o.mu.RUnlock()
//	data, err := json.Marshal(o.data)
//	if err != nil {
//		return err
//	}
//	return os.WriteFile(filename, data, 0644)
//}
