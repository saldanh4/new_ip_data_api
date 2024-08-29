package repository

import (
	"net/http"
	"new_ip_data_api/model"
)

func (ipRepo *IpDataRepository) StoreIpData(ipDataInfo model.IpDataInfo) (int, string, int, error) {

	var id int
	query, err := ipRepo.connection.Prepare(INSERT_IP_DATA_QUERY + " RETURNING id")
	if err != nil {

		message = "Erro ao executar query no banco de dados."
		return http.StatusInternalServerError, message, 0, err
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
		message = "Erro ao executar query no banco de dados."
		return http.StatusInternalServerError, message, 0, err
	}
	message = "Cadastro realizado com sucesso"
	query.Close()
	return http.StatusOK, message, id, nil
}
