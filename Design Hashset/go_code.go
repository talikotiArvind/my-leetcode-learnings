package main
import "fmt"

const bucketSize = 500

type MyHashSet struct {
    buckets [bucketSize][] int
}

func (s *MyHashSet) hash (key int) int {
    return key % bucketSize
}

func (s *MyHashSet) Add (key int) {
    h := s.hash(key)
    for _, k := range s.buckets[h] {
        if k == key {
            return
        }
    }
    s.buckets[h] = append(s.buckets[h], key)
}

func (s * MyHashSet) Remove (key int) {
    h := s.hash(key)
    b := s.buckets[h]
    
    for i, k := range b {
        if k == key {
            s.buckets[h] = append(b[:i], b[i+1:]...)
        }
    }
}

func (s *MyHashSet) Contains (k int) bool {
    h := s.hash(k)
    for _, key := range s.buckets[h] {
        if k == key {
            return true
        }
    }
    return false
}

func main() {
	hs := &MyHashSet{}

	hs.Add(1)
	hs.Add(2)
	hs.Add(1001) // collides with 1 (1001 % 1000 == 1)

	fmt.Println(hs.Contains(1))    // true
	fmt.Println(hs.Contains(2))    // true
	fmt.Println(hs.Contains(3))    // false
	fmt.Println(hs.Contains(1001)) // true
}
