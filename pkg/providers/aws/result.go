package aws

import (
	"fmt"
	"regexp"
	"strings"
)

var azRegexp = regexp.MustCompile(`\d`)

// Result implements provider.Result
type Result struct {
	responses map[string]string
}

// Name returns the Cloud result name
func (r Result) Name() string {
	return "AWS"
}

// ToEval returns a string which should be able to be eval'd in a shell
func (r Result) ToEval() string {
	r.responses["CF_CLOUD"] = r.Name()
	exports := make([]string, 0)

	for k, v := range r.responses {
		exports = append(exports, fmt.Sprintf("export %s=%s", k, v))
	}
	return strings.Join(exports, "\n")
}

func (r Result) String() string {
	r.responses["CF_CLOUD"] = r.Name()
	items := make([]string, 0)
	for k, v := range r.responses {
		items = append(items, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(items, "\n")
}
