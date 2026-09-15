package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowCounter struct {
	mu         sync.Mutex
	limit      int
	window     time.Duration
	timestamps []time.Time
}

func NewSlidingWindowCounter(limit int, window time.Duration) *SlidingWindowCounter {
	return &SlidingWindowCounter{limit: limit, window: window}
}

func (s *SlidingWindowCounter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	i := 0
	for i < len(s.timestamps) && now.Sub(s.timestamps[i]) > s.window {
		i++
	}
	s.timestamps = s.timestamps[i:]
	if len(s.timestamps) < s.limit {
		s.timestamps = append(s.timestamps, now)
		return true
	}
	return false
}

func main() {
	limiter := NewSlidingWindowCounter(3, time.Second)
	for i := 0; i < 6; i++ {
		fmt.Println(i, limiter.Allow())
		time.Sleep(300 * time.Millisecond)
	}
}