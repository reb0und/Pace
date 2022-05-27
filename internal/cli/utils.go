package cli

import "os"

const (
	ClearConsole = "\u001Bc"
	Black        = "\u001B[0;30m"
	Red          = "\u001B[0;31m"
	Green        = "\u001B[0;32m"
	Yellow       = "\u001B[0;33m"
	Blue         = "\u001B[0;34m"
	Purple       = "\u001B[0;35m"
	Cyan         = "\u001B[0;36m"
	White        = "\u001B[0;37m"
	BlackBold    = "\u001B[1;30m"
	RedBold      = "\u001B[1;31m"
	GreenBold    = "\u001B[1;32m"
	YellowBold   = "\u001B[1;33m"
	BlueBold     = "\u001B[1;34m"
	PurpleBold   = "\u001B[1;35m"
	CyanBold     = "\u001B[1;36m"
	WhiteBold    = "\u001B[1;37m"
	Reset        = "\u001B[0m"
)

func Clear() {
	_, err := os.Stdout.Write([]byte(ClearConsole))
	if err != nil {
		return
	}
}
