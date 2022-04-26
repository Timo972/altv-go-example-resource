package commands

import (
	"afterlife-rp.de/go/core/utility"
	"github.com/timo972/altv-go-pkg/alt"
)

func handleRespawn(p *alt.Player, args []string) {
	alt.LogInfo("Respawning player")

	if !p.Valid() {
		alt.LogWarning("Player is not valid")
		return
	}

	pos := p.Position()
	randomPosition := utility.RandomPositionAround(pos, 2, 8)
	p.Spawn(randomPosition, 3000)
}
