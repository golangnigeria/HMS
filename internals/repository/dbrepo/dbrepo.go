package dbrepo

import (
	"database/sql"

	"github.com/golangnigeria/kinicart/internals/config"
	"github.com/golangnigeria/kinicart/internals/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgreRepo 
func NewPostgreRepo(conn *sql.DB, a *config.AppConfig)repository.DatabaseRepo{
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}