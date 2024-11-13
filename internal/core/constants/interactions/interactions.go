package interactions

import (
	"fmt"
)

type Interaction string

const (
	HeroDelete Interaction = "HeroDelete"

	SpecialDelimeterForInteractions = "#_#"
)

func (i Interaction) WithID(id int64) Interaction {
	return i + Interaction(fmt.Sprintf("%s%v", SpecialDelimeterForInteractions, id))
}
