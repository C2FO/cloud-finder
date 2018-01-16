package cloudfinder

import (
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
)

// Result is what we get when we check a cloud provider.
type Result interface {
	// print a string that can be eval'd in bash.
	ToEval() string
}

// Options configure cloudfinder and change how it behaves.
type Options struct {
	HTTPTimeout time.Duration
}

// CloudFinder is the thing that interacts with all of the providers for us and
// simply returns a result to us.
type CloudFinder struct {
	opts *Options
}

// New creates and returns a new CloudFinder
func New(opts *Options) *CloudFinder {
	return &CloudFinder{
		opts: opts,
	}
}

// Discover which cloud provider we are in an return a Result accordingly.
// Returns nil if we are not in a cloud provider, or the cloud provider could
// not be determined.
func (cf *CloudFinder) Discover() Result {
	ch := make(chan provider.Result)
	for _, pro := range registry {
		go func(p provider.Provider) {
			res := p.Check(&provider.Options{
				HTTPTimeout: cf.opts.HTTPTimeout,
			})
			if res != nil {
				ch <- res
			}
		}(pro)
	}
	return nil
}
