package azure

import (
	"fmt"
	"testing"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/stretchr/testify/assert"
)

func TestResultName(t *testing.T) {
	r := Result{}
	assert.Equal(t, "azure", r.Provider().Name())
}

func TestAzureResultImplementsProviderResult(t *testing.T) {
	assert.Implements(t, (*provider.Result)(nil), new(Result))
}

func TestToEval(t *testing.T) {
	var result provider.Result

	withTestRoutes(t, mockResp, func(t *testing.T) {
		result = new(Provider).Check(
			&provider.Options{HTTPTimeout: 5 * time.Second},
		)
		assert.NotNil(t, result, "Result should not be nil.")
	})

	evalOutput := result.ToEval()

	for k, v := range result.(Result).properties() {
		assert.Contains(t, evalOutput, fmt.Sprintf("export %s=%v", k, v))
	}
}

func TestString(t *testing.T) {
	var result provider.Result

	withTestRoutes(t, mockResp, func(t *testing.T) {
		result = new(Provider).Check(
			&provider.Options{HTTPTimeout: 5 * time.Second},
		)
		assert.NotNil(t, result, "Result should not be nil.")
	})

	output := result.String()

	for k, v := range result.(Result).properties() {
		assert.Contains(t, output, fmt.Sprintf("%s=%v", k, v))
	}
}
