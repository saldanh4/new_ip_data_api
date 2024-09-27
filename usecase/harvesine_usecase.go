package usecase

import (
	"math"
	l "new_ip_data_api/config/logger"

	"go.uber.org/zap"
)

const (
	latSe float64 = -23.5505
	lonSe float64 = -46.6333
)

type IpDt struct {
	Ip       string
	Lat      float64
	Lon      float64
	Distance float64
}

func (ipUseCase *IpDataUsecase) DistanciaPcaSe(ipOne, ipTwo string) (int, string, []IpDt, error) {

	var ipResult []IpDt

	status, message, ipList, err := ipUseCase.repository.DistanciaPcaSe(ipOne, ipTwo)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		return status, message, nil, err
	}

	var ipDt []IpDt
	var ip string
	var la, lo, result float64

	for _, data := range ipList {
		ip = data.Query
		la = float64(data.Lat)
		lo = float64(data.Lon)
		//ipDt = []IpDt{{Ip: ip, Lat: la, Lon: lo}}
		ipDt = append(ipDt, IpDt{Ip: ip, Lat: la, Lon: lo})
	}

	for i := range ipDt {
		result = Haversine(latSe, lonSe, ipDt[i].Lat, ipDt[i].Lon)
		ipDt[i].Distance = result
	}

	a := ipDt[0].Distance
	b := ipDt[1].Distance

	if a < b {
		ipResult = append(ipResult, ipDt[0])
	} else if b < a {
		ipResult = append(ipResult, ipDt[1])
	}

	l.Logger.Warn(message, zap.Int("status", status))
	return status, message, ipResult, err
}

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Raio da Terra em quilômetros

	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLon := (lon2 - lon1) * (math.Pi / 180.0)

	lat1 = lat1 * (math.Pi / 180.0)
	lat2 = lat2 * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
