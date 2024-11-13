package interactions

import (
	"fmt"
)

type Interaction string

const (
	HeroCreationEnterName Interaction = "HeroCreationEnterName"

	HeroDelete Interaction = "HeroDelete"

	SpecialDelimeterForInteractions = "#_#"
)

func New(interaction Interaction, additionalFields ...string) Interaction {
	for _, field := range additionalFields {
		interaction += Interaction(fmt.Sprintf("%s%s", SpecialDelimeterForInteractions, field))
	}

	return interaction
}

func (i Interaction) With(additionalFields ...string) Interaction {
	return New(i, additionalFields...)
}

func (i Interaction) WithID(id int64) Interaction {
	return i + Interaction(fmt.Sprintf("%s%v", SpecialDelimeterForInteractions, id))
}
