package model

import "gorm.io/gorm"

type NetInfo struct {
	gorm.Model

	NPeer int `json:"n_peer"`
	BlockHeight int `gorm:"index:,sort:desc" json:"block_height"`
}

type Peer struct {
	gorm.Model

	BlockHeight int `gorm:"index:,sort:desc" json:"block_height"`

	RemoteIP         string               `json:"remote_ip"`
	DefaultNodeID string `json:"default_node_id"`
	Score int `json:"score"`

	// NodeInfo         p2p.DefaultNodeInfo  `json:"node_info"`
	// IsOutbound       bool                 `json:"is_outbound"`
	// ConnectionStatus p2p.ConnectionStatus `json:"connection_status"`
}
