package main

import (
	"log"
	"os"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
	_ "github.com/c2fo/cloud-finder/pkg/providers/aws"
	_ "github.com/c2fo/cloud-finder/pkg/providers/gcp"
)

var debug = true

func main() {

	if debug {
		log.Printf("Registered the following providers: %v", cloudfinder.Providers())
	}

	cf := cloudfinder.New(
		&cloudfinder.Options{
			HTTPTimeout: 5 * time.Second,
		},
	)
	result := cf.Discover()
	if result == nil {
		os.Exit(1)
	}
}
