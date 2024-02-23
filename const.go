package low

import (
	config "github.com/238Studio/child-nodes-config-service"
	database "github.com/238Studio/child-nodes-database-service"
	device "github.com/238Studio/child-nodes-device-service"
	"github.com/238Studio/child-nodes-net-service/ws"
)

const (
	configPath = "./config" //配置文件路径

	databaseConfig  = "database"
	websocketConfig = "websocket"
	deviceConfig    = "device"

	configDataBaseName = "database_name"

	configWsUrl = "ws_url"
	configPing  = "ping"
	configPong  = "pong"

	configBaud                  = "baud"
	configReadTimeOut           = "read_time_out"
	configMaxResendTimes        = "max_resend_times"
	configRevBufferWaitTimeOut  = "rev_buffer_wait_time_out"
	configSendBufferWaitTimeOut = "send_buffer_wait_time_out"
)

var (
	configManger        *config.ConfigManager
	databaseService     *database.DatabaseAPP
	websocketServiceAPP *ws.WebsocketServiceApp
	SerialApp           *device.SerialApp
)
