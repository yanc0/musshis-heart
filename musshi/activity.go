package musshi

type Activity string

const (
	Sleeping    Activity = "sleeping"
	Playing     Activity = "playing"
	Reproducing Activity = "reproducing"
	Dying       Activity = "dying"
	Dead        Activity = "dead"
)

func (a Activity) idealBPM() int {
	switch a {
	case Sleeping:
		return 60
	case Playing:
		return 135
	case Reproducing:
		return 400
	case Dying:
		return 20
	}
	return 0
}
