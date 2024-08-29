package repository

import (
	"database/sql"
	"net/http"
	"new_ip_data_api/model"
)

func (ipRepo *IpDataRepository) GetTotalSearchByIP(ipNumber string) (int, string, *model.IpDataInfo, error) {

	var ipData model.IpDataInfo

	query, err := ipRepo.connection.Prepare(SEARCH_BY_IP_QUERY)
	if err != nil {
		message = "Erro ao efetuar consulta no banco de dados."
		return http.StatusInternalServerError, message, &model.IpDataInfo{}, err
	}

	if err := query.QueryRow(ipNumber).Scan(
		&ipData.Query,
		&ipData.Isp,
		&ipData.Country,
		&ipData.Count); err != nil {
		if err == sql.ErrNoRows {
			message = "IP " + ipNumber + "não localizado no banco de dados!"
			return http.StatusNotFound, message, nil, err
		}
		message = "Erro ao efetuar consulta no banco de dados."
		return http.StatusInternalServerError, message, nil, err
	}
	message = "Consulta realizada com sucesso"
	query.Close()
	return http.StatusOK, message, &ipData, err
}
