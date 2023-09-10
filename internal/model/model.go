package model

// PluginInfo - info received from plugin
type PluginInfo struct {
	PluginType    string
	PluginVersion string
	CliType       string
	CliVersion    string
	OSName        string
	IdeType       string
	IdeVersion    string
	Events        []Events
}

type Events struct {
	Uid            string
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
	OSName                      string
	Uid                         string
}

// DataModel - internal structure for store data before dispatching
type DataModel struct {
	PluginInfo     PluginInfo
	AggregatorInfo AggregatorInfo
}
