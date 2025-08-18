package database

import (
	"database/sql"
)

/*
	TODO: Database is set up via docker-compose.
	TODO: Volume should be moved to separated directory.
	TODO: Migrations are written with golang-migrate and should be run from golang code.

	Here should be connection or interface returned. Better interface to
	replace connection with pool if needed. Maybe move interface to repos.
*/

func NewConn() sql.Conn {
	panic("not implemented")
}
