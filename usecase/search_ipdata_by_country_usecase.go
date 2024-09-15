package usecase

import (
	l "new_ip_data_api/config/logger"
	"new_ip_data_api/model"

	"go.uber.org/zap"
)

func (ipUseCase *IpDataUsecase) GetTotalSearchByCountry(givenCountry string) (int, string, []model.IpDataInfo, error) {

	status, message, countryList, err := ipUseCase.repository.GetTotalSearchByCountry(givenCountry)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		return status, message, nil, err
	}

	l.Logger.Warn(message, zap.Int("status", status))
	return status, message, countryList, err
}
