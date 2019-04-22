package gcp

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/jarcoal/httpmock"
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

func registerHTTPMockResponse(method, relativePath, response string) {
	fullPath := strings.Join([]string{baseURL, relativePath}, "")
	httpmock.RegisterResponder(method, fullPath, httpmock.NewStringResponder(http.StatusOK, response))
}

func withTestRoutes(t *testing.T, f func(t *testing.T)) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerHTTPMockResponse("GET", "/instance/hostname", `my-instance-name.c.my-application.internal`)
	registerHTTPMockResponse("GET", "/instance/name", `my-instance-name`)
	registerHTTPMockResponse("GET", "/instance/id", `7779877288063105047`)
	registerHTTPMockResponse("GET", "/instance/image", `projects/centos-cloud/global/images/centos-7-v20170719`)
	registerHTTPMockResponse("GET", "/instance/machine-type", `projects/354678626795/machineTypes/n1-standard-2`)
	registerHTTPMockResponse("GET", "/instance/cpu-platform", `Intel Broadwell`)
	registerHTTPMockResponse("GET", "/instance/zone", `projects/354678626795/zones/us-west1-a`)

	f(t)
}

func TestGCPProvider(t *testing.T) {
	gcpProvider := &Provider{}
	providerOptions := &provider.Options{HTTPTimeout: 5 * time.Second}

	withTestRoutes(t, func(t *testing.T) {
		result := gcpProvider.Check(providerOptions)
		assert.NotNil(t, result, "Result should not be nil.")

		gcpResult, ok := result.(Result)
		assert.True(t, ok, "Result should be a gcp Result.")

		assert.Equal(t, `my-instance-name.c.my-application.internal`, gcpResult.Hostname())
		assert.Equal(t, `my-instance-name`, gcpResult.InstanceName())
		assert.Equal(t, `7779877288063105047`, gcpResult.InstanceID())
		assert.Equal(t, `projects/centos-cloud/global/images/centos-7-v20170719`, gcpResult.Image())
		assert.Equal(t, `projects/354678626795/machineTypes/n1-standard-2`, gcpResult.MachineType())
		assert.Equal(t, `Intel Broadwell`, gcpResult.CPUPlatform())
		assert.Equal(t, `us-west1-a`, gcpResult.Zone())
		assert.Equal(t, `us-west1`, gcpResult.Region())
	})
}
