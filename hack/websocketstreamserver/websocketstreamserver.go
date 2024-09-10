package websocketstreamserver

import (
	"github.com/Yokosa/web-config-server/config"
	"github.com/Yokosa/web-config-server/network"
	"github.com/Yokosa/web-config-server/service"
)

const ID = "websocket-stream-server"
const FriendlyName = "Websocket Stream Server"

const frameGrabberMainstreamService = "websocket-stream-server-framegrabber-mainstream"
const frameGrabberSubstreamService = "websocket-stream-server-framegrabber-substream"
const mainServiceName = "websocket-stream-server-websocket-stream-server"

type WebsocketStreamConfig struct {
	Enable bool `json:"enable"`
}

func GetServiceStatus() service.ServiceStatus {
	return service.Status(FriendlyName, service.Runit, mainServiceName)
}

func SaveConfig(websocketStreamConfig WebsocketStreamConfig) bool {

	success := config.Save(ID, websocketStreamConfig)

	if !success {
		return false
	}

	if websocketStreamConfig.Enable {
		config.EnableService(ID)

		service.Restart(service.Perp, "fetch_av")

		service.Restart(service.Runit, frameGrabberMainstreamService)
		service.Restart(service.Runit, frameGrabberSubstreamService)

		service.Restart(service.Runit, mainServiceName)
	} else {
		config.DisableService(ID)

		service.Restart(service.Perp, "fetch_av")

		service.Stop(service.Runit, frameGrabberMainstreamService)
		service.Stop(service.Runit, frameGrabberSubstreamService)

		service.Stop(service.Runit, mainServiceName)
	}

	return true
}

func Info() string {
	return ``
}

func Endpoints() string {
	ip := network.GetIP()

	return `{
		"main": "ws://` + ip + `:558/main",
		"sub": "ws://` + ip + `:558/sub"
	}`
}
