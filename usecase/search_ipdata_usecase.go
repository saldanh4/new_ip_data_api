package usecase

import (
	l "new_ip_data_api/config/logger"
	"new_ip_data_api/model"

	"go.uber.org/zap"
)

func (ipUseCase *IpDataUsecase) GetTotalSearchByIP(ipNumber string) (int, string, *model.IpDataInfo, error) {

	status, message, ipData, err := ipUseCase.repository.GetTotalSearchByIP(ipNumber)
	if err != nil {
		l.Logger.Info(message, zap.Int("status", status))
		return status, message, nil, err
	}
	l.Logger.Info(message, zap.Int("status", status))
	return status, message, ipData, err
}
