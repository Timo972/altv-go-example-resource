package commands

import (
	"github.com/timo972/altv-go-pkg/alt"
)

func Register() {
	registerCmd, _ := alt.Import[alt.ExternFunction]("chat", "registerCmd")

	registerCmd.Call("respawn", handleRespawn)
	registerCmd.Call("spawn", handleRespawn)

	alt.On.ConsoleCommand(func(command string, args []string) {
		alt.LogInfo("[Console]", command, args)
	})
}
