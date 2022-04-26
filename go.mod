module afterlife-rp.de/go/core

go 1.18

// require github.com/shockdev04/altv-go-pkg v0.0.0-20210812162411-ed5f7c81bf1d // indirect
replace github.com/timo972/altv-go-pkg => ../../altv-go-pkg

replace altv-updater => ../../../OpenSource\pkg.go\altv-updater

require altv-updater v0.0.0-0-0

require (
	golang.org/x/exp v0.0.0-20220414153411-bcd21879b8fd // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jedib0t/go-pretty/v6 v6.3.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/spf13/cobra v1.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	//	github.com/shockdev04/altv-go-pkg v0.0.0-20210901183431-3250d7261696
	github.com/timo972/altv-go-pkg v0.0.0-20220413145743-bf3a9446ee38 // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
)
