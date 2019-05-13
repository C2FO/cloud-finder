package gcp

import (
	"fmt"
	"strings"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
)

// Result implements provider.Result
type Result struct {
	responses map[string]string
}

// Provider returns the Provider that made the Result
func (r Result) Provider() provider.Provider {
	return &Provider{}
}

// ToEval returns a string which should be able to be eval'd in a shell
func (r Result) ToEval() string {
	r.responses["CF_CLOUD"] = strings.ToUpper(r.Provider().Name())
	exports := make([]string, 0)

	for k, v := range r.responses {
		exports = append(exports, fmt.Sprintf("export %s=%s", k, v))
	}
	return strings.Join(exports, "\n")
}

func (r Result) String() string {
	r.responses["CF_CLOUD"] = strings.ToUpper(r.Provider().Name())
	items := make([]string, 0)
	for k, v := range r.responses {
		items = append(items, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(items, "\n")
}

// Hostname returns the instance's hostname.
func (r Result) Hostname() string {
	return r.responses["GCP_HOSTNAME"]
}

// InstanceName returns the instance's Name.
func (r Result) InstanceName() string {
	return r.responses["GCP_INSTANCE_NAME"]
}

// InstanceID returns the instance's ID.
func (r Result) InstanceID() string {
	return r.responses["GCP_INSTANCE_ID"]
}

// Image returns the image used to create the instance.
func (r Result) Image() string {
	return r.responses["GCP_IMAGE"]
}

// MachineType returns the machine type of the instance.
func (r Result) MachineType() string {
	return r.responses["GCP_MACHINE_TYPE"]
}

// CPUPlatform returns the CPU Platform of the instance.
func (r Result) CPUPlatform() string {
	return r.responses["GCP_CPU_PLATFORM"]
}

// Zone returns the zone that the instance is in.
func (r Result) Zone() string {
	return r.responses["GCP_ZONE"]
}

// Region returns the region that the instance is in.
func (r Result) Region() string {
	return r.responses["GCP_REGION"]
}
