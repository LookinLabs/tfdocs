package tfData

type Instance struct {
	Attributes map[string]interface{} `json:"attributes"`
}

type Resource struct {
	Type          string     `json:"type"`
	Name          string     `json:"name"`
	Instances     []Instance `json:"instances"`
	Mode          string     `json:"mode"`
	ProviderName  string     `json:"provider_name"`
	ModuleAddress string     `json:"module_address"`
}

type TFStateResources struct {
	Resources []Resource `json:"resources"`
}
