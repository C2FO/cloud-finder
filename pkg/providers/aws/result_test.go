package aws

import (
	"testing"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/stretchr/testify/assert"
)

func TestResultName(t *testing.T) {
	r := Result{}
	assert.Equal(t, "AWS", r.Name())
}

func TestAWSResultImplementsProviderResult(t *testing.T) {
	assert.Implements(t, (*provider.Result)(nil), new(Result))
}
