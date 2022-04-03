package scenes

import (
	"fmt"

	"github.com/yanc0/musshis-heart/musshi"
)

func describeActivity(m *musshi.Musshi) string {
	switch m.Activity() {
	case musshi.Sleeping:
		return fmt.Sprintf("Your Musshi sleeps quietly.\nYou can rest at around %d BPM.", m.GetIdealBPM())
	case musshi.Playing:
		return fmt.Sprintf("Your Musshi plays with his friends and burns energy.\nYou gotta keep the beat around %d BPM.", m.GetIdealBPM())
	case musshi.Loving:
		return fmt.Sprintf("Your Musshi has found love.\nYou must beat wildly at around %d BPM.", m.GetIdealBPM())
	case musshi.Dying:
		return fmt.Sprintf("Your Musshi had a great life.\nYou let him go slowly at around %d BPM.", m.GetIdealBPM())
	default:
		return string(m.Activity())
	}
}
