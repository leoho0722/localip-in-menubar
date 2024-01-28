package menubar

import (
	"fmt"
	"github.com/caseymrm/menuet"
	"leoho.io/localip-in-menubar/network"
	"time"
)

// Start starts the menubar app
func Start() {
	go setupMenuBar()
	menuet.App().RunApplication()
}

// setupMenuBar sets up the menubar
func setupMenuBar() {
	for {
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: fmt.Sprintf("IP: %s", network.GetLocalNetworkIPv4Address()),
		})
		time.Sleep(time.Second)
	}
}
