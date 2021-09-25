package main

import (
	"github.com/getlantern/systray"
	"log"
	"os/exec"
	"time"
	"wireguard-state/config"
	"wireguard-state/connectionstate"
	"wireguard-state/menubar"
)

func onToggleConnection(config *config.WgsConfig, state *connectionstate.ConnectionState, menubar *menubar.MenuBar) {
	connectionCommand := config.WgUpCommand
	if state.Up {
		connectionCommand = config.WgDownCommand
	}

	err := exec.Command("bash", "-c", connectionCommand).Run()
	if err != nil {
		log.Println("Error toggle connecion:", err)
	} else {
		onUpdateState(config, state, menubar)
	}
}

func onExit(ticker *time.Ticker) {
	ticker.Stop()
	systray.Quit()
}

func onUpdateState(config *config.WgsConfig, state *connectionstate.ConnectionState, menubar *menubar.MenuBar) {
	state.Update(config)
	menubar.Update(state)
}
