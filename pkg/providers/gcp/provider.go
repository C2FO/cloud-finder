package gcp

import (
	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
)

// Provider is the GCP cloudfinder provider
type Provider struct{}

// Check checks if we are in this cloud providers cloud. If so, we return the
// result. Else, nil.
func (p *Provider) Check(opts *provider.Options) provider.Result {
	return nil
}

func init() {
	cloudfinder.Register("gcp", &Provider{})
}
