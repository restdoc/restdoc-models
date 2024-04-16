package Config

type ModelConfig struct {
	Debug      bool
	Timeout    int
	SaaSDomain string
	SqlDB      string
}

var DefaultConfig Config
