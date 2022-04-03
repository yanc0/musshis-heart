package musshi

import (
	"sync"
	"time"
)

type Heart struct {
	mu    *sync.Mutex
	beats []time.Time
}

func NewHeart() *Heart {
	return &Heart{
		mu:    &sync.Mutex{},
		beats: make([]time.Time, 0),
	}
}

func (h *Heart) Beat(t time.Time) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.beats = append(h.beats, t)
}

func (h *Heart) JustBeaten() bool {
	return time.Since(h.beats[len(h.beats)]) < 100*time.Millisecond
}

func (h *Heart) Electrocardiogram() []float64 {
	var ecg = make([]float64, 0)
	end := time.Now()
	start := end.Add(-6 * time.Second)

	for i := 0; i < 30; i++ {
		end = start.Add(200 * time.Millisecond)
		toAdd := []float64{1, 1}
		if h.hasBeatenBetween(start, end) {
			toAdd = []float64{5, -1}
		}
		ecg = append(ecg, toAdd...)
		start = end
	}
	return ecg
}

func (h *Heart) hasBeatenBetween(start time.Time, end time.Time) bool {
	for _, beat := range h.beats {
		if start.Before(beat) && end.After(beat) {
			return true
		}
	}
	return false
}

func (h *Heart) BeatsPerMinute() int {
	h.mu.Lock()
	defer h.mu.Unlock()

	// keep only the last 6 seconds beats
	lastBeats := make([]time.Time, 0)
	for _, beat := range h.beats {
		if time.Since(beat) < time.Second*6 {
			lastBeats = append(lastBeats, beat)
		}
	}

	h.beats = lastBeats

	numBeats := len(h.beats)

	if numBeats <= 1 {
		return 0
	}

	if numBeats <= 2 {
		return int(time.Minute / h.beats[numBeats-1].Sub(h.beats[numBeats-2]))
	}

	if numBeats <= 3 {
		return (int(time.Minute/h.beats[numBeats-1].Sub(h.beats[numBeats-2])) +
			int(time.Minute/h.beats[numBeats-2].Sub(h.beats[numBeats-3]))) / 2
	}

	return (int(time.Minute/h.beats[numBeats-1].Sub(h.beats[numBeats-2])) +
		int(time.Minute/h.beats[numBeats-2].Sub(h.beats[numBeats-3])) +
		int(time.Minute/h.beats[numBeats-3].Sub(h.beats[numBeats-4]))) / 3

}
