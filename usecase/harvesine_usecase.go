package usecase

import (
	l "new_ip_data_api/config/logger"
	"new_ip_data_api/model"

	"go.uber.org/zap"
)

func (ipUseCase *IpDataUsecase) DistanciaPcaSe(ipOne, ipTwo string) (int, string, []model.IpDataInfo, error) {

	status, message, ipList, err := ipUseCase.repository.DistanciaPcaSe(ipOne, ipTwo)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		return status, message, nil, err
	}

	l.Logger.Warn(message, zap.Int("status", status))
	return status, message, ipList, err
}
