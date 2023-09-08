package hashtables

import (
	"errors"
	"fmt"
	"hash/fnv"
)

type ChainingHashTable[T comparable] struct {
	table  []*chain[T]
	keyLen int
}

type chain[T comparable] struct {
	Key   string
	Value T
	Next  *chain[T]
}

func NewChainingHT[T comparable]() *ChainingHashTable[T] {
	return &ChainingHashTable[T]{table: make([]*chain[T], 8), keyLen: 0}
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
		ht.keyLen++

		if ht.keyLen > len(ht.table) {
			ht.grow()
		}
		return
	}

	for item := ht.table[i]; item != nil; item = item.Next {
		if item.Key == k {
			item.Value = v
			return
		}

		if item.Next == nil {
			item.Next = &chain[T]{k, v, nil}
			ht.keyLen++
			if ht.keyLen > len(ht.table) {
				ht.grow()
			}
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
			ht.keyLen--
			if ht.keyLen < len(ht.table)/4 {
				ht.shrink()
			}
			return value, nil
		}
	}

	var null T

	return null, errors.New("key not found")
}

func (ht *ChainingHashTable[T]) Print() {
	for i, item := range ht.table {
		fmt.Printf("%d: ", i)
		for ; item != nil; item = item.Next {
			fmt.Printf("(%s, %v) ", item.Key, item.Value)
		}
		fmt.Println()
	}
}

func (ht *ChainingHashTable[T]) grow() {
	newHt := make([]*chain[T], len(ht.table)*2)

	for _, item := range ht.table {
		for ; item != nil; item = item.Next {
			h := fnv.New32a()
			h.Write([]byte(item.Key))
			i := h.Sum32() % uint32(len(newHt))

			if newHt[i] == nil {
				newHt[i] = &chain[T]{item.Key, item.Value, nil}
				continue
			}

			for newItem := newHt[i]; newItem != nil; newItem = newItem.Next {
				if newItem.Next == nil {
					newItem.Next = &chain[T]{item.Key, item.Value, nil}
					break
				}
			}
		}
	}

	ht.table = newHt
}

func (ht *ChainingHashTable[T]) shrink() {
	newHt := make([]*chain[T], len(ht.table)/2)

	for _, item := range ht.table {
		for ; item != nil; item = item.Next {
			h := fnv.New32a()
			h.Write([]byte(item.Key))
			i := h.Sum32() % uint32(len(newHt))

			if newHt[i] == nil {
				newHt[i] = &chain[T]{item.Key, item.Value, nil}
				continue
			}

			for newItem := newHt[i]; newItem != nil; newItem = newItem.Next {
				if newItem.Next == nil {
					newItem.Next = &chain[T]{item.Key, item.Value, nil}
					break
				}
			}
		}
	}

	ht.table = newHt
}

func (ht *ChainingHashTable[T]) LoadFactor() float32 {
	return float32(ht.keyLen) / float32(len(ht.table))
}
