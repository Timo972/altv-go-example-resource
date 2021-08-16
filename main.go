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
		alt.LogInfo(fmt.Sprintf("Player %s %v connected with ip %s", p.Name(), p.ID(), p.IP()))
		// mp_m_freemode_01 -> alt.Hash function missing
		p.SetModel(1885233650)
		p.Spawn(alt.Position{X: 0, Y: 0, Z: 80}, 0)
		p.SetArmour(200)
		p.SetHealth(120)
		p.SetClothes(11, 7, 0, 2)
		p.GiveWeapon(3756226112, 999, false)
		p.GiveWeapon(2228681469, 999, true)
		p.AddWeaponComponent(2228681469, 0x837445AA)
		p.SetWeather(0)
		pos := p.Position()
		checkpoint := alt.CreateCheckpoint(0, pos.X+15, pos.Y+10, pos.Z+30, 5, 10, alt.RGBA{R: 100, G: 100, B: 100, A: 100})
		println(checkpoint.IsPlayersOnly())
		alt.CreateVehicle(3630826055, p.Position(), p.Rotation())
	})

	//alt.On.ConsoleCommand(func(command string, args []string) {
	//	alt.LogColored(fmt.Sprintf("~y~%v ~r~Args: %v", command, args))
	//})
}

//export OnStop
func OnStop() {
	alt.LogInfo("Resource Stopped")
}
