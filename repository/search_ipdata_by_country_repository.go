package repository

import (
	"database/sql"
	"fmt"
	"new_ip_data_api/model"
)

const searchCountryQuery = "SELECT * FROM ip_data_endpoints  WHERE country = $1 GROUP BY id order by isp"

func (ipRepo *IpDataRepository) GetTotalSearchByCountry(givenCountry string) ([]model.IpDataInfo, error) {

	query, err := ipRepo.connection.Query(searchCountryQuery, givenCountry)
	if err != nil {
		fmt.Println("Implementar log: ", err)
		return []model.IpDataInfo{}, err
	}

	var countryList []model.IpDataInfo
	var ipData model.IpDataInfo

	for query.Next() {
		err = query.Scan(
			&ipData.Id,
			&ipData.As,
			&ipData.City,
			&ipData.Country,
			&ipData.CountryCode,
			&ipData.Isp,
			&ipData.Lat,
			&ipData.Lon,
			&ipData.Org,
			&ipData.Query,
			&ipData.Region,
			&ipData.RegionName,
			&ipData.Status,
			&ipData.Timezone,
			&ipData.Zip,
			&ipData.TimeStamp)

		if err != nil {
			fmt.Println(err)
			return []model.IpDataInfo{}, err
		}
		countryList = append(countryList, ipData)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return countryList, nil
}
