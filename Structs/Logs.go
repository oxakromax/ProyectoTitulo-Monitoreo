package Structs

import "time"

type LogMessage struct {
	Level           string      `json:"Level"`
	WindowsIdentity string      `json:"WindowsIdentity"`
	ProcessName     string      `json:"ProcessName"`
	TimeStamp       time.Time   `json:"TimeStamp"`
	Message         string      `json:"Message"`
	JobKey          string      `json:"JobKey"`
	RawMessage      string      `json:"RawMessage"`
	RobotName       string      `json:"RobotName"`
	HostMachineName string      `json:"HostMachineName"`
	MachineId       int         `json:"MachineId"`
	RuntimeType     interface{} `json:"RuntimeType"`
	Id              int         `json:"Id"`
}
type LogResponse struct {
	OdataContext string       `json:"@odata.context"`
	OdataCount   int          `json:"@odata.count"`
	Value        []LogMessage `json:"value"`
}
