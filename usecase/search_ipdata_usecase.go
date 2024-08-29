package usecase

import "new_ip_data_api/model"

func (ipUseCase *IpDataUsecase) GetTotalSearchByIP(ipNumber string) (int, string, *model.IpDataInfo, error) {

	status, message, ipData, err := ipUseCase.repository.GetTotalSearchByIP(ipNumber)
	if err != nil {
		return status, message, nil, err
	}

	return status, message, ipData, err
}
