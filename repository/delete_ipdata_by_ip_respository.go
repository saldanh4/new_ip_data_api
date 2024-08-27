package repository

import (
	"fmt"
)

const deleteIpQuery = "DELETE FROM ip_data_endpoints WHERE query = $1"

func (ipRepo *IpDataRepository) DeleteIpDataByIp(ipNumber string) (string, error) {
	//Criar query para salvar o IP no banco de dados e retornar o ID para que seja usado no use case
	var message string

	query, err := ipRepo.connection.Prepare(deleteIpQuery)
	// var id int
	//
	fmt.Printf("query: %v\n", query)
	if err != nil {
		fmt.Println("implementar log: ", err)
		return "deu ruim", err
	}

	err = query.QueryRow(ipNumber)

	fmt.Printf("query: %v\n", query)
	message = "Apagado os registros para o IP " + ipNumber

	query.Close()
	return message, nil
}
