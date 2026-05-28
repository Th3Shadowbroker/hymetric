package config

type Configuration struct {
	Http   Http    `mapstructure:"http"`
	Cache  Cache   `mapstructure:"cache"`
	Timers []Timer `mapstructure:"timers"`
}
