package model

// PluginInfo - info received from plugin
type PluginInfo struct {
	Events []Events
}

type Events struct {
	Id             string
	CreatedAt      string
	Type           string
	Project        string
	ProjectBaseDir string
	Language       string
	Target         string
	Branch         string
	Timezone       string
	Params         map[string]any
}

// AggregatorInfo - info about project from aggregator
type AggregatorInfo struct {
	GitBranchesByProjectBaseDir map[string]string
	Id                          []string
}

// DataModel - internal structure for store data before dispatching
type DataModel struct {
	PluginInfo     PluginInfo
	AggregatorInfo AggregatorInfo
}
