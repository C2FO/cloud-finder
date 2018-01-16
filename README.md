# Cloud Finder

This project should serve as a way to help processes discover which cloud,
region, endpoint, etc. they need for dynamic configuration. 


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
eval $(cloud-finder --output=eval)
echo $CLOUD_PROVIDER
```
