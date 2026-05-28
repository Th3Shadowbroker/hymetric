package config

import "time"

type Cache struct {
	Ttl             time.Duration `mapstructure:"ttl"`
	CleanupInterval time.Duration `mapstructure:"cleanupInterval"`
}
