package main

import (
	"fmt"
	"os"

	manifestdata "github.com/Bloem-Studios/bloem-community-drondeseries-aiometadata"
	"github.com/Bloem-Studios/bloem-community-drondeseries-aiometadata/internal/plugin"

	"github.com/Bloem-Studios/bloem-plugin-sdk/pkg/pluginsdk/manifest"
	"github.com/Bloem-Studios/bloem-plugin-sdk/pkg/pluginsdk/runtime"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{Name: "github.com/Bloem-Studios/bloem-community-drondeseries-aiometadata", Level: hclog.Info, Output: os.Stderr, JSONFormat: true})
	p := plugin.New(logger, nil)
	m, err := manifest.LoadWithChecksum(manifestdata.JSON, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	p.SetManifest(m)
	runtime.Serve(runtime.ServeConfig{Logger: logger, Servers: runtime.CapabilityServers{Runtime: p, MetadataProvider: p}})
}
