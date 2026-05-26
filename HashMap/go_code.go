package main

import "fmt"

const bucketSize = 1000

type entry struct {
	key, val int
}

type MyHashMap struct {
	buckets [bucketSize][]entry
}

func (m *MyHashMap) hash(key int) int {
	return key % bucketSize
}

func (m *MyHashMap) Put(key, val int) {
	h := m.hash(key)
	for i, e := range m.buckets[h] {
		if e.key == key {
			m.buckets[h][i].val = val // update existing
			return
		}
	}
	m.buckets[h] = append(m.buckets[h], entry{key, val})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for _, e := range m.buckets[h] {
		if e.key == key {
			return e.val
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	b := m.buckets[h]
	for i, e := range b {
		if e.key == key {
			m.buckets[h] = append(b[:i], b[i+1:]...)
			return
		}
	}
}

func main() {
	hm := &MyHashMap{}

	hm.Put(1, 10)
	hm.Put(2, 20)
	hm.Put(1001, 99) // collides with 1 (1001 % 1000 == 1)

	fmt.Println(hm.Get(1))    // 10
	fmt.Println(hm.Get(2))    // 20
	fmt.Println(hm.Get(3))    // -1 (not found)
	fmt.Println(hm.Get(1001)) // 99

	hm.Put(1, 42) // update key 1
	fmt.Println(hm.Get(1)) // 42

	hm.Remove(2)
	fmt.Println(hm.Get(2)) // -1 (removed)

	// Show collision bucket: keys 1 and 1001 share bucket[1]
	fmt.Println("bucket[1]:", hm.buckets[1]) // [{1 42} {1001 99}]
}
