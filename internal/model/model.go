package model

// PluginInfo - info received from plugin
type PluginInfo struct {
	PluginType    string
	PluginVersion string
	CliType       string
	CliVersion    string
	DeviceName    string
	Events        []Events
}

type Events struct {
	Uid       string
	CreatedAt string
	Type      string
	Project   string
	Language  string
	Target    string
	Branch    string
	Params    map[string]any
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
