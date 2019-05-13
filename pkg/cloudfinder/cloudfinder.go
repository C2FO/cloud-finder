package cloudfinder

import (
	"sync"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/c2fo/cloud-finder/pkg/logging"
)

// Options configure cloudfinder and change how it behaves.
type Options struct {
	Timeout     time.Duration
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
func (cf *CloudFinder) Discover() provider.Result {
	ch := make(chan provider.Result)
	wg := sync.WaitGroup{}

	// Lock the registry so we don't get an unexpected change in the number of
	// providers between when we count them and when we iterate over them
	registryMu.Lock()
	defer registryMu.Unlock()

	// Add the number of providers to a wait group. Start a goroutine that waits
	// until all providers are done. Then send a nil on the channel to immediately
	// force a return from this function so we don't have to wait for the
	// timeout to occur.
	wg.Add(len(registry))
	go func() {
		wg.Wait()
		ch <- nil
		logging.Printf("No provider returned a successfull result.")
	}()

	opts := &provider.Options{
		HTTPTimeout: cf.opts.HTTPTimeout,
	}
	for _, pro := range registry {
		go func(p provider.Provider) {
			r := p.Check(opts)
			if r != nil {
				ch <- r
			}
			wg.Done()
		}(pro)
	}

	select {
	case r := <-ch:
		return r
	case <-time.After(cf.opts.Timeout):
		logging.Printf("Was not able to discover cloud provider before timeout of %0.0f seconds.", cf.opts.Timeout.Seconds())
		return nil
	}
}
