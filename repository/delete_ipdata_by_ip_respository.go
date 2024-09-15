package repository

import (
	"net/http"
	l "new_ip_data_api/config/logger"

	"go.uber.org/zap"
)

func (ipRepo *IpDataRepository) DeleteIpDataByIp(ipNumber string) (int, string, error) {

	err := ipRepo.connection.QueryRow(SELECT_IP_EXISTS_QUERY, ipNumber).Scan(&exists)
	if err != nil {
		message = "Erro ao efetuar consulta no banco de dados."
		l.Logger.Error(message, zap.Error(err))
		return http.StatusInternalServerError, message, err
	}

	if !exists {
		message = "IP: " + ipNumber + ". Não consta na base de dados!"
		l.Logger.Warn(message, zap.Int("status", http.StatusNotFound))
		return http.StatusNotFound, message, err
	}

	query, err := ipRepo.connection.Prepare(DELETE_IP_QUERY)

	if err != nil {
		message = "Erro ao efetuar consulta no banco de dados."
		l.Logger.Error(message, zap.Error(err))
		return http.StatusInternalServerError, message, err
	}

	err = query.QueryRow(ipNumber)

	message = "Todos os registros para o IP " + ipNumber + " foram apagados!"
	l.Logger.Info(message, zap.Int("status", http.StatusOK))
	query.Close()
	return http.StatusOK, message, nil
}
