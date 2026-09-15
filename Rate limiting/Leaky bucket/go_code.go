package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	mu        sync.Mutex
	capacity  int
	leakRate  float64
	queue     []time.Time
	lastLeak  time.Time
}

func NewLeakyBucket(capacity int, leakRate float64) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		lastLeak: time.Now(),
	}
}

func (b *LeakyBucket) leak() {
	now := time.Now()
	elapsed := now.Sub(b.lastLeak).Seconds()
	leaked := int(elapsed * b.leakRate)
	if leaked > 0 {
		if leaked > len(b.queue) {
			leaked = len(b.queue)
		}
		b.queue = b.queue[leaked:]
		b.lastLeak = now
	}
}

func (b *LeakyBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.leak()
	if len(b.queue) < b.capacity {
		b.queue = append(b.queue, time.Now())
		return true
	}
	return false
}

func main() {
	bucket := NewLeakyBucket(3, 1)
	for i := 0; i < 6; i++ {
		fmt.Println(i, bucket.Allow())
		time.Sleep(300 * time.Millisecond)
	}
}