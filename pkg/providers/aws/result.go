package aws

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
)

var azRegexp = regexp.MustCompile(`\d`)

// Result implements provider.Result
type Result struct {
	responses map[string]string
}

// Provider returns the Provider that made the Result.
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

// AmiID returns the AMI ID returned from the metadata service.
func (r Result) AmiID() string {
	return r.responses["AWS_AMI_ID"]
}

// AmiLaunchIndex returns the launch index of the AMI.
func (r Result) AmiLaunchIndex() (int64, error) {
	strIndex := r.responses["AWS_AMI_LAUNCH_INDEX"]
	return strconv.ParseInt(strIndex, 10, 64)
}

// Hostname returns the instance's hostname.
func (r Result) Hostname() string {
	return r.responses["AWS_HOSTNAME"]
}

// InstanceID returns the instance ID.
func (r Result) InstanceID() string {
	return r.responses["AWS_INSTANCE_ID"]
}

// InstanceType returns the instance type.
func (r Result) InstanceType() string {
	return r.responses["AWS_INSTANCE_TYPE"]
}

// LocalIPv4 returns a net.IP representing the local IPv4 of the instance.
func (r Result) LocalIPv4() net.IP {
	return net.ParseIP(r.responses["AWS_LOCAL_IPV4"])
}

// MAC returns a net.HardwareAddr and error from parsing the MAC returned
// from the metadata service.
func (r Result) MAC() (net.HardwareAddr, error) {
	return net.ParseMAC(r.responses["AWS_MAC"])
}

// AvailabilityZone returns the availability zone of the instance.
func (r Result) AvailabilityZone() string {
	return r.responses["AWS_AVAILABILITY_ZONE"]
}

// Region returns the region that the instance resides in.
func (r Result) Region() string {
	return r.responses["AWS_REGION"]
}

// Domain returns the AWS domain of the region.
func (r Result) Domain() string {
	return r.responses["AWS_DOMAIN"]
}
