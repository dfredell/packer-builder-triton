package main

import (
	"github.com/jen20/packer-builder-triton/triton"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(triton.Builder))
	server.Serve()
}
