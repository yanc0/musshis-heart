package musshi_test

import (
	"testing"
	"time"

	"github.com/yanc0/musshis-heart/musshi"
)

func TestHeartbeatsPerMinute(t *testing.T) {
	heart := musshi.NewHeart()

	start := time.Time{}

	if bpm := heart.BeatsPerMinute(); bpm != 0 {
		t.Fatalf("unexpected beats per minute, wants %d, got %d\n", 0, bpm)
	}

	for i := 0; i < 10; i++ {
		heart.Beat(start.Add(time.Duration(i) * time.Second))
	}

	if bpm := heart.BeatsPerMinute(); bpm != 60 {
		t.Fatalf("unexpected beats per minute, wants %d, got %d\n", 60, bpm)
	}
}
