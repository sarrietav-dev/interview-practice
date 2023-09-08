package hashtables_test

import (
	"fmt"
	"testing"

	hashtables "github.com/sarrietav-dev/interview-practice/hash_tables"
)

func TestAdd(t *testing.T) {
	ht := hashtables.NewChainingHT[int]()
	ht.Add("key1", 1)
	ht.Add("key2", 2)

	val1, err1 := ht.Get("key1")
	if err1 != nil || val1 != 1 {
		t.Errorf("Expected value 1 for key1, got %v with error %v", val1, err1)
	}

	val2, err2 := ht.Get("key2")
	if err2 != nil || val2 != 2 {
		t.Errorf("Expected value 2 for key2, got %v with error %v", val2, err2)
	}
}

func TestAddTwoKeysWithSameHash(t *testing.T) {
	ht := hashtables.NewChainingHT[int]()
	ht.Add("hello", 1)
	ht.Add("world", 2)

	val1, err1 := ht.Get("hello")

	if err1 != nil || val1 != 1 {
		t.Errorf("Expected value 1 for hello, got %v with error %v", val1, err1)
	}

	val2, err2 := ht.Get("world")

	if err2 != nil || val2 != 2 {
		t.Errorf("Expected value 2 for world, got %v with error %v", val2, err2)
	}
}

func TestAddMultipleKeys(t *testing.T) {
	ht := hashtables.NewChainingHT[int]()
	ht.Add("key1", 1)
	ht.Add("key2", 2)
	ht.Add("key3", 3)
	ht.Add("key4", 4)

	val1, err1 := ht.Get("key1")
	if err1 != nil || val1 != 1 {
		t.Errorf("Expected value 1 for key1, got %v with error %v", val1, err1)
	}

	val2, err2 := ht.Get("key2")
	if err2 != nil || val2 != 2 {
		t.Errorf("Expected value 2 for key2, got %v with error %v", val2, err2)
	}

	val3, err3 := ht.Get("key3")
	if err3 != nil || val3 != 3 {
		t.Errorf("Expected value 3 for key3, got %v with error %v", val3, err3)
	}

	val4, err4 := ht.Get("key4")
	if err4 != nil || val4 != 4 {
		t.Errorf("Expected value 4 for key4, got %v with error %v", val4, err4)
	}
}

func TestGetUnknownKey(t *testing.T) {
	var ht hashtables.HashTable[string] = hashtables.NewChainingHT[string]()

	ht.Add("key", "value")

	_, err := ht.Get("notakey")

	if err == nil {
		t.Errorf("Error expected to be returned but it returned nil")
	}
}

func TestKeyDeletion(t *testing.T) {
	var ht hashtables.HashTable[string] = hashtables.NewChainingHT[string]()

	ht.Add("key", "value")

	_, err := ht.Del("key")

	if err != nil {
		t.Errorf("Error expected to be nil but instead %s was returned", err)
	}

	value, err := ht.Get("key")

	fmt.Println(value)

	if err == nil {
		t.Errorf("Error expected to be returned but it returned nil")
	}
}

func TestKeyDeletionWhenKeyDoNotExist(t *testing.T) {
	var ht hashtables.HashTable[string] = hashtables.NewChainingHT[string]()
	_, err := ht.Del("key")

	if err == nil {
		t.Errorf("Error expected to be returned but it returned nil")
	}
}

func TestKeyUpdate(t *testing.T) {
	var ht hashtables.HashTable[string] = hashtables.NewChainingHT[string]()

	ht.Add("key", "value")

	ht.Add("key", "newvalue")

	value, err := ht.Get("key")

	if err != nil {
		t.Errorf("Error expected to be nil but instead %s was returned", err)
	}

	if value != "newvalue" {
		t.Errorf("Expected value to be newvalue but instead %s was returned", value)
	}
}

func TestDeleteOneKeyInAChain(t *testing.T) {
	var ht hashtables.HashTable[string] = hashtables.NewChainingHT[string]()

	ht.Add("hello", "sebas")
	ht.Add("world", "sebas")

	_, err := ht.Del("hello")

	if err != nil {
		t.Errorf("Error expected to be nil but instead %s was returned", err)
	}

	value, err := ht.Get("world")

	if err != nil {
		t.Errorf("Error expected to be nil but instead %s was returned", err)
	}

	if value != "sebas" {
		t.Errorf("Expected value to be sebas but instead %s was returned", value)
	}
}

func TestDeleteMiddleItemInAChain(t *testing.T) {
	var ht hashtables.HashTable[int] = hashtables.NewChainingHT[int]()

	ht.Add("hello", 1)
	ht.Add("world", 2)
	ht.Add("thing>", 3)

	_, err := ht.Del("world")

	if err != nil {
		t.Errorf("Error expected to be nil but instead %s was returned", err)
	}

	_, err = ht.Get("world")

	if err == nil {
		t.Errorf("Error expected to be returned but it returned nil")
	}

	_, err = ht.Get("thing>")

	if err != nil {
		t.Errorf("Error expected to be nil but instead %s was returned", err)
	}

	_, err = ht.Get("hello")

	if err != nil {
		t.Errorf("Error expected to be nil but instead %s was returned", err)
	}
}

func TestPrint(t *testing.T) {
	ht := hashtables.NewChainingHT[int]()

	ht.Add("hello", 1)
	ht.Add("world", 2)
	ht.Add("thing>", 3)
	ht.Add("sebas", 4)
	ht.Add("isa", 5)

	ht.Print()
}

func TestLoadFactor(t *testing.T) {
	ht := hashtables.NewChainingHT[int]()

	ht.Add("hello", 1)
	ht.Add("world", 2)
	ht.Add("thing>", 3)
	ht.Add("sebas", 4)
	ht.Add("isa", 5)
	ht.Add("isa2", 6)
	ht.Add("isa3", 7)
	ht.Add("isa4", 8)
	ht.Add("isa5", 9)

	if ht.LoadFactor() >= 1 {
		t.Errorf("Load factor expected to be less than 1 but instead %f was returned", ht.LoadFactor())
	}
}
