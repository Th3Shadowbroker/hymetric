package config

type Timer struct {
	Name        string `mapstructure:"name"`
	DisplayName string `mapstructure:"displayName"`
	Icon        string `mapstructure:"icon"`
	Url         string `mapstructure:"url"`
}
