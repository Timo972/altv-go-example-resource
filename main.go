package main

import "C"
import (
	"fmt"

	"github.com/shockdev04/altv-go-pkg/alt"
)

func main() {}

//export OnStart
func OnStart() {
	alt.LogColored("~g~Resource Started")
	alt.On.PlayerConnect(func(p *alt.Player) {
		alt.LogInfo(fmt.Sprintf("Player %s %v connected with ip %s", p.Name(), p.GetID(), p.GetIP()))
		// mp_m_freemode_01 -> alt.Hash function missing
		p.SetModel(1885233650)
		p.Spawn(alt.Position{X: 0, Y: 0, Z: 80}, 0)
		p.SetArmour(200)
		p.SetHealth(120)
		p.SetClothes(11, 7, 0, 2)
		p.GiveWeapon(3756226112, 999, true)
		p.SetWeather(0)
	})
}

//export OnStop
func OnStop() {
	alt.LogInfo("Resource Stopped")
}
