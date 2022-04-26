package main

import "C"
import (
	"afterlife-rp.de/go/core/commands"
	"github.com/timo972/altv-go-pkg/alt"
)

func main() {}

type VehicleInfo struct {
	Vehicle *alt.Vehicle `alt:"vehicle"`
	Price   uint64       `alt:"price"`
	Name    string       `alt:"name"`
	Labels  [][]string   `alt:"labels"`
}

//export OnStart
func OnStart() {
	alt.LogColored("~g~Go Resource Started")

	alt.SetMetaData("name", alt.Resource.Name)

	commands.Register()

	veh, err := alt.CreateVehicle(alt.Hash("adder"), alt.Vector3{X: -10, Y: 0, Z: 0}, alt.Vector3{X: 0, Y: 0, Z: 0})
	if err != nil {
		alt.LogError(err.Error())
		return
	}

	alt.LogInfo("Vehicle ID:", veh.ID())

	var resourceName string
	alt.MetaData("name", &resourceName)

	alt.LogInfo("Resource Name:", resourceName)

	vInfo := VehicleInfo{
		Vehicle: veh,
		Price:   1000,
		Name:    "Adder",
		Labels:  [][]string{{"test", "test2"}, {"hello", "world"}},
	}

	// alt.LogInfo(fmt.Sprintf("Vehicle Ptr: %v", veh.Ptr))
	alt.SetMetaData("vehicle", vInfo)

	// alt.Emit("vInfo", vInfo)
	err = alt.Emit("vInfo", vInfo)
	if err != nil {
		alt.LogError("Emit err:", err.Error())
	}

	var vInfo2 VehicleInfo
	ok := alt.MetaData("vehicle", &vInfo2)
	if ok {
		// alt.LogInfo(fmt.Sprintf("vInfo2: %v, Ptr: %v, Valid: %v, ID: %v", vInfo2, vInfo2.Ptr, vInfo2.Valid(), vInfo2.ID()))
		alt.LogInfo("vInfo2:", vInfo2.Labels, vInfo2.Name, vInfo2.Price, vInfo2.Vehicle)
	} else {
		alt.LogError("vInfo2 not found")
	}

	alt.On.ServerEvent("user::login", func(userName string, password string, v map[string]interface{}) {
		alt.LogInfo("User login:", userName, password)
		alt.LogInfo("Users vehicle:", v)
	})

	alt.Emit("user::login", "Timo", "123", vInfo2)

	/*veh.SetMetaData("test", map[string]VehicleInfo{"Hello": vInfo})

	var isTest = map[string]VehicleInfo{}
	ok := veh.MetaData("test", &isTest)
	if !ok {
		alt.LogWarning("MetaData not found")
	}
	alt.LogInfo("This is a test:", isTest)*/

	//veh.SetMetaData("vehicleInfo", vInfo)
	//
	//var vInfo2 VehicleInfo
	//veh.MetaData("vehicleInfo", &vInfo2)
	//alt.LogWarning("Vehicle Info:", fmt.Sprintf("VehicleInfo{Name: %v, Price: %v, Labels: %v, Vehicle: %v}", vInfo2.Name, vInfo2.Price, vInfo2.Labels, vInfo2.Vehicle.String()))
	//alt.Emit("foo", veh, "Hello World", 7, true, "f", "b", "x", 9, "z")
}

//export OnStop
func OnStop() {
	alt.LogColored("~r~Go Resource Stopped")
}
