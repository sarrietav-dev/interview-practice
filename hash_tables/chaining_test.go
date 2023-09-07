package hashtables_test

import (
	"fmt"
	"testing"

	hashtables "github.com/sarrietav-dev/interview-practice/hash_tables"
)

func TestAddToChain(t *testing.T) {
	var ht hashtables.HashTable[string] = hashtables.NewChainingHT[string]()

	ht.Add("key", "value")

	val, _ := ht.Get("key")

	if val != "value" {
		t.Errorf("Expected 'value' to be found, but instead %s was found.", val)
	}
}

func TestErrorWhenKeyDoNotExist(t *testing.T) {
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
