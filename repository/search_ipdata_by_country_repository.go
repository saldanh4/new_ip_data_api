package repository

import (
	"errors"
	"net/http"
	l "new_ip_data_api/config/logger"
	"new_ip_data_api/model"

	"go.uber.org/zap"
)

func (ipRepo *IpDataRepository) GetTotalSearchByCountry(givenCountry string) (int, string, []model.IpDataInfo, error) {
	message = "Erro interno."
	err := ipRepo.connection.QueryRow(SELECT_COUNTRY_EXISTS_QUERY, givenCountry).Scan(&exists)
	if err != nil {

		l.Logger.Error(message, zap.Error(err))
		return http.StatusInternalServerError, message, []model.IpDataInfo{}, err
	}

	if !exists {
		message = "Não foram localizados dados para " + givenCountry + " em nosso banco de dados."
		err := errors.New(message)
		l.Logger.Warn(message, zap.Int("status", http.StatusNotFound))
		return http.StatusNotFound, message, []model.IpDataInfo{}, err
	}

	query, err := ipRepo.connection.Query(SEARCH_COUNTRY_QUERY, givenCountry)
	if err != nil {

		l.Logger.Error(message, zap.Error(err))
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

			l.Logger.Error(message, zap.Error(err))
			return http.StatusInternalServerError, message, []model.IpDataInfo{}, err
		}
		countryList = append(countryList, ipData)
	}
	message = "Consulta realizada com sucesso"
	l.Logger.Info(message, zap.Int("status", http.StatusOK))
	query.Close()
	return http.StatusOK, message, countryList, nil
}
