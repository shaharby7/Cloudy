package types

type ClusterId string
type ClusterStatus string

const (
	ClusterStatus_NEW ClusterStatus = "NEW"
)

type ClusterSpecs struct {
	PublicKey string           `json:"public_key"`
	NodePools []*NodePoolSpecs `json:"node_pools"`
}

type NodePoolSpecs struct {
}
