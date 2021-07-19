package lang

type Translator interface {
	Translate(key string, replace ... string) string
}
