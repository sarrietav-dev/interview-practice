package hashtables

type HashTable[T comparable] interface {
	Add(k string, v T)
	Get(k string) (T, error)
	Del(k string) (T, error)
}
