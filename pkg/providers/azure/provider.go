package azure

import (
	"encoding/json"
	"net/http"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/c2fo/cloud-finder/pkg/contrib"
	"github.com/c2fo/cloud-finder/pkg/logging"
)

// Constants for the Azure provider
const (
	name    = "azure"
	baseURL = "http://169.254.169.254"
)

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
	httpClient := &http.Client{Timeout: opts.HTTPTimeout}
	client := contrib.NewClient(httpClient)
	client.SetBaseURL(baseURL)
	client.SetHeader("Metadata", "true")

	resp, err := client.Get("/metadata/instance?api-version=2020-09-01")
	if err != nil {
		logging.Printf("Error in provider %s: %s", name, err)
		return nil
	}

	imr := Result{}
	err = json.Unmarshal([]byte(resp), &imr)
	if err != nil {
		logging.Printf("Error unmarshalling resp in provider %s: %s", name, err)
		return nil
	}

	return imr
}
