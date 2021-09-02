package main

import "C"
import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/shockdev04/altv-go-pkg/alt"
)

func main() {}

func randomNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomSpawn() alt.Position {
	return Spawns[uint(randomNumber(0, (len(Spawns)-1)))]
}

func getWeaponModel(name string) (uint32, error) {
	valid := false
	for _, wName := range WeaponModels {
		if wName == name {
			valid = true
		}
	}
	if valid {
		return alt.Hash(name), nil
	} else {
		return 0, errors.New("weapon model invalid")
	}
}

func maxVehicle(vehicle *alt.Vehicle) error {
	if vehicle.ModKitsCount() < 1 {
		return errors.New("vehicle has no modkits")
	}

	if vehicle.ModKit() == 0 {
		vehicle.SetModKit(1)
	}

	var modTypes uint8 = 48
	var i uint8 = 0

	for ; i < modTypes; i++ {
		modMax := vehicle.ModsCount(i)
		if modMax > 0 {
			vehicle.SetMod(i, modMax)
		}
	}

	return nil
}

//export OnStart
func OnStart() {
	alt.LogColored("~g~Resource Started")

	list := []interface{}{0, 1, 2, 3, "test", true, false}

	listMValue := alt.CreateMValue(list)

	alt.LogInfo("created list mvalue")

	valInterface := listMValue.GetValue()

	valArray, ok := valInterface.([]interface{})

	if ok {
		for _, val := range valArray {
			println("List val:", val)
		}
	}

	//alt.EmitServer("foo", "bar", 1, 2, false, nil)
	//alt.CreateColShapeCircle(0, 0, 0, 10)

	/*alt.SetTimeout(func() {
		alt.LogWarning("timeout after 1 sec")
	}, 1000)

	alt.SetTimeout(func() {
		alt.LogWarning("Timeout after 3 sec")

		interval := alt.SetInterval(func() {
			alt.LogWarning("Interval every 3 sec")
		}, 3000)

		alt.SetTimeout(func() {
			alt.ClearTimer(interval)
			alt.LogWarning("3 sec Interval clear after 15 sec")

			alt.SetTimeout(func() {
				alt.LogWarning("Timeout after a clear, after 3 sec")

				alt.NextTick(func() {
					alt.LogWarning("Next tick")
				})

				et := alt.EveryTick(func() {

				})

				alt.SetTimeout(func() {
					alt.ClearTimer(et)
					alt.LogWarning("clear everytick after 5 sec")
				}, 5000)

			}, 3000)

		}, 15000)

	}, 3000)

	return*/

	alt.On.AllServerEvents(func(eventName string, args ...interface{}) {
		alt.LogInfo(eventName)
		for _, arg := range args {
			println(arg)
		}

	})

	alt.On.AllClientEvents(func(player *alt.Player, eventName string, args ...interface{}) {
		alt.LogInfo(fmt.Sprintf("received event from player %v. name: %v, args: %v", player.Name(), eventName, args))
		alt.EmitClient(player, "test", "Are you receiving at least this?")
		player.Emit("test", "maybe this?")
		player.Emit("test", "Are you receiving this?", 1, 2, 1.2, true, false)
		alt.EmitAllClients("test", "fucking receive this message")

		alt.LogInfo("preparing for EmitClients")

		players := make([]*alt.Player, 0)
		players = append(players, player)
		alt.EmitClients(players, "test", "multiple clients emit test")
	})
	alt.On.EntityEnterColShape(func(c *alt.ColShape, entity interface{}) {
		if entity == nil {
			alt.LogError("No entity entered ColShape")
			return
		}
		v, err := entity.(*alt.Vehicle)
		if err {
			alt.LogInfo("Vehicle entered ColShape")
			driver := v.Driver()
			if driver.Valid() {
				v.SetNumberplateText(driver.Name())
			}
			return
		}
		p, err := entity.(*alt.Player)
		if err {
			alt.LogInfo("Player entered ColShape")
			alt.LogInfo(p.Name())
			//veh, err := alt.CreateVehicle(alt.Hash("krieger"), c.Position(), p.Rotation())
			//if err != nil {
			//	alt.LogError(err.Error())
			//	return
			//}
			//veh.SetNumberplateText(p.Name())
			return
		}

	})
	alt.On.PlayerConnect(func(p *alt.Player) {
		alt.LogInfo(fmt.Sprintf("Player %s %v connected with ip %s", p.Name(), p.ID(), p.IP()))

		if strings.Contains(strings.ToLower(p.Name()), "admin") {
			p.Kick("Please remove the word admin from your name!")
			return
		}

		p.SetModel(alt.Hash(SpawnModels[uint(randomNumber(0, len(SpawnModels)-1))]))
		//p.SetMetaData("vehicles", make([]*alt.Vehicle, 0))

		spawn := getRandomSpawn()

		p.Spawn(spawn, 0)

		//TODO fix timers and replace it with timeout 1000
		go func() {
			if p.Valid() {
				playerCount := len(alt.GetPlayers())
				//	//TODO add resource imports to module to use js chat
				alt.LogInfo(fmt.Sprintf("%v has joined the Server.. (%v players online)", p.Name(), playerCount))
				alt.LogInfo("Press T and type /help to see all available commands..")
			}
			//TODO clear timeout to reduce timer count
		}()

		// mp_m_freemode_01 -> alt.Hash function missing
		// p.SetModel(1885233650)
		// p.Spawn(alt.Position{X: 1070.206, Y: -711.958, Z: 58.483}, 0)
		// p.SetArmour(200)
		// p.SetHealth(120)
		// p.SetClothes(11, 7, 0, 2)
		// p.GiveWeapon(3756226112, 999, false)
		// p.GiveWeapon(2228681469, 999, false)
		// p.AddWeaponComponent(2228681469, 0x837445AA)
		// p.SetWeather(0)
	})

	alt.On.PlayerDeath(func(p *alt.Player, killer interface{}, weapon uint32) {
		spawn := getRandomSpawn()
		//TODO fix timers and replace it with timeout 3000
		//go func() {
		if p.Valid() {
			p.Spawn(spawn, 0)
			p.ClearBloodDamage()
		}
		//TODO clear timer
		//}()

		if killer != nil {
			if killer.(*alt.BaseObject).Type == alt.PlayerObject {
				alt.LogInfo(fmt.Sprintf("~r~%v ~s~killed ~b~%v", killer.(*alt.Player).Name(), p.Name()))
			}
		} else {
			alt.LogInfo(fmt.Sprintf("~s~Suicide ~b~%v", p.Name()))
		}
	})

	alt.On.PlayerDisconnect(func(p *alt.Player, reason string) {
		playerCount := len(alt.GetPlayers())
		alt.LogColored(fmt.Sprintf("%v has left the Server.. (%v players online)", p.Name(), playerCount))
		vehicles := p.GetMetaData("vehicles")
		vehs, ok := vehicles.([]*alt.Vehicle)
		if ok {
			p.DeleteMetaData("vehicles")
		}
		for _, vehicle := range vehs {
			vehicle.Destroy()
		}
	})

	alt.On.ConsoleCommand(func(command string, args []string) {
		if command == "veh" && len(args) == 2 {
			players := alt.GetPlayersByName(args[0])
			model := args[1]

			if players == nil || len(players) < 1 {
				alt.LogError(fmt.Sprintf("No player with name %v found", args[0]))
				return
			}

			player := players[0]

			rot := player.Rotation()
			pos := GetPositionInFront(player.Position(), rot, 3)

			veh, err := alt.CreateVehicle(alt.Hash(model), pos, rot)

			if err != nil {
				alt.LogError(fmt.Sprintf("Could not create vehicle, invalid model: %v", args[1]))
				return
			}

			veh.SetNumberplateText("GO<3")

			playerVehiclesMeta := player.GetMetaData("vehicles")

			if playerVehiclesMeta == nil {
				return
			}

			playerVehicles := playerVehiclesMeta.([]*alt.Vehicle)

			if len(playerVehicles) > 2 {
				lastVeh := playerVehicles[0]
				lastVeh.Destroy()
				playerVehicles = playerVehicles[1:]
			}

			playerVehicles = append(playerVehicles, veh)

			player.SetMetaData("vehicles", playerVehicles)
		}

		if command == "vehmax" && len(args) == 1 {
			players := alt.GetPlayersByName(args[0])
			if players == nil || len(players) < 1 {
				alt.LogError(fmt.Sprintf("No player with name %v found", args[0]))
				return
			}

			vehicle := players[0].Vehicle()

			if vehicle == nil || !vehicle.Valid() {
				alt.LogError("Player is not in a vehicle")
				return
			}

			err := maxVehicle(vehicle)

			if err != nil {
				alt.LogError(err.Error())
				return
			}
		}

		if command == "weapon" && len(args) == 2 {
			players := alt.GetPlayersByName(args[0])
			weaponName := args[1]

			if players == nil || len(players) < 1 {
				alt.LogError(fmt.Sprintf("No player with name %v found", args[0]))
				return
			}

			model, err := getWeaponModel(weaponName)

			if err != nil {
				alt.LogError(err.Error())
				return
			}

			player := players[0]

			player.GiveWeapon(model, 999, true)
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
