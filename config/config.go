package Config

type ModelConfig struct {
	Debug      bool
	Timeout    int
	SaaSDomain string
	Mysql      string
}

var DefaultConfig Config
