package types

type MachineSpecs struct {
	CPU           int    `json:"cpu"`
	RAM           int    `json:"ram"`
	User_data_b64 string `json:"user_data_b64"`
}

type MachineStatus string

const (
	NEW         MachineStatus = "NEW"
	CREATING    MachineStatus = "CREATING"
	RUNNING     MachineStatus = "RUNNING"
	TERMINATING MachineStatus = "TERMINATING"
	TERMINATED  MachineStatus = "TERMINATED"
	ERROR       MachineStatus = "ERROR"
)
