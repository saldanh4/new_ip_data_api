package usecase

func (ipUseCase *IpDataUsecase) DeleteIpDataByIp(ipNumber string) (int, string, error) {
	status, message, err := ipUseCase.repository.DeleteIpDataByIp(ipNumber)
	if err != nil {
		return status, message, err
	}
	return status, message, nil

}
