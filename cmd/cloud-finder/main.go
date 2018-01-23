package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
	"github.com/c2fo/cloud-finder/pkg/logging"
	_ "github.com/c2fo/cloud-finder/pkg/providers/aws"
	_ "github.com/c2fo/cloud-finder/pkg/providers/gcp"
)

var isDebug bool

func init() {
	flag.BoolVar(&isDebug, "debug", false, "Enable debug output")
}

func main() {
	flag.Parse()

	if isDebug {
		logging.EnableDebug()
	}

	logging.Printf("Registered the following providers: %v", cloudfinder.Providers())

	cf := cloudfinder.New(
		&cloudfinder.Options{
			Timeout:     30 * time.Second,
			HTTPTimeout: 5 * time.Second,
		},
	)

	result := cf.Discover()
	if result == nil {
		logging.Fatalf("Unable to determine which cloud we are in")
	}

	fmt.Println(result.ToEval())
}
