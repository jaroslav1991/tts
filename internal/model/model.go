package model

import "time"

// PluginInfo - info received from plugin
type PluginInfo struct {
	Program     string
	Duration    time.Duration
	PathProject string
}

// AggregatorInfo - info about project from aggregator
type AggregatorInfo struct {
	CurrentGitBranch string
}

// DataModel - internal structure for store data before dispatching
type DataModel struct {
	PluginInfo     PluginInfo
	AggregatorInfo AggregatorInfo
}
