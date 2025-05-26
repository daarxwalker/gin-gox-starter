package database_config

import (
	"github.com/spf13/viper"
	
	"common/pkg/env"
)

const (
	Uri = "database-uri"
)

func Read(v *viper.Viper) {
	if env.Development() {
		v.Set(Uri, "postgresql://starter:starter@starter-postgres:5432/starter")
	}
}
