package main

import (
	"fmt"
	"hash/fnv"
)

type Entry struct {
	Key   string
	Value interface{}
	Next  *Entry
}

type HashTable struct {
	buckets []*Entry
	size    int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*Entry, size),
		size:    size,
	}
}

func (ht *HashTable) hash(key string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % ht.size
}

func (ht *HashTable) Put(key string, value interface{}) {
	index := ht.hash(key)
	entry := ht.buckets[index]

	if entry == nil {
		ht.buckets[index] = &Entry{Key: key, Value: value}
		return
	}

	prev := entry
	for entry != nil {
		if entry.Key == key {
			entry.Value = value
			return
		}
		prev = entry
		entry = entry.Next
	}

	prev.Next = &Entry{Key: key, Value: value}
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := ht.hash(key)
	entry := ht.buckets[index]

	for entry != nil {
		if entry.Key == key {
			return entry.Value, true
		}
		entry = entry.Next
	}
	return nil, false
}

func (ht *HashTable) Remove(key string) bool {
	index := ht.hash(key)
	entry := ht.buckets[index]

	if entry == nil {
		return false
	}

	if entry.Key == key {
		ht.buckets[index] = entry.Next
		return true
	}

	prev := entry
	entry = entry.Next
	for entry != nil {
		if entry.Key == key {
			prev.Next = entry.Next
			return true
		}
		prev = entry
		entry = entry.Next
	}
	return false
}

func (ht *HashTable) Print() {
	for i, bucket := range ht.buckets {
		fmt.Printf("Bucket %d: ", i)
		entry := bucket
		for entry != nil {
			fmt.Printf("[Key: %s, Value: %v] -> ", entry.Key, entry.Value)
			entry = entry.Next
		}
		fmt.Println("nil")
	}
}

func main() {
	ht := NewHashTable(10)

	ht.Put("Имя", "Данил")
	ht.Put("Язык", "Golang")
	ht.Put("Хобби", "Оптимизировать код")

	fmt.Println("Хэш-таблица")

	ht.Print()

	value, found := ht.Get("Язык")
	if found {
		fmt.Printf("\nНайден ключ 'Язык': %v\n", value)
	} else {
		fmt.Println("\nКлюч 'Язык' не найден")
	}

	ht.Remove("Имя")
	fmt.Println("\nХэш-таблица после удаления ключа 'Имя'")
	ht.Print()

}
