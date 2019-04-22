package gcp

import (
	"net/http"
	"strings"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/c2fo/cloud-finder/pkg/contrib"
	"github.com/c2fo/cloud-finder/pkg/logging"
)

const (
	baseURL = "http://metadata.google.internal/computeMetadata/v1"
	name    = "gcp"
)

// Provider is the GCP cloudfinder provider
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
	client.SetHeader("Metadata-Flavor", "Google")

	requests := map[string]string{
		"GCP_HOSTNAME":      "/instance/hostname",
		"GCP_INSTANCE_NAME": "/instance/name",
		"GCP_INSTANCE_ID":   "/instance/id",
		"GCP_IMAGE":         "/instance/image",
		"GCP_MACHINE_TYPE":  "/instance/machine-type",
		"GCP_CPU_PLATFORM":  "/instance/cpu-platform",
		"GCP_ZONE":          "/instance/zone",
	}

	responses, err := client.GetAll(requests)
	if err != nil {
		logging.Printf("Error in provider %s: %s", name, err)
		return nil
	}

	responses["GCP_ZONE"] = zone(responses["GCP_ZONE"])
	responses["GCP_REGION"] = region(responses["GCP_ZONE"])

	return Result{
		responses: responses,
	}
}

func zone(z string) string {
	pieces := strings.Split(z, "/")
	return pieces[len(pieces)-1]
}

func region(zone string) string {
	i := strings.LastIndex(zone, "-")
	if i < 0 {
		return zone
	}
	return zone[:i]
}
