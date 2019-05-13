package provider

import (
	"fmt"
	"time"
)

// Options that get passed to the Check method of each Provider
type Options struct {
	HTTPTimeout time.Duration
}

// Result from calling check on that particular Provider
type Result interface {
	fmt.Stringer
	Provider() Provider
	ToEval() string
}

// Provider interface that providers implement.
type Provider interface {
	Check(*Options) Result
	Name() string
}
