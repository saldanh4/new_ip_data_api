package repository

import (
	"errors"
	"fmt"
	"net/http"
	"new_ip_data_api/model"
)

func (ipRepo *IpDataRepository) GetTotalSearchByCountry(givenCountry string) (int, string, []model.IpDataInfo, error) {

	err := ipRepo.connection.QueryRow(SELECT_COUNTRY_EXISTS_QUERY, givenCountry).Scan(&exists)
	if err != nil {
		message = "Erro ao efetuar consulta no banco de dados."
		return http.StatusInternalServerError, message, []model.IpDataInfo{}, err
	}

	if !exists {
		message = "Não foram localizados dados para " + givenCountry + " em nosso banco de dados."
		err := errors.New(message)
		return http.StatusNotFound, message, []model.IpDataInfo{}, err
	}

	query, err := ipRepo.connection.Query(SEARCH_COUNTRY_QUERY, givenCountry)
	if err != nil {
		fmt.Println("Implementar log: ", err)
		message = "Erro ao efetuar consulta no banco de dados."
		return http.StatusInternalServerError, message, []model.IpDataInfo{}, err
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
			message = "Erro ao efetuar consulta no banco de dados."
			return http.StatusInternalServerError, message, []model.IpDataInfo{}, err
		}
		countryList = append(countryList, ipData)
	}
	message = "Consulta realizada com sucesso"
	query.Close()
	return http.StatusOK, message, countryList, nil
}
