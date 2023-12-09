package types

type MachineId string

type MachineSpecs struct {
	CPU           int    `json:"cpu"`
	RAM           int    `json:"ram"`
	User_data_b64 string `json:"user_data_b64"`
}

type MachineStatus string

const (
	MachineStatus_NEW         MachineStatus = "NEW"
	MachineStatus_CREATING    MachineStatus = "CREATING"
	MachineStatus_RUNNING     MachineStatus = "RUNNING"
	MachineStatus_TERMINATING MachineStatus = "TERMINATING"
	MachineStatus_TERMINATED  MachineStatus = "TERMINATED"
	MachineStatus_ERROR       MachineStatus = "ERROR"
)
