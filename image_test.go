package cast

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageId(t *testing.T) {
	b := bytes.NewBufferString("test data")
	s, err := ImageId(b)

	assert.Nil(t, err)
	assert.Equal(t, "578fbe87", s)

}
