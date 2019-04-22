# Cloud Finder

[![Build Status](https://travis-ci.org/C2FO/cloud-finder.svg?branch=master)](https://travis-ci.org/C2FO/cloud-finder) [![codecov](https://codecov.io/gh/c2fo/cloud-finder/branch/master/graph/badge.svg)](https://codecov.io/gh/c2fo/cloud-finder)

This project should serve as a way to help processes discover which cloud,
region, endpoint, etc. they need for dynamic configuration. 

<!-- TOC depthFrom:1 depthTo:6 withLinks:1 updateOnSave:1 orderedList:0 -->

- [Cloud Finder](#cloud-finder)
- [Rationale](#rationale)
- [Usage](#usage)
	- [Exit Codes](#exit-codes)
	- [Environment Variables](#environment-variables)
		- [AWS](#aws)
		- [GCP](#gcp)

<!-- /TOC -->

# Rationale

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

# Usage

The most common use case for this right now is to call cloud-finder in a
subprocess and to eval its output like so:

```sh
eval $(cloud-finder -output=eval)
echo $CF_CLOUD
```


## Exit Codes

0 - Able to determine a Cloud Provider
1 - Not able to determine a Cloud Provider

## Environment Variables

### AWS

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
```

### GCP

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
