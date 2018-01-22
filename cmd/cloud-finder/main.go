package main

import (
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
	"github.com/c2fo/cloud-finder/pkg/logging"
	_ "github.com/c2fo/cloud-finder/pkg/providers/aws"
	_ "github.com/c2fo/cloud-finder/pkg/providers/gcp"
)

func main() {
	logging.EnableDebug()
	logging.Printf("Registered the following providers: %v", cloudfinder.Providers())

	cf := cloudfinder.New(
		&cloudfinder.Options{
			HTTPTimeout: 5 * time.Second,
		},
	)

	result := cf.Discover()
	if result == nil {
		logging.Fatalf("Unable to determine which cloud we are in")
	}
}
