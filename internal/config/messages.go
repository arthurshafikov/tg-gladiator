package config

type MessagesBag struct {
	RU Messages
	EN Messages
}

type Messages struct {
	StartSuccess string
	Help         string

	OpenedMenu     string
	DefaultBackBtn string

	MenuItemMyHeroes string
	MyHeroesList     string

	HeroName string

	Classes map[string]string
	Errors  map[string]string
}
