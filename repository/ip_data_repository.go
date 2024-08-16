package repository

import (
	"database/sql"
)

const previousQuery = "INSERT INTO ip_data_endpoints (as_number, city, country, countrycode, isp, lat, lon, org, query, region, regionname, status, timezone, zip, time_stamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)"

type IpDataRepository struct {
	connection *sql.DB
}

func NewIpDataRepository(connection *sql.DB) IpDataRepository {
	return IpDataRepository{
		connection: connection,
	}
}
