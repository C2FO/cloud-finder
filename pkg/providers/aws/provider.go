package aws

import (
	"net/http"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/c2fo/cloud-finder/pkg/contrib"
	"github.com/c2fo/cloud-finder/pkg/logging"
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
	httpClient := &http.Client{Timeout: opts.HTTPTimeout}
	client := contrib.NewClient(httpClient)
	client.SetBaseURL("http://169.254.169.254")

	requests := map[string]string{
		"AWS_AMI_ID":            "/latest/meta-data/ami-id",
		"AWS_AMI_LAUNCH_INDEX":  "/latest/meta-data/ami-launch-index",
		"AWS_HOSTNAME":          "/latest/meta-data/hostname",
		"AWS_INSTANCE_ID":       "/latest/meta-data/instance-id",
		"AWS_INSTANCE_TYPE":     "/latest/meta-data/instance-type",
		"AWS_LOCAL_IPV4":        "/latest/meta-data/local-ipv4",
		"AWS_MAC":               "/latest/meta-data/mac",
		"AWS_AVAILABILITY_ZONE": "/latest/meta-data/placement/availability-zone",
	}

	responses, err := client.GetAll(requests)
	if err != nil {
		logging.Printf("Error in provider %s: %s", name, err)
		return nil
	}
	responses["AWS_REGION"] = region(responses["AWS_AVAILABILITY_ZONE"])

	return Result{
		responses: responses,
	}
}

func region(az string) string {
	loc := azRegexp.FindStringIndex(az)
	if loc == nil {
		return ""
	}
	return az[:loc[1]]
}
