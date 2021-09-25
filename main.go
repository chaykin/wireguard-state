package main

import (
	"github.com/getlantern/systray"
	"time"
	"wireguard-state/config"
	"wireguard-state/connectionstate"
	"wireguard-state/menubar"
)

const (
	configFile = "wgs.yaml"
)

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	cfg := config.ReadConfig(configFile)
	state := &connectionstate.ConnectionState{}
	menu := menubar.New()
	ticker := time.NewTicker(time.Duration(cfg.RefreshRate) * time.Second)

	go processRequests(cfg, state, menu, ticker)
}

func processRequests(config *config.WgsConfig, state *connectionstate.ConnectionState, menubar *menubar.MenuBar, ticker *time.Ticker) {
	onUpdateState(config, state, menubar)

	for {
		select {
		case <-menubar.OnConnectionClick():
			onToggleConnection(config, state, menubar)
		case <-menubar.OnQuitClick():
			onExit(ticker)
		case <-ticker.C:
			onUpdateState(config, state, menubar)
		}
	}
}
