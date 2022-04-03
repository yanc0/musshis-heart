package musshi

import (
	"math/rand"
	"time"
)

type Musshi struct {
	Heart               *Heart
	IdealBeatsPerMinute int
	LifeTimeExpectancy  time.Duration
	lastAltered         time.Time
	BornAt              time.Time
	DeadAt              time.Time
	random              int
}

func NewMusshi() *Musshi {
	rand.Seed(time.Now().UnixNano())
	random := rand.Int()
	return &Musshi{
		Heart:               NewHeart(),
		IdealBeatsPerMinute: Sleeping.idealBPM(),
		LifeTimeExpectancy:  time.Second * 130,
		lastAltered:         time.Now(),
		BornAt:              time.Now(),
		DeadAt:              time.Time{},
		random:              random % 10,
	}
}

func (m *Musshi) GetIdealBPM() int {
	return m.Activity().idealBPM() + m.random
}

func (m *Musshi) Age() time.Duration {
	return time.Since(m.BornAt)
}

func (m *Musshi) GetCondition() Condition {
	difference := float64(m.Heart.BeatsPerMinute()) / float64(m.Activity().idealBPM())

	if difference > 2 {
		return VeryHighBPM
	}
	if difference > 1.2 {
		return TooHighBPM
	}
	if difference < 0.1 {
		return VeryLowBPM
	}
	if difference < 0.8 {
		return TooLowBPM
	}

	return idealBPM
}

func (m *Musshi) AlterLifeTimeExpectancy() {
	if time.Since(m.lastAltered) < time.Second {
		return
	}
	if !m.GetCondition().IsGood() {
		m.lastAltered = time.Now()

		if m.GetCondition() == VeryLowBPM || m.GetCondition() == VeryHighBPM {
			m.LifeTimeExpectancy -= time.Second * 3
			return
		}
		m.LifeTimeExpectancy -= time.Second * 1

	}
}

func (m *Musshi) Activity() Activity {
	switch {
	case m.Age() < time.Second*20:
		return Sleeping
	case m.Age() < time.Second*45:
		return Playing
	case m.Age() < time.Second*60:
		return Sleeping
	case m.Age() < time.Second*70:
		return Playing
	case m.Age() < time.Second*80:
		return Sleeping
	case m.Age() < time.Second*100:
		return Loving
	case m.Age() < time.Second*110:
		return Sleeping
	}
	return Dying
}

func (m *Musshi) Alive() bool {
	return m.Age() < m.LifeTimeExpectancy
}

func (m *Musshi) SetDeathTime(t time.Time) {
	m.DeadAt = t
}
