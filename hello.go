package main

import (
	"github.com/cvtienhoven/tray_plugin"
	"github.com/cvtienhoven/tray_sensu_plugin"
)

func main() {
	config := tray_plugin.Config{{"url", "http://sensu.prd.intra.tkppensioen.nl:4567/events"}, {"tag", "type=system"}, {"subscription", "owner=pdp"}}
	tray_sensu_plugin.GetStatus(config)
}
