package dbrepo

import (
	"database/sql"

	"github.com/golangnigeria/kinicart/internals/config"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

