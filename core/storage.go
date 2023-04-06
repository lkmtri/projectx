package core

type Storage interface {
	Put(*Block) error
}

type InMemoryStore struct{}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{}
}

func (s *InMemoryStore) Put(block *Block) error {
	return nil
}
