# Docker Machine CloudStack Driver

Docker Machine CloudStack Driver is a driver for [Docker Machine](https://docs.docker.com/machine/).
It allows to create Docker hosts on [Apache CloudStack](https://cloudstack.apache.org/) and
[Accelerite CloudPlatform](http://cloudplatform.accelerite.com/).

## Added Functionality
- The `--cloudstack-userdata-file` now supports passing urls
- In the userdata file if `ssh_authorized_keys` then the driver grabs the first public key and expects that `ssh_keys`
contains a private key so that docker-machine can ssh into the VM after creation (Deprecated)
- It also checks if the key in the userdata file exists in cloudstack based on the public key fingerprint, and if it does writes it to the new
machines StoragePath (/home/$USER/.docker/machine/machines/$MACHINE_NAME)


## Requirements

* [Docker Machine](https://docs.docker.com/machine/) 0.5.1 or later

## Installation

Download the binary from following link and put it within your PATH (ex. `/usr/local/bin`)

https://github.com/dahendel/docker-machine-driver-cloudstack/releases/latest

## Usage

```bash
docker-machine create -d cloudstack \
  --cloudstack-api-url CLOUDSTACK_API_URL \
  --cloudstack-api-key CLOUDSTACK_API_KEY \
  --cloudstack-secret-key CLOUDSTACK_SECRET_KEY \
  --cloudstack-template "Ubuntu Server 14.04" \
  --cloudstack-zone "zone01" \
  --cloudstack-service-offering "Small" \
  --cloudstack-expunge \
  docker-machine
```

## Acknowledgement

The driver is originally written by [@svanharmelen](https://github.com/svanharmelen), [@atsaki](https://github.com/atsaki) and [@andrestc](https://github.com/andrestc).

