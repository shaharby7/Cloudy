package vms

// type Machine_id string

type CreateVmOptions struct {
	User_data_b64  string `json:"user_data_b64"`
	IpAllocationId string `json:"ip_allocation_id"`
}

type DeleteVmOptions struct {
	Machine_id string `json:"machine_id"`
}

type Result struct {
	Machine_id string `json:"machine_id"`
}
type GetVmOptions struct {
	Machine_id string `json:"machine_id"`
}

type GetVmResults struct {
	Machine_id string `json:"machine_id"`
	IP         string `json:"ip"`
}
