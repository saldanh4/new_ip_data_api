package repository

import (
	"database/sql"
)

type IpDataRepository struct {
	connection *sql.DB
}

func NewIpDataRepository(connection *sql.DB) IpDataRepository {
	return IpDataRepository{
		connection: connection,
	}
}
