package lang

import (
	"fmt"
)

type MapTranslator struct {
	messages map[string]string
}

var _ Translator = &MapTranslator{}

func NewMapTranslator() *MapTranslator {
	return &MapTranslator{
		messages,
	}
}

func (m MapTranslator) Translate(key string, replace ...string) string {
	value := messages[key]

	return fmt.Sprintf(value, replace)
}

