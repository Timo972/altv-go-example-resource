package main

import "C"
import (
	"fmt"
	"strconv"

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
		p.Spawn(alt.Position{X: 1070.206, Y: -711.958, Z: 58.483}, 0)
		p.SetArmour(200)
		p.SetHealth(120)
		p.SetClothes(11, 7, 0, 2)
		p.GiveWeapon(3756226112, 999, false)
		p.GiveWeapon(2228681469, 999, false)
		p.AddWeaponComponent(2228681469, 0x837445AA)
		p.SetWeather(0)
	})

	alt.On.ConsoleCommand(func(command string, args []string) {
		if command == "veh" {
			model, _ := strconv.ParseUint(args[0], 2, 32)
			veh := alt.CreateVehicle(uint32(model), alt.Position{X: 1070.206, Y: -711.958, Z: 58.483},
				alt.Rotation{X: 1070.206, Y: -711.958, Z: 58.483})
			veh.SetNumberplateText("GO<3")
		}
	})

	//alt.On.ConsoleCommand(func(command string, args []string) {
	//	alt.LogColored(fmt.Sprintf("~y~%v ~r~Args: %v", command, args))
	//})
}

//export OnStop
func OnStop() {
	alt.LogInfo("Resource Stopped")
}
