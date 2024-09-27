package repository

import (
	"errors"
	"net/http"
	l "new_ip_data_api/config/logger"
	"new_ip_data_api/model"

	"go.uber.org/zap"
)

var query string = `
SELECT 
	EXISTS(SELECT 1 FROM ip_data_endpoints WHERE query = $1) AS ip1_exists,
	EXISTS(SELECT 1 FROM ip_data_endpoints WHERE query = $2) AS ip2_exists
`

func (ipRepo *IpDataRepository) DistanciaPcaSe(ipOne, ipTwo string) (int, string, []model.IpDataInfo, error) {
	message = "Erro interno"
	var ipList []model.IpDataInfo
	var ipData model.IpDataInfo

	var ip1_exists, ip2_exists bool

	err := ipRepo.connection.QueryRow(query, ipOne, ipTwo).Scan(&ip1_exists, &ip2_exists)
	if err != nil {
		l.Logger.Error(message, zap.Error(err))
		return http.StatusInternalServerError, message, []model.IpDataInfo{}, err
	}

	if !ip1_exists && !ip2_exists {
		message = "Não foram localizados dados para os IPs informados em nosso banco de dados."
		err := errors.New(message)
		l.Logger.Warn(message, zap.Int("status", http.StatusNotFound))
		return http.StatusNotFound, message, []model.IpDataInfo{}, err
	} else if !ip1_exists {
		message = "Não foram localizados dados para o IP " + ipOne + " em nosso banco de dados."
		err := errors.New(message)
		l.Logger.Warn(message, zap.Int("status", http.StatusNotFound))
		return http.StatusNotFound, message, []model.IpDataInfo{}, err
	} else if !ip2_exists {
		message = "Não foram localizados dados para o IP " + ipTwo + " em nosso banco de dados."
		err := errors.New(message)
		l.Logger.Warn(message, zap.Int("status", http.StatusNotFound))
		return http.StatusNotFound, message, []model.IpDataInfo{}, err
	}

	query, err := ipRepo.connection.Query("select distinct query, lat, lon FROM ip_data_endpoints WHERE query = $1 or query = $2 order by query", ipOne, ipTwo)
	if err != nil {
		message = "Erro ao efetuar consulta no banco de dados."
		l.Logger.Error(message, zap.Error(err))
		return http.StatusInternalServerError, message, []model.IpDataInfo{}, err
	}

	for query.Next() {
		err = query.Scan(
			&ipData.Query,
			&ipData.Lat,
			&ipData.Lon)
		if err != nil {
			message = "Erro interno."
			l.Logger.Error(message, zap.Error(err))
			return http.StatusInternalServerError, message, []model.IpDataInfo{}, err
		}
		ipList = append(ipList, ipData)
	}

	message = "Consulta realizada com sucesso"
	l.Logger.Info(message, zap.Int("status", http.StatusOK))
	query.Close()
	return http.StatusOK, message, ipList, nil

}
