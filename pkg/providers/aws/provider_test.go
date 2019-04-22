package aws

import (
	"testing"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/stretchr/testify/assert"
)

func TestProviderName(t *testing.T) {
	p := Provider{}
	assert.Equal(t, "aws", p.Name())
}

func TestRegion(t *testing.T) {
	assert.Equal(t, "us-west-2", region("us-west-2a"))
	assert.Equal(t, "", region("us-west"))
	assert.Equal(t, "ap-south-1", region("ap-south-1a"))
	assert.Equal(t, "us-gov-west-1", region("us-gov-west-1"))
}

func TestAWSProviderImplementsProvider(t *testing.T) {
	assert.Implements(t, (*provider.Provider)(nil), new(Provider))
}
