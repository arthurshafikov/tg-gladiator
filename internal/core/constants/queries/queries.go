package queries

import (
	"fmt"
)

type Query string

const (
	OpenMenu Query = "OpenMenu"

	SpecialDelimeterInQueryCallback = "#_#"
)

func New(query Query, additionalFields ...string) Query {
	for _, field := range additionalFields {
		query += Query(fmt.Sprintf("%s%s", SpecialDelimeterInQueryCallback, field))
	}

	return query
}

func (i Query) With(additionalFields ...string) Query {
	return New(i, additionalFields...)
}

func (i Query) WithID(id int64) Query {
	return i + Query(fmt.Sprintf("%s%v", SpecialDelimeterInQueryCallback, id))
}
