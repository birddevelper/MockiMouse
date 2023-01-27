package config

type Config struct {
	MockServer struct {
		ContextPath string     `yaml:"contextPath,omitempty"`
		Port        int        `yaml:"port"`
		Endpoints   []EndPoint `yaml:"endpoints"`
	} `yaml:"MockServer"`
}

type EndPoint struct {
	Name        string     `yaml:"name"`
	Path        string     `yaml:"path"`
	ContentType string     `yaml:"contentType"`
	Consume     string     `yaml:"consume"`
	Method      string     `yaml:"method"`
	Delay       int        `yaml:"delay,omitempty"`
	Scenarios   []Scenario `yaml:"scenarios"`
}

type Scenario struct {
	Description string `yaml:"description"`
	Condition   struct {
		Params []Param `yaml:"param"`
	} `yaml:"condition,omitempty"`
	Response string `yaml:"response"`
	Status   int    `yaml:"status"`
}

type Param struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`    // header/body/query/form/path
	Operand string `yaml:"operand"` // eq/gt/lt/geq/leq/con
	Value   string `yaml:"value"`
}
