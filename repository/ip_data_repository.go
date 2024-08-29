package repository

import (
	"database/sql"
)

const (
	DELETE_IP_QUERY      = "DELETE FROM ip_data_endpoints WHERE query = $1"
	SEARCH_BY_IP_QUERY   = "SELECT query, isp, country, COUNT(*) as qtd FROM ip_data_endpoints  WHERE query = $1 GROUP BY query, isp, country"
	SEARCH_COUNTRY_QUERY = "SELECT * FROM ip_data_endpoints  WHERE country = $1 GROUP BY id order by isp"

	SELECT_COUNTRY_EXISTS_QUERY = "SELECT EXISTS(SELECT * FROM ip_data_endpoints WHERE country = $1)"
	SELECT_IP_EXISTS_QUERY      = "SELECT EXISTS(SELECT * FROM ip_data_endpoints WHERE query = $1)"

	INSERT_IP_DATA_QUERY = "INSERT INTO ip_data_endpoints (as_number, city, country, countrycode, isp, lat, lon, org, query, region, regionname, status, timezone, zip, time_stamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)"
)

var (
	message string
	exists  bool
)

type IpDataRepository struct {
	connection *sql.DB
}

func NewIpDataRepository(connection *sql.DB) IpDataRepository {
	return IpDataRepository{
		connection: connection,
	}
}
