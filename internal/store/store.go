package store

import (
	"errors"
	"time"
)

type Store struct {
	Data    map[string]interface{}
	MaxSize int
}

func New(maxSize int) *Store {
	return &Store{
		Data:    make(map[string]interface{}),
		MaxSize: maxSize,
	}
}

// Принимает ключ, значение и время жизни записи, создает запись
func (s *Store) Set(key string, value interface{}, ttl int) error {
	if len(s.Data) < s.MaxSize {
		s.Data[key] = value
		if ttl != 0 {
			liveTime := time.Duration(ttl) * time.Second
			time.AfterFunc(liveTime, func() {
				delete(s.Data, key)
			})
		}
		return nil
	}
	return errors.New("storage is full")
}

// Возвращает значение для заданного ключа
func (s *Store) Get(key string) interface{} {
	return s.Data[key]
}

// Возвращает список всех существующих ключей
func (s *Store) GetAllKeys() []string {
	keys := make([]string, 0, len(s.Data))
	for k := range s.Data {
		keys = append(keys, k)
	}
	return keys
}

// Удаляет запись по ключу
func (s *Store) Delete(key string) {
	delete(s.Data, key)
}
