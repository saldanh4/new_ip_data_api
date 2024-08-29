package repository

import (
	"fmt"
	"net/http"
)

func (ipRepo *IpDataRepository) DeleteIpDataByIp(ipNumber string) (int, string, error) {

	err := ipRepo.connection.QueryRow(SELECT_IP_EXISTS_QUERY, ipNumber).Scan(&exists)
	if err != nil {
		message = "Erro ao efetuar consulta no banco de dados. Erro: "
		return http.StatusInternalServerError, message, err
	}

	if !exists {
		message = "IP: " + ipNumber + ". Não consta na base de dados!"
		return http.StatusNotFound, message, err
	}

	query, err := ipRepo.connection.Prepare(DELETE_IP_QUERY)

	if err != nil {
		fmt.Println("implementar log: ", err)
		message = "Erro ao efetuar consulta no banco de dados. Erro: "
		return http.StatusInternalServerError, message, err
	}

	err = query.QueryRow(ipNumber)

	message = "Todos os registros para o IP " + ipNumber + " foram apagados!"

	query.Close()
	return http.StatusOK, message, nil
}
