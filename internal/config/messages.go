package config

type MessagesBag struct {
	RU Messages
	EN Messages
}

type Messages struct {
	Errors map[string]string
}
