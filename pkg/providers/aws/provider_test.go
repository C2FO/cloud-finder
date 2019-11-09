package aws

import (
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/jarcoal/httpmock"
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
	assert.Equal(t, "cn-north-1", region("cn-north-1a"))
}

func TestAWSProviderImplementsProvider(t *testing.T) {
	assert.Implements(t, (*provider.Provider)(nil), new(Provider))
}

func registerHTTPMockResponse(method, relativePath, response string) {
	fullPath := strings.Join([]string{baseURL, relativePath}, "")
	httpmock.RegisterResponder(method, fullPath, httpmock.NewStringResponder(http.StatusOK, response))
}

func withTestRoutes(t *testing.T, f func(t *testing.T)) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerHTTPMockResponse("GET", "/latest/meta-data/ami-id", `ami-abcd1234`)
	registerHTTPMockResponse("GET", "/latest/meta-data/ami-launch-index", `0`)
	registerHTTPMockResponse("GET", "/latest/meta-data/hostname", `ip-10-0-1-181`)
	registerHTTPMockResponse("GET", "/latest/meta-data/instance-id", `i-abcd1234`)
	registerHTTPMockResponse("GET", "/latest/meta-data/instance-type", `t2.medium`)
	registerHTTPMockResponse("GET", "/latest/meta-data/local-ipv4", `10.0.1.181`)
	registerHTTPMockResponse("GET", "/latest/meta-data/mac", `0a:2e:31:ec:fa:45`)
	registerHTTPMockResponse("GET", "/latest/meta-data/placement/availability-zone", `us-west-2c`)
	registerHTTPMockResponse("GET", "/latest/meta-data/services/domain", `amazonaws.com`)

	f(t)
}

func TestAWSProvider(t *testing.T) {
	awsProvider := &Provider{}
	providerOptions := &provider.Options{HTTPTimeout: 5 * time.Second}

	withTestRoutes(t, func(t *testing.T) {
		result := awsProvider.Check(providerOptions)
		assert.NotNil(t, result, "Result should not be nil.")

		awsResult, ok := result.(Result)
		assert.True(t, ok, "Result should be an AWS Result.")

		assert.Equal(t, `ami-abcd1234`, awsResult.AmiID())
		launchIndex, err := awsResult.AmiLaunchIndex()
		assert.NoError(t, err)

		assert.Equal(t, int64(0), launchIndex)
		assert.Equal(t, `ip-10-0-1-181`, awsResult.Hostname())
		assert.Equal(t, `i-abcd1234`, awsResult.InstanceID())
		assert.Equal(t, `t2.medium`, awsResult.InstanceType())
		assert.Equal(t, net.IPv4(0x0a, 0x00, 0x01, 0xb5), awsResult.LocalIPv4())

		mac, err := awsResult.MAC()
		assert.NoError(t, err)
		assert.Equal(t, `0a:2e:31:ec:fa:45`, mac.String())

		assert.Equal(t, "us-west-2c", awsResult.AvailabilityZone())
		assert.Equal(t, "us-west-2", awsResult.Region())
		assert.Equal(t, "amazonaws.com", awsResult.Domain())
	})
}
