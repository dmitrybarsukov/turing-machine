package parser

type yamlRoot struct {
	Validators map[string]yamlValidator `yaml:"validators"`
	Tests      map[int]map[string]bool  `yaml:"tests"`
}

type yamlValidator struct {
	Compare         *yamlValidatorCompare `yaml:"compare"`
	Count           *yamlValidatorCount   `yaml:"count"`
	Parity          *yamlValidatorParity  `yaml:"parity"`
	MajorParity     bool                  `yaml:"major_parity"`
	HasRepetitions  bool                  `yaml:"has_repetitions"`
	HasPair         bool                  `yaml:"has_pair"`
	GreatestItem    bool                  `yaml:"greatest_item"`
	LeastItem       bool                  `yaml:"least_item"`
	OutstandingItem bool                  `yaml:"outstanding_item"`
	HasOrder        bool                  `yaml:"has_order"`
	HasSequence     string                `yaml:"has_sequence"`
}

type yamlValidatorCompare struct {
	Item    string   `yaml:"item"`
	Sum     []string `yaml:"sum"`
	Target  string   `yaml:"target"`
	AnyPair bool     `yaml:"any_pair"`
	AnyItem bool     `yaml:"any_item"`
	Compare string   `yaml:"compare"`
}

type yamlValidatorCount struct {
	Number int    `yaml:"number"`
	Parity string `yaml:"parity"`
	OneOf  []int  `yaml:"one_of"`
}

type yamlValidatorParity struct {
	Item string `yaml:"item"`
	Sum  bool   `yaml:"sum"`
}
