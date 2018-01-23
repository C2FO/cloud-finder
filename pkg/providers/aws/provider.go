package aws

import (
	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
)

const name = "aws"

// Provider is the AWS cloudfinder provider
type Provider struct{}

func init() {
	cloudfinder.Register(name, &Provider{})
}

// Name returns the provider name as needed by provider.Provider
func (p *Provider) Name() string {
	return name
}

// Check checks if we are in this cloud providers cloud. If so, we return the
// result. Else, nil.
func (p *Provider) Check(opts *provider.Options) provider.Result {
	return nil
}
