package musshi

type Condition string

const (
	VeryHighBPM Condition = "VeryHighBPM"
	TooHighBPM  Condition = "TooHighBPM"
	TooLowBPM   Condition = "TooLowBPM"
	VeryLowBPM  Condition = "VeryLowBPM"

	idealBPM Condition = "idealBPM"
	Unknown  Condition = "Unknown"
)

func (c Condition) IsGood() bool {
	if c == idealBPM {
		return true
	}
	return false
}
