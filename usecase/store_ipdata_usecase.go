package usecase

import "new_ip_data_api/model"

func (ipUseCase *IpDataUsecase) StoreIpData(ipData model.IpDataInfo) (int, string, model.IpDataInfo, error) {
	//pegar o ID gerado ao salvar o IP no banco e retornar em um tipo de Ipdata para exibição
	status, message, ipDataID, err := ipUseCase.repository.StoreIpData(ipData)
	if err != nil {
		return status, message, model.IpDataInfo{}, err
	}

	ipData.Id = ipDataID

	return status, message, ipData, nil
}
