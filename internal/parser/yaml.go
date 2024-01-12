package parser

type yamlRoot struct {
	Validators map[string]yamlValidator `yaml:"validators"`
	Tests      map[int]map[string]bool  `yaml:"tests"`
}

type yamlValidator struct {
	Compare        *yamlValidatorCompare `yaml:"compare"`
	Count          *yamlValidatorCount   `yaml:"count"`
	Parity         *yamlValidatorParity  `yaml:"parity"`
	HasMoreParity  bool                  `yaml:"has_more_parity"`
	HasRepetitions bool                  `yaml:"has_repetitions"`
	HasPair        bool                  `yaml:"has_pair"`
	GreatestItem   bool                  `yaml:"greatest_item"`
	LeastItem      bool                  `yaml:"least_item"`
	OutlierItem    bool                  `yaml:"outlier_item"`
	HasOrder       bool                  `yaml:"has_order"`
}

type yamlValidatorCompare struct {
	Item   string   `yaml:"item"`
	Sum    []string `yaml:"sum"`
	Target string   `yaml:"target"`
	Multi  bool     `yaml:"multi"`
}

type yamlValidatorCount struct {
	Number       int    `yaml:"number"`
	Parity       string `yaml:"parity"`
	OneOfNumbers []int  `yaml:"one_of_numbers"`
}

type yamlValidatorParity struct {
	Item string `yaml:"item"`
	Sum  bool   `yaml:"sum"`
}
