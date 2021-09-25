package menubar

import (
	"github.com/getlantern/systray"
	"github.com/krlc/bravo/icon"
	"wireguard-state/connectionstate"
)

type MenuBar struct {
	connectionItem *systray.MenuItem
	rxItem         *systray.MenuItem
	txItem         *systray.MenuItem
	quitItem       *systray.MenuItem
}

func New() *MenuBar {
	menuBar := &MenuBar{}
	menuBar.connectionItem = addMenuItem("Connect", false)
	systray.AddSeparator()

	menuBar.rxItem = addMenuItem("Rx", true)
	menuBar.txItem = addMenuItem("Tx", true)

	menuBar.quitItem = addMenuItem("Quit", false)

	return menuBar
}

func (menuBar *MenuBar) OnConnectionClick() chan struct{} {
	return menuBar.connectionItem.ClickedCh
}

func (menuBar *MenuBar) OnQuitClick() chan struct{} {
	return menuBar.quitItem.ClickedCh
}

func (menuBar *MenuBar) Update(state *connectionstate.ConnectionState) {
	if state.Up {
		systray.SetIcon(icon.DataEnabled)
		systray.SetTooltip("WireGuard is running")

		menuBar.connectionItem.Check()
		menuBar.connectionItem.SetTitle("Connected")

		menuBar.rxItem.Show()
		menuBar.txItem.Show()

		menuBar.rxItem.SetTitle("Rx: " + state.Rx)
		menuBar.txItem.SetTitle("Tx: " + state.Tx)
	} else {
		systray.SetIcon(icon.DataDisabled)
		systray.SetTooltip("WireGuard is not running")

		menuBar.connectionItem.Uncheck()
		menuBar.connectionItem.SetTitle("Connect")

		menuBar.rxItem.Hide()
		menuBar.txItem.Hide()
	}
}

func addMenuItem(title string, disable bool) *systray.MenuItem {
	menuItem := systray.AddMenuItemCheckbox(title, "", false)
	if disable {
		menuItem.Disable()
	}

	return menuItem
}
