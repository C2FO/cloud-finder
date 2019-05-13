package gcp

import (
	"fmt"
	"testing"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/stretchr/testify/assert"
)

var propertiesMap = map[string]string{
	"GCP_HOSTNAME":     "my-instance-name.c.my-application.internal",
	"GCP_INSTANCE_ID":  "7779877288063105047",
	"GCP_IMAGE":        "projects/centos-cloud/global/images/centos-7-v20170719",
	"GCP_MACHINE_TYPE": "projects/354678626795/machineTypes/n1-standard-2",
	"GCP_CPU_PLATFORM": "Intel Broadwell",
	"GCP_ZONE":         "us-west1-a",
	"GCP_REGION":       "us-west1",
}

func TestResultName(t *testing.T) {
	r := Result{}
	assert.Equal(t, "gcp", r.Provider().Name())
}

func TestGCPResultImplementsProviderResult(t *testing.T) {
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
