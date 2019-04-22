package gcp

import (
	"testing"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/stretchr/testify/assert"
)

func TestZone(t *testing.T) {
	assert.Equal(t, "us-west1-a", zone("projects/123456789/zones/us-west1-a"))
	assert.Equal(t, "us-west1-a", zone("us-west1-a"))
}

func TestRegion(t *testing.T) {
	assert.Equal(t, "us-west1", region("us-west1-a"))
}
func TestGCPProviderImplementsProvider(t *testing.T) {
	assert.Implements(t, (*provider.Provider)(nil), new(Provider))
}
