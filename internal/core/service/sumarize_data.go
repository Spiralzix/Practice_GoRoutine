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

func (s *CovidService) GetReport() (*Response, error) {
	covidData, err := s.repo.FetchData()
	if err != nil {
		return nil, err
	}
	var mapbyProvince = make(map[string]int)
	var mapbyAge = map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}

	var wg sync.WaitGroup
	for _, value := range covidData.Data {
		wg.Add(1)
		// go func(value Response) {
		defer wg.Done()
		mapbyProvince[value.Province] += 1
		switch age := value.Age; {
		case age <= 30:
			mapbyAge["0-30"] += 1
		case age == 31 && age <= 60:
			mapbyAge["31-60"] += 1
		case age >= 61:
			mapbyAge["61+"] += 1
		default:
			mapbyAge["N/A"] += 1
		}
		// }(value)
	}

	response := &Response{
		Province: mapbyProvince,
		AgeGroup: mapbyAge,
	}
	return response, nil
}
