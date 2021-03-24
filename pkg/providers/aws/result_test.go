package aws

import (
	"fmt"
	"testing"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/stretchr/testify/assert"
)

var propertiesMap = map[string]string{
	"CF_CLOUD":              "AWS",
	"AWS_AMI_ID":            "ami-abcd1234",
	"AWS_AMI_LAUNCH_INDEX":  "0",
	"AWS_AVAILABILITY_ZONE": "us-west-2c",
	"AWS_HOSTNAME":          "ip-10-0-1-181",
	"AWS_INSTANCE_ID":       "i-abcd1234",
	"AWS_INSTANCE_TYPE":     "t2.medium",
	"AWS_LOCAL_IPV4":        "10.0.1.181",
	"AWS_MAC":               "0a:2e:31:ec:fa:45",
	"AWS_REGION":            "us-west-2",
}

func TestResultName(t *testing.T) {
	r := Result{}
	assert.Equal(t, "aws", r.Provider().Name())
}

func TestAWSResultImplementsProviderResult(t *testing.T) {
	assert.Implements(t, (*provider.Result)(nil), new(Result))
}

func TestToEval(t *testing.T) {
	var result provider.Result

	withTestRoutes(t, func(t *testing.T) {
		result = new(Provider).Check(
			&provider.Options{HTTPTimeout: 5 * time.Second},
		)
		assert.NotNil(t, result, "Result should not be nil.")
	})

	evalOutput := result.ToEval()

	for k, v := range propertiesMap {
		assert.Contains(t, evalOutput, fmt.Sprintf("export %s=%v", k, v))
	}
}

func TestString(t *testing.T) {
	var result provider.Result

	withTestRoutes(t, func(t *testing.T) {
		result = new(Provider).Check(
			&provider.Options{HTTPTimeout: 5 * time.Second},
		)
		assert.NotNil(t, result, "Result should not be nil.")
	})

	output := result.String()

	for k, v := range propertiesMap {
		assert.Contains(t, output, fmt.Sprintf("%s=%v", k, v))
	}
}
