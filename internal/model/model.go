package model

// PluginInfo - info received from plugin
type PluginInfo struct {
	Uid           string
	PluginType    string
	PluginVersion string
	IdeType       string
	IdeVersion    string
	Events        []Events
}

type Events struct {
	CreatedAt      string
	Type           string
	Project        string
	ProjectBaseDir string
	Language       string
	Target         string
	Branch         string
	Params         map[string]any
}

// AggregatorInfo - info about project from aggregator
type AggregatorInfo struct {
	GitBranchesByProjectBaseDir map[string]string
}

// DataModel - internal structure for store data before dispatching
type DataModel struct {
	PluginInfo     PluginInfo
	AggregatorInfo AggregatorInfo
}
