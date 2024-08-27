package repository

import (
	"fmt"
	"new_ip_data_api/model"
)

const previousQuery = "INSERT INTO ip_data_endpoints (as_number, city, country, countrycode, isp, lat, lon, org, query, region, regionname, status, timezone, zip, time_stamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)"

func (ipRepo *IpDataRepository) StoreIpData(ipDataInfo model.IpDataInfo) (int, error) {
	//Criar query para salvar o IP no banco de dados e retornar o ID para que seja usado no use case
	var id int
	query, err := ipRepo.connection.Prepare(previousQuery + " RETURNING id")
	if err != nil {
		fmt.Println("implementar log: ", err)
		return 0, err
	}

	err = query.QueryRow(
		ipDataInfo.As,
		ipDataInfo.City,
		ipDataInfo.Country,
		ipDataInfo.CountryCode,
		ipDataInfo.Isp,
		ipDataInfo.Lat,
		ipDataInfo.Lon,
		ipDataInfo.Org,
		ipDataInfo.Query,
		ipDataInfo.Region,
		ipDataInfo.RegionName,
		ipDataInfo.Status,
		ipDataInfo.Timezone,
		ipDataInfo.Zip,
		ipDataInfo.TimeStamp).Scan(&id)
	if err != nil {
		fmt.Println("implementar log: ", err)
		return 0, err
	}

	query.Close()
	return id, nil
}
