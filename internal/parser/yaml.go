package parser

type yamlRoot struct {
	Validators map[string]yamlValidator `yaml:"validators"`
	Tests      map[int]map[string]bool  `yaml:"tests"`
}

type yamlValidator struct {
	Compare        *yamlValidatorCompare `yaml:"compare"`
	Counter        *yamlValidatorCounter `yaml:"counter"`
	Parity         *yamlValidatorParity  `yaml:"parity"`
	HasMoreParity  bool                  `yaml:"has_more_parity"`
	HasRepetitions bool                  `yaml:"has_repetitions"`
	HasPair        bool                  `yaml:"has_pair"`
	GreatestItem   bool                  `yaml:"greatest_item"`
	LeastItem      bool                  `yaml:"least_item"`
	HasOrder       bool                  `yaml:"has_order"`
}

type yamlValidatorCompare struct {
	Item   string   `yaml:"item"`
	Sum    []string `yaml:"sum"`
	Target string   `yaml:"target"`
}

type yamlValidatorCounter struct {
	Number int    `yaml:"number"`
	Parity string `yaml:"parity"`
}

type yamlValidatorParity struct {
	Item string `yaml:"item"`
	Sum  bool   `yaml:"sum"`
}
