package types

type ClusterId string
type ClusterStatus string

const (
	ClusterStatus_NEW      ClusterStatus = "NEW"
	ClusterStatus_CREATING ClusterStatus = "CREATING"
	ClusterStatus_ERROR    ClusterStatus = "ERROR"
	ClusterStatus_READY    ClusterStatus = "READY"
)

type ClusterSpecs struct {
	PublicKey string           `json:"public_key"`
	NodePools []*NodePoolSpecs `json:"node_pools"`
}

type NodePoolSpecs struct {
}
