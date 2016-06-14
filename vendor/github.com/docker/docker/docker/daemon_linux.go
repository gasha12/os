// +build daemon

package docker

import (
	systemdDaemon "github.com/coreos/go-systemd/daemon"
)

// notifySystem sends a message to the host when the server is ready to be used
func notifySystem() {
	// Tell the init daemon we are accepting requests
	go systemdDaemon.SdNotify("READY=1")
}
