package lang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessage_Success(t *testing.T) {
	translator := NewMapTranslator()

	message := translator.Translate("error.url.parameter", "id")

	assert.EqualValues(t, "Malformed parameter [id]", message)
}
