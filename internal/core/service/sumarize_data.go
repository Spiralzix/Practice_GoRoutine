package service

import (
	"sync"

	"github.com/Spiralzix/LinemanAssignment/external"
)

type CovidReport struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}
type CovidService struct {
	record external.ICovidRecord
}

type ICovidService interface {
	GetReport() (*CovidReport, error)
}

func NewCOVIDService(record external.ICovidRecord) ICovidService {
	return &CovidService{
		record: record,
	}
}

// count case of covid by province
func countByProvince(covidData *external.CovidHistoricalData) map[string]int {
	var mapbyProvince = make(map[string]int)
	for _, value := range covidData.Data {
		mapbyProvince[value.Province] += 1
	}
	return mapbyProvince
}

// count case of covid by age
func countByAge(covidData *external.CovidHistoricalData) map[string]int {
	var mapbyAge = map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}
	for _, value := range covidData.Data {
		switch age := value.Age; {
		case age <= 30:
			mapbyAge["0-30"]++
		case age >= 31 && age <= 60:
			mapbyAge["31-60"]++
		case age >= 61:
			mapbyAge["61+"]++
		default:
			mapbyAge["N/A"]++
		}
	}
	return mapbyAge
}

func (s *CovidService) GetReport() (*CovidReport, error) {
	covidData, err := s.record.FetchData()
	if err != nil {
		return nil, err
	}
	var (
		mapByProvince map[string]int
		mapByAge      map[string]int
		wg            sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		mapByProvince = countByProvince(covidData)
	}()

	go func() {
		defer wg.Done()
		mapByAge = countByAge(covidData)
	}()

	wg.Wait()
	result := &CovidReport{
		Province: mapByProvince,
		AgeGroup: mapByAge,
	}
	return result, nil
}
