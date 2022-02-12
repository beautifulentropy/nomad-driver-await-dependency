package main

import (
	await "github.com/beautifulentropy/nomad-driver-await-dependency/await"
	log "github.com/hashicorp/go-hclog"

	"github.com/hashicorp/nomad/plugins"
)

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of a nomad driver plugin
func factory(log log.Logger) interface{} {
	return await.NewPlugin(log)
}
