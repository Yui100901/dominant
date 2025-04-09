package data

import (
	"dominant/domain/impl/data/common"
	"dominant/global"
	"dominant/persistence"
	"dominant/persistence/model"
	"github.com/Yui100901/MyGo/log_utils"
	"github.com/Yui100901/MyGo/struct_utils"
)

//
// @Author yfy2001
// @Date 2025/3/12 20 49
//

// SaveDeviceTelemetry 保存设备传数据
// 这里传入的是设备id
func SaveDeviceTelemetry(id string, m common.DeviceMessage) {
	telemetry := m.ConvertToTelemetry(id)
	// 同时存入历史表和最新表
	SaveTelemetry(telemetry)
	SaveTelemetryLatest(telemetry)
}

func SaveDevice(device *common.Device) error {
	pDevice := &model.Device{}
	err := struct_utils.ConvertStruct(device, pDevice)
	if err != nil {
		log_utils.Error.Println(err)
		return err
	}
	return persistence.SaveOrUpdate(pDevice)
}

func GetDeviceByIdFromMap(id string) *common.Device {
	device, ok := global.DeviceMap.Get(id)
	if ok {
		return device
	}
	return &common.Device{}
}

func GetDeviceByIdFromPersistence(id string) *common.Device {
	pDevice, err := persistence.GetByID[model.Device](id)
	if err != nil {
		return nil
	}
	device := &common.Device{}
	err = struct_utils.ConvertStruct(pDevice, device)
	if err != nil {
		log_utils.Error.Println(err)
		return nil
	}
	return device
}

func GetDeviceList() []common.Device {
	pDeviceList, err := persistence.GetList[model.Device]()
	if err != nil {
		log_utils.Error.Println(err)
		return nil
	}
	var deviceList []common.Device
	for _, pDevice := range pDeviceList {
		device := common.Device{}
		err := struct_utils.ConvertStruct(pDevice, &device)
		if err != nil {
			log_utils.Error.Println(err)
			return nil
		}
		deviceList = append(deviceList, device)
	}
	return deviceList
}

func DeleteDevice(idList []string) {
	deleteList := make([]any, len(idList))
	for _, id := range idList {
		deleteList = append(deleteList, id)
	}
	err := persistence.BatchDeleteByIdList[model.Device](deleteList)
	if err != nil {
		log_utils.Error.Println(err)
		return
	}
}

func CacheAllDevice() {
	deviceList := GetDeviceList()
	for _, device := range deviceList {
		global.DeviceMap.Set(device.ID, &device)
	}
}

func SaveTelemetry(telemetry *common.Telemetry) {
	pTelemetry := &model.Telemetry{}
	err := struct_utils.ConvertStruct(telemetry, pTelemetry)
	//log_utils.Info.Printf("Telemetry %+v\n,Position%+v,Status %+v", telemetry, telemetry.Position, telemetry.Status)
	//log_utils.Info.Printf("PTelemetry %+v\n,Position%+v,Status %+v", pTelemetry, pTelemetry.Position, pTelemetry.Status)
	if err != nil {
		log_utils.Error.Println(err)
		return
	}
	err = persistence.SaveOrUpdate[model.Telemetry](pTelemetry)
	if err != nil {
		log_utils.Error.Println(err)
		return
	}
}

func SaveTelemetryLatest(telemetry *common.Telemetry) {
	pTelemetryLatest := &model.TelemetryLatest{}
	err := struct_utils.ConvertStruct(telemetry, pTelemetryLatest)
	if err != nil {
		log_utils.Error.Println(err)
		return
	}
	err = persistence.SaveOrUpdate[model.TelemetryLatest](pTelemetryLatest)
	if err != nil {
		log_utils.Error.Println(err)
		return
	}
}

func GetTelemetryLatestList() []common.Telemetry {
	pTelemetryList, err := persistence.GetList[model.TelemetryLatest]()
	if err != nil {
		log_utils.Error.Println(err)
		return nil
	}
	var telemetryList []common.Telemetry
	for _, pTelemetry := range pTelemetryList {
		telemetry := common.Telemetry{}
		err := struct_utils.ConvertStruct(pTelemetry, &telemetry)
		if err != nil {
			log_utils.Error.Println(err)
			return nil
		}
		telemetryList = append(telemetryList, telemetry)
	}
	return telemetryList
}

func SaveTelemetryVirtual(telemetry *common.Telemetry) {
	pTelemetryVirtual := &model.TelemetryVirtual{}
	err := struct_utils.ConvertStruct(telemetry, pTelemetryVirtual)
	if err != nil {
		log_utils.Error.Println(err)
		return
	}
	err = persistence.SaveOrUpdate[model.TelemetryVirtual](pTelemetryVirtual)
	if err != nil {
		log_utils.Error.Println(err)
		return
	}
}

func GetTelemetryVirtualList() []common.Telemetry {
	pTelemetryList, err := persistence.GetList[model.TelemetryVirtual]()
	if err != nil {
		log_utils.Error.Println(err)
		return nil
	}
	var telemetryList []common.Telemetry
	for _, pTelemetry := range pTelemetryList {
		telemetry := common.Telemetry{}
		err := struct_utils.ConvertStruct(pTelemetry, &telemetry)
		if err != nil {
			log_utils.Error.Println(err)
			return nil
		}
		telemetryList = append(telemetryList, telemetry)
	}
	return telemetryList
}
