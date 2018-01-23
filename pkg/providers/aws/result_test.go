package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultName(t *testing.T) {
	r := Result{}
	assert.Equal(t, "AWS", r.Name())
}
