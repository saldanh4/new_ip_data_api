package usecase

import (
	l "new_ip_data_api/config/logger"
	"new_ip_data_api/model"

	"go.uber.org/zap"
)

func (ipUseCase *IpDataUsecase) StoreIpData(ipData model.IpDataInfo) (int, string, model.IpDataInfo, error) {
	//pegar o ID gerado ao salvar o IP no banco e retornar em um tipo de Ipdata para exibição
	status, message, ipDataID, err := ipUseCase.repository.StoreIpData(ipData)
	if err != nil {
		l.Logger.Info(message, zap.Int("status", status))
		return status, message, model.IpDataInfo{}, err
	}

	l.Logger.Info(message, zap.Int("status", status))
	ipData.Id = ipDataID

	return status, message, ipData, nil
}
