package events

type Event string

const (
	HeroLevelUp Event = "HeroLevelUp"
	AnalyticEvent   Event = "AnalyticEvent"
)

func (e Event) ToString() string {
	return string(e)
}
