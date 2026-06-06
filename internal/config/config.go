package config

type Configuration struct {
	Http     Http     `mapstructure:"http"`
	Cache    Cache    `mapstructure:"cache"`
	Defaults Defaults `mapstructure:"defaults"`
	Timers   []Timer  `mapstructure:"timers"`
}
