package main

import (
	"github.com/dahendel/docker-machine-driver-cloudstack"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(cloudstack.NewDriver("", ""))
}
