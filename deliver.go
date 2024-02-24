package low

import (
	"errors"

	config "github.com/238Studio/child-nodes-config-service"
	database "github.com/238Studio/child-nodes-database-service"
	device "github.com/238Studio/child-nodes-device-service"
	"github.com/238Studio/child-nodes-error-manager/errpack"
	"github.com/238Studio/child-nodes-net-service/ws"
)

// GetConfigManager 获取配置文件管理器对象
// 传入:无
// 传出:配置文件管理器对象，错误
func GetConfigManager() (*config.ConfigManager, error) {
	//防止空值
	if configManger == nil {
		return nil, errpack.NewError(errpack.CommonException, 0, errors.New("配置文件管理器对象为空")) //TODO:确定模块
	}

	return configManger, nil
}

// GetDatabaseService 获取数据库服务对象
// 传入:无
// 传出:数据库服务对象，错误
func GetDatabaseService() (*database.DatabaseAPP, error) {
	//防止空值
	if databaseService == nil {
		return nil, errpack.NewError(errpack.CommonException, 0, errors.New("数据库服务对象为空")) //TODO:确定模块
	}

	return databaseService, nil
}

// GetWebsocketServiceAPP 获取websocket服务对象
// 传入:无
// 传出:websocket服务对象，错误
func GetWebsocketServiceAPP() (*ws.WebsocketServiceApp, error) {
	//防止空值
	if websocketServiceAPP == nil {
		return nil, errpack.NewError(errpack.CommonException, 0, errors.New("websocket服务对象为空")) //TODO:确定模块
	}

	return websocketServiceAPP, nil
}

// GetSerialApp 获取串口服务对象
// 传入:无
// 传出:串口服务对象，错误
func GetSerialApp() (*device.SerialApp, error) {
	//防止空值
	if SerialApp == nil {
		return nil, errpack.NewError(errpack.CommonException, 0, errors.New("串口服务对象为空")) //TODO:确定模块
	}

	return SerialApp, nil
}
