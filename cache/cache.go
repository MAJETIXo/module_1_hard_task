package cache

import "sync"

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	return &Cache{store: make(map[string]string)}
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	c.store[k] = v
	c.mu.Unlock()
}

func (c *Cache) Get(k string) (v string, ok bool) { // Я сделал указатели на кэш без них не получалось!!!
	c.mu.RLock()
	v, ok = c.store[k]
	c.mu.RUnlock()
	return
}
