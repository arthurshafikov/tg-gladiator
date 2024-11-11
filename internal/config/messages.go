package config

type MessagesBag struct {
	RU Messages
	EN Messages
}

type Messages struct {
	StartSuccess string

	Errors map[string]string
}
