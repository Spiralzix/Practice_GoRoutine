package service

import (
	"sync"

	"github.com/Spiralzix/LinemanAssignment/internal/core/repositorys"
)

// type Response struct {
// 	Response []struct {
// 		Province []string `json:"Province"`
// 		AgeGroup []string `json:"AgeGroup"`
// 	} `json:"response"`
// }

type Response struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}
type CovidService struct {
	repo repositorys.ICovidRepo
}

type ICovidService interface {
	GetReport() (*Response, error)
}

func NewCOVIDService(repo repositorys.ICovidRepo) ICovidService {
	return &CovidService{
		repo: repo,
	}
}

func countByProvince(covidData *repositorys.CovidHistoricalData) map[string]int {
	var mapbyProvince = make(map[string]int)
	for _, value := range covidData.Data {
		mapbyProvince[value.Province] += 1
	}
	return mapbyProvince
}

func countByAge(covidData *repositorys.CovidHistoricalData) map[string]int {
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

func (s *CovidService) GetReport() (*Response, error) {
	covidData, err := s.repo.FetchData()
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

	response := &Response{
		Province: mapByProvince,
		AgeGroup: mapByAge,
	}
	return response, nil
}
