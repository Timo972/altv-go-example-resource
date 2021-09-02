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

func getRandomSpawn() alt.Vector3 {
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

func SendNotificationToPlayer(p *alt.Player, message string, textColor uint, bgColor uint, blink bool) {
	alt.EmitClient(p, "freeroam:sendNotification", textColor, bgColor, message, blink)
}

func SendNotificationToAllPlayer(message string, textColor uint, bgColor uint, blink bool) {
	alt.EmitAllClients("freeroam:sendNotification", textColor, bgColor, message, blink)
}

//export OnStart
func OnStart() {
	alt.LogColored("~g~Go Resource Started")
	chatBroadcast := alt.Import("chat", "broadcast").(alt.MValueFunc)
	chatSend := alt.Import("chat", "send").(alt.MValueFunc)
	chatRegisterCmd := alt.Import("chat", "registerCmd").(alt.MValueFunc)
	//chatMutePlayer := alt.Import("chat", "mutePlayer").(alt.MValueFunc)

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
		p.SetMetaData("vehicles", make([]*alt.Vehicle, 0))

		spawn := getRandomSpawn()

		p.Spawn(spawn, 0)

		p.Emit("freeroam:spawned")
		p.Emit("freeroam:Interiors")

		//TODO fix timers and replace it with timeout 1000
		go func() {
			if p != nil && p.Valid() {
				playerCount := len(alt.GetPlayers())
				chatBroadcast(fmt.Sprintf("{1cacd4}%v {ffffff}has {00ff00}joined {ffffff}the Server.. (%v players online)", p.Name(), playerCount))
				chatSend(p, "{80eb34}Press {34dfeb}T {80eb34}and type {34dfeb}/help {80eb34}to see all available commands..")
			}
			//TODO clear timeout to reduce timer count
		}()

		// mp_m_freemode_01 hashed
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
			p.Emit("freeroam:switchInOutPlayer", true)
			p.ClearBloodDamage()
		}
		//TODO clear timer
		//}()

		if killer != nil {
			if player, ok := killer.(*alt.Player); ok {
				alt.LogInfo(player.Name(), "gave", p.Name(), "the rest!")
				SendNotificationToAllPlayer(fmt.Sprintf(`~r~<C>%v</C> ~s~killed ~b~<C>%v</C>`, player.Name(), p.Name()), 0, 2, false)
			} else if vehicle, ok := killer.(*alt.Vehicle); ok {
				alt.LogInfo(vehicle.NumberPlateText(), "(Vehicle) gave", p.Name(), "the rest!")
				SendNotificationToAllPlayer(fmt.Sprintf(`~r~<C>%v</C> (Vehicle) ~s~killed ~b~<C>%v</C>`, vehicle.NumberPlateText(), p.Name()), 0, 2, false)
			}
		} else {
			alt.LogInfo(p.Name(), "died!")
			SendNotificationToAllPlayer(fmt.Sprintf("~s~Suicide ~b~%v", p.Name()), 0, 2, false)
		}
	})

	alt.On.PlayerDisconnect(func(p *alt.Player, reason string) {
		playerCount := len(alt.GetPlayers())
		chatBroadcast(fmt.Sprintf("%v has left the Server.. (%v players online)", p.Name(), playerCount))
		vehicles := p.GetMetaData("vehicles")
		vehs, ok := vehicles.([]*alt.Vehicle)
		if ok {
			for _, vehicle := range vehs {
				vehicle.Destroy()
			}
		}
		p.DeleteMetaData("vehicles")
	})

	// =============================== Commands Begin ==================================================

	chatRegisterCmd("help", func(args ...interface{}) interface{} {
		player, ok := args[0].(*alt.Player)
		if !ok {
			return nil
		}
		chatSend(player, "{ff0000}========== {eb4034}HELP {ff0000} ==========")
		chatSend(player, "{ff0000}= {34abeb}/veh {40eb34}(model)   {ffffff} Spawn a Vehicle")
		chatSend(player, "{ff0000}= {34abeb}/tp {40eb34}(targetPlayer)   {ffffff} Teleport to Player")
		chatSend(player, "{ff0000}= {34abeb}/model {40eb34}(modelName)   {ffffff} Change Player Model")
		chatSend(player, "{ff0000}= {34abeb}/weapon {40eb34}(weaponName)   {ffffff} Get specified weapon")
		chatSend(player, "{ff0000}= {34abeb}/weapons    {ffffff} Get all weapons")
		chatSend(player, "{ff0000} ========================")
		return nil
	})

	chatRegisterCmd("veh", func(args ...interface{}) interface{} {
		player, ok := args[0].(*alt.Player)
		if !ok {
			return nil
		}

		if len(args) < 2 {
			chatSend(player, "Usage: /veh (vehicleModel)")
			return nil
		}

		model, ok := args[1].(string)
		if !ok {
			chatSend(player, "Usage: /veh (vehicleModel)")
			return nil
		}

		pos := player.Position()
		vehicle, err := alt.CreateVehicle(alt.Hash(model), pos, alt.Vector3{X: 0, Y: 0, Z: 0})
		if err != nil {
			chatSend(player, fmt.Sprintf(`{ff0000} Vehicle Model {ff9500}%v {ff0000}does not exist..`, model))
			alt.LogError(err.Error())
			return nil
		}

		vehicle.SetNumberplateText("GO<3")

		pVehs := player.GetMetaData("vehicles")
		playerVehicles, ok := pVehs.([]*alt.Vehicle)
		if !ok {
			return nil
		}

		if len(playerVehicles) >= 3 {
			toDestroy := playerVehicles[0]
			playerVehicles = playerVehicles[1:]

			if toDestroy != nil && toDestroy.Valid() {
				toDestroy.Destroy()
			}
		}

		playerVehicles = append(playerVehicles, vehicle)
		player.SetMetaData("vehicles", playerVehicles)

		return nil
	})

	chatRegisterCmd("pos", func(args ...interface{}) interface{} {
		player, ok := args[0].(*alt.Player)
		if !ok {
			return nil
		}

		pos := player.Position()
		alt.LogInfo("Position:", pos)
		chatSend(player, fmt.Sprintf("Position: %v, %v, %v", pos.X, pos.Y, pos.Z))
		return nil
	})

	chatRegisterCmd("tp", func(args ...interface{}) interface{} {
		player, ok := args[0].(*alt.Player)
		if !ok {
			return nil
		}

		if len(args) < 2 {
			chatSend(player, "Usage: /tp (target player)")
			return nil
		}

		targetName, ok := args[1].(string)
		if !ok {
			chatSend(player, "Usage: /tp (target player)")
			return nil
		}

		targetPlayers := alt.GetPlayersByName(targetName)

		if len(targetPlayers) < 1 {
			chatSend(player, fmt.Sprintf("{ff0000} Player {ff9500}%v {ff0000}not found..", targetName))
			return nil
		}

		player.SetPosition(targetPlayers[0].Position())
		chatSend(player, fmt.Sprintf("You got teleported to {1cacd4}%v{ffffff}", targetName))

		return nil
	})

	chatRegisterCmd("model", func(args ...interface{}) interface{} {
		player, ok := args[0].(*alt.Player)
		if !ok {
			return nil
		}

		if len(args) < 2 {
			chatSend(player, "Usage: /model (modelName)")
			return nil
		}

		model, ok := args[1].(string)
		if !ok {
			chatSend(player, "Usage: /model (modelName)")
			return nil
		}

		player.SetModel(alt.Hash(model))

		return nil
	})

	chatRegisterCmd("weapon", func(args ...interface{}) interface{} {
		player, ok := args[0].(*alt.Player)
		if !ok {
			return nil
		}

		if len(args) < 2 {
			chatSend(player, "Usage: /weapon (modelName)")
			return nil
		}

		model, ok := args[1].(string)
		if !ok {
			chatSend(player, "Usage: /weapon (modelName)")
			return nil
		}

		player.GiveWeapon(alt.Hash("weapon_"+model), 500, true)

		return nil
	})

	chatRegisterCmd("weapons", func(args ...interface{}) interface{} {
		player, ok := args[0].(*alt.Player)
		if !ok {
			return nil
		}

		for _, weapon := range WeaponModels {
			player.GiveWeapon(alt.Hash("weapon_"+weapon), 500, true)
		}

		return nil
	})

	chatRegisterCmd("vehmax", func(args ...interface{}) interface{} {
		player, ok := args[0].(*alt.Player)
		if !ok {
			return nil
		}

		vehicle := player.Vehicle()

		if vehicle == nil || !vehicle.Valid() {
			alt.LogError("Player is not in a vehicle")
			chatSend(player, "Usage: /vehmax [only when in vehicle]")
			return nil
		}

		err := maxVehicle(vehicle)

		if err != nil {
			alt.LogError(err.Error())
			return nil
		}

		return nil
	})
}

//export OnStop
func OnStop() {
	alt.LogColored("~r~Go Resource Stopped")
}
