package hashtables

import (
	"errors"
	"hash/fnv"
)

type ChainingHashTable[T comparable] struct {
	table []*chain[T]
}

type chain[T comparable] struct {
	Key   string
	Value T
	Next  *chain[T]
}

func NewChainingHT[T comparable]() *ChainingHashTable[T] {
	return &ChainingHashTable[T]{make([]*chain[T], 8)}
}

func (ht *ChainingHashTable[T]) hashFunction(k string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(k))
	return uint32(h.Sum32() % uint32(len(ht.table)))
}

func (ht *ChainingHashTable[T]) Add(k string, v T) {
	i := ht.hashFunction(k)

	if ht.table[i] == nil {
		ht.table[i] = &chain[T]{k, v, nil}
		return
	}

	for item := ht.table[i]; item != nil; item = item.Next {
		if item.Next == nil {
			item.Next = &chain[T]{k, v, nil}
			return
		}
	}
}

func (ht *ChainingHashTable[T]) Get(k string) (T, error) {
	i := ht.hashFunction(k)

	for item := ht.table[i]; item != nil; item = item.Next {
		if item.Key == k {
			return item.Value, nil
		}
	}

	var null T

	return null, errors.New("key not found")
}

func (ht *ChainingHashTable[T]) Del(k string) (T, error) {
	i := ht.hashFunction(k)

	for item := &ht.table[i]; *item != nil; item = &(*item).Next {
		if (**item).Key == k {
			value := (*item).Value
			*item = (*item).Next
			return value, nil
		}
	}

	var null T

	return null, errors.New("key not found")
}

func (ht *ChainingHashTable[T]) Set(k string, v T) error {
	i := ht.hashFunction(k)

	for item := ht.table[i]; item != nil; item = item.Next {
		if item.Key == k {
			item.Value = v
			return nil
		}
	}

	return errors.New("key not found")
}
