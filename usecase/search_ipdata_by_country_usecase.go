package usecase

import "new_ip_data_api/model"

func (ipUseCase *IpDataUsecase) GetTotalSearchByCountry(givenCountry string) (int, string, []model.IpDataInfo, error) {

	status, message, countryList, err := ipUseCase.repository.GetTotalSearchByCountry(givenCountry)
	if err != nil {
		return status, message, nil, err
	}

	return status, message, countryList, err
}
