package store

import "time"

type Store struct {
	Data map[string]interface{}
}

func New() *Store {
	return &Store{
		Data: make(map[string]interface{}),
	}
}

// Принимает ключ, значение и время жизни записи, создает запись
func (s *Store) Set(key string, value interface{}, ttl int) {
	s.Data[key] = value
	if ttl != 0 {
		liveTime := time.Duration(ttl) * time.Second
		time.AfterFunc(liveTime, func() {
			delete(s.Data, key)
		})
	}
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
