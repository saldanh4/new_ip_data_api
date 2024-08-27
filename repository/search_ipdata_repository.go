package repository

import (
	"database/sql"
	"fmt"
	"new_ip_data_api/model"
)

const searchIpQuery = "SELECT query, isp, country, COUNT(*) as qtd FROM ip_data_endpoints  WHERE query = $1 GROUP BY query, isp, country"

func (ipRepo *IpDataRepository) GetTotalSearchByIP(ipNumber string) (*model.IpDataInfo, error) {
	//var ipList []model.IpDataInfo
	var ipData model.IpDataInfo

	query, err := ipRepo.connection.Prepare(searchIpQuery)
	if err != nil {
		fmt.Println("Implementar log: ", err)
		return &model.IpDataInfo{}, err
	}

	err = query.QueryRow(ipNumber).Scan(
		&ipData.Query,
		&ipData.Isp,
		&ipData.Country,
		&ipData.Count,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &ipData, nil
}
