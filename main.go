package main

import "C"
import (
	"fmt"

	"github.com/shockdev04/altv-go-pkg/alt"
)

func main() {}

//export OnStart
func OnStart() {
	alt.LogInfo("Resource Started")
	alt.On.PlayerConnect(func(p *alt.Player) {
		alt.LogInfo(fmt.Sprintf("Player %s connected", p.Name()))
	})
}

//export OnStop
func OnStop() {
	alt.LogInfo("Resource Stopped")
}
