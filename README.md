# Cloud Finder

![Build Status](https://github.com/c2fo/cloud-finder/actions/workflows/ci.yml/badge.svg?branch=master) [![codecov](https://codecov.io/gh/c2fo/cloud-finder/branch/master/graph/badge.svg)](https://codecov.io/gh/c2fo/cloud-finder)

This project should serve as a way to help processes discover which cloud,
region, endpoint, etc. they need for dynamic configuration.

<!-- TOC depthFrom:1 depthTo:6 withLinks:1 updateOnSave:1 orderedList:0 -->

- [Cloud Finder](#cloud-finder)
	- [Rationale](#rationale)
	- [Usage](#usage)
		- [Binary](#binary)
			- [Exit Codes](#exit-codes)
			- [Environment Variables](#environment-variables)
			- [AWS](#aws)
			- [GCP](#gcp)
			- [Azure](#azure)
		- [In Go Code](#in-go-code)

<!-- /TOC -->

## Rationale

When deploying docker containers to different clouds, you might often want to
configure your application based on specific rules. Things that might be
important to you may include:

* Cloud Provider
* Region
* Availability Zone

Ideally, this logic could reside in the application. Upon application boot,
the application could figure out all of these things for itself and configure
itself accordingly. However, we often must deal with what we have. In some
cases, getting that logic into the main application might not be feasible.

## Usage

### Binary

The most common use case for this right now is to call cloud-finder in a
subprocess and to eval its output like so:

```sh
eval $(cloud-finder -output=eval)
echo $CF_CLOUD
```


#### Exit Codes

0 - Able to determine a Cloud Provider
1 - Not able to determine a Cloud Provider

#### Environment Variables

#### AWS

```
CF_CLOUD=AWS
AWS_LOCAL_IPV4=1.1.1.1
AWS_INSTANCE_TYPE=t2.medium
AWS_AMI_LAUNCH_INDEX=0
AWS_REGION=us-west-2
AWS_MAC=0a:2e:31:ea:ab:35
AWS_AVAILABILITY_ZONE=us-west-2c
AWS_AMI_ID=ami-edfc2abd
AWS_HOSTNAME=ip-1-1-1-1
AWS_INSTANCE_ID=i-8af98240
AWS_DOMAIN=amazonaws.com
```

#### GCP

```
GCP_INSTANCE_NAME=my-instance-name
GCP_MACHINE_TYPE=projects/234678626784/machineTypes/n1-standard-4
GCP_CPU_PLATFORM=Intel Broadwell
GCP_ZONE=us-west1-b
GCP_REGION=us-west1
CF_CLOUD=GCP
GCP_HOSTNAME=my-instance-name.c.my-application.internal
GCP_IMAGE=projects/gke-node-images/global/images/gke-1-8-6-gke-0-cos-stable-63-10032-71-0-p-v180105-pre
GCP_INSTANCE_ID=2693006570498178293
```

#### Azure

```
AZURE_MAC=000D3A6C292C
AZURE_LOCATION=UAENorth
AZURE_ZONE=
AZURE_TAGS=creationSource:aks-aks-default-00000000-vmss;orchestrator:Kubernetes:1.18.14;poolName:default;resourceNameSuffix:00000000
AZURE_VM_SIZE=Standard_B8MS
AZURE_VM_ID=00000000-856d-4a02-bbcf-00000000
CF_CLOUD=AZURE
AZURE_PRIVATE_IPV4=10.10.1.10
```

### In Go Code

You can also import this in your go code and use it to determine which cloud you are in like so:

```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/c2fo/cloud-finder/pkg/providers/aws"
	"github.com/c2fo/cloud-finder/pkg/providers/gcp"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder"
)

func main() {
	log.Printf("Registered the following providers: %v", cloudfinder.Providers())

	cf := cloudfinder.New(
		&cloudfinder.Options{
			Timeout:     30 * time.Second,
			HTTPTimeout: 5 * time.Second,
		},
	)

	result := cf.Discover()
	if result == nil {
		log.Fatalf("Unable to determine which cloud we are in")
	}

	var datacenter string

	switch v := result.(type) {
	case gcp.Result:
		datacenter = v.Region()
	case aws.Result:
		datacenter = v.Region()
	default:
		log.Fatalf("No datacenter configuration has been defined for provider: %s\n", result.Name())
	}

	fmt.Printf("Determined we are in datacenter: %s\n", datacenter)
	fmt.Println(v.ToString())
}
```
