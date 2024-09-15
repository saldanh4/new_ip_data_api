package usecase

import (
	l "new_ip_data_api/config/logger"

	"go.uber.org/zap"
)

func (ipUseCase *IpDataUsecase) DeleteIpDataByIp(ipNumber string) (int, string, error) {
	status, message, err := ipUseCase.repository.DeleteIpDataByIp(ipNumber)
	if err != nil {
		l.Logger.Info(message, zap.Int("status", status))
		return status, message, err
	}
	l.Logger.Info(message, zap.Int("status", status))
	return status, message, nil

}
