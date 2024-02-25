package low

import (
	"testing"
)

func TestInitConfigService(t *testing.T) {
	initConfigService()

	err := registerInitConfigService()
	if err != nil {
		t.Error(err)
	}

	databaseName, err := configManger.ReadConfig(databaseConfig, configDataBaseName)
	if err != nil {
		t.Error(err)
	}
	t.Log(databaseName)
}

func TestInitDataBaseService(t *testing.T) {
	initConfigService()

	err := registerInitConfigService()
	if err != nil {
		t.Error(err)
	}

	err = initDataBaseService()
	if err != nil {
		t.Error(err)
	}

	var TestStruct struct {
		// ID      int
		UserName string
		PassWord string
		Avtar    string
	}
	err = databaseService.CreateTable("test1", &TestStruct)
	if err != nil {
		t.Error(err)
	}
}

func TestInitWebsocket(t *testing.T) {
	initConfigService()

	err := registerInitConfigService()
	if err != nil {
		t.Error(err)
	}

	err = initWebsocket()
	if err != nil {
		t.Error(err)
	}

	websocketServiceAPP.StartRead()
	test := websocketServiceAPP.InitModelMessageChan("test")
	test.StartWrite()
	test.WriteMessage("test", false, "hello world")

	t.Log(<-test.ReadMessage)
}

func TestInitDeviceService(t *testing.T) {
	initConfigService()

	err := registerInitConfigService()
	if err != nil {
		t.Error(err)
	}

	err = initDeviceService()
	if err != nil {
		t.Error(err)
	}
}
