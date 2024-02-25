package low

import (
	"strconv"
	"time"

	config "github.com/238Studio/child-nodes-config-service"
	database "github.com/238Studio/child-nodes-database-service"
	device "github.com/238Studio/child-nodes-device-service"
	"github.com/238Studio/child-nodes-error-manager/errpack"
	"github.com/238Studio/child-nodes-net-service/ws"
)

//初始化次序:配置文件管理器、数据库管理器、websocket(网络管理器)、设备管理器

// InitLowLevel 初始化下层
// 传入:无
// 传出:错误
func InitLowLevel() {
	initConfigService()

	//初始化配置文件
	err := registerInitConfigService()
	if err != nil {
		panic(err)
	}

	err = initDataBaseService()
	if err != nil {
		panic(err)
	}

	err = initWebsocket()
	if err != nil {
		panic(err)
	}

	err = initDeviceService()
	if err != nil {
		panic(err)
	}
}

//TODO:字段规划,表单规划

// 初始化配置文件管理器
// 传入:无
// 传出:错误
func initConfigService() {
	configManger = config.InitConfigManager(configPath)
}

// 注册初始化配置文件
// 传入:无
// 传出:错误
func registerInitConfigService() error {
	err := configManger.InitModuleConfig(databaseConfig)
	if err != nil {
		return err
	}

	err = configManger.InitModuleConfig(websocketConfig)
	if err != nil {
		return err
	}

	err = configManger.InitModuleConfig(deviceConfig)
	if err != nil {
		return err
	}

	return nil
}

// 初始化数据库管理器
// 传入:无
// 传出:错误
func initDataBaseService() error {
	databaseName, err := configManger.ReadConfig(databaseConfig, configDataBaseName)
	if err != nil {
		return err
	}

	databaseService, err = database.InitSQLiteDatabase(databaseName, "")
	if err != nil {
		return err
	}

	return nil
}

// 初始化websocket
// 传入:无
// 传出:错误
func initWebsocket() error {
	wsURL, err := configManger.ReadConfig(websocketConfig, configWsUrl)
	if err != nil {
		return err
	}

	pingTimeString, err := configManger.ReadConfig(websocketConfig, configPing)
	if err != nil {
		return err
	}
	ping, err := strconv.Atoi(pingTimeString)
	if err != nil {
		return errpack.NewError(errpack.TrivialException, 0, err) //TODO:错误模块
	}

	pongTimeString, err := configManger.ReadConfig(websocketConfig, configPong)
	if err != nil {
		return err
	}
	pong, err := strconv.Atoi(pongTimeString)
	if err != nil {
		return err
	}

	websocketServiceAPP, err = ws.InitWebsocketService(wsURL, ping, pong)
	if err != nil {
		return errpack.NewError(errpack.TrivialException, 0, err) //TODO:错误模块
	}

	return nil
}

// 初始化设备管理器
// 传入:无
// 传出:错误
func initDeviceService() error {
	baudString, err := configManger.ReadConfig(deviceConfig, configBaud)
	if err != nil {
		return err
	}
	baud, err := strconv.Atoi(baudString)
	if err != nil {
		return errpack.NewError(errpack.TrivialException, 0, err) //TODO:错误模块
	}

	readTimeOutString, err := configManger.ReadConfig(deviceConfig, configReadTimeOut)
	if err != nil {
		return err
	}
	readTimeOut, err := time.ParseDuration(readTimeOutString)
	if err != nil {
		return errpack.NewError(errpack.TrivialException, 0, err) //TODO:错误模块
	}

	maxResendTimesString, err := configManger.ReadConfig(deviceConfig, configMaxResendTimes)
	if err != nil {
		return err
	}
	maxResendTimes, err := strconv.Atoi(maxResendTimesString)
	if err != nil {
		return errpack.NewError(errpack.TrivialException, 0, err) //TODO:错误模块
	}

	revBufferWaitTimeOutString, err := configManger.ReadConfig(deviceConfig, configRevBufferWaitTimeOut)
	if err != nil {
		return err
	}
	revBufferWaitTimeout, err := strconv.ParseInt(revBufferWaitTimeOutString, 10, 64)
	if err != nil {
		return errpack.NewError(errpack.TrivialException, 0, err) //TODO:错误模块
	}

	SendBufferWaitTimeOutString, err := configManger.ReadConfig(deviceConfig, configSendBufferWaitTimeOut)
	if err != nil {
		return err
	}
	SendBufferWaitTimeOut, err := strconv.ParseInt(SendBufferWaitTimeOutString, 10, 64)
	if err != nil {
		return errpack.NewError(errpack.TrivialException, 0, err) //TODO:错误模块
	}

	SerialApp = device.InitSerialApp(baud, readTimeOut, maxResendTimes, revBufferWaitTimeout, SendBufferWaitTimeOut)

	return nil
}
