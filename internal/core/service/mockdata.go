package service

import "github.com/Spiralzix/LinemanAssignment/internal/core/repositorys"

type CovidHistoricalData struct {
	Data []CovidData `json:"Data"`
}

type CovidData struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             int    `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     int    `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine int    `json:"StatQuarantine"`
}

func (c *CovidHistoricalData) FetchData() (*repositorys.CovidHistoricalData, error) {
	data := repositorys.CovidHistoricalData{
		Data: []repositorys.CovidData{
			{
				ConfirmDate:    "2022-06-15",
				No:             1,
				Age:            31,
				Gender:         "ชาย",
				GenderEn:       "Male",
				Nation:         "ไทย",
				NationEn:       "Thai",
				Province:       "Bangkok",
				ProvinceId:     1,
				District:       "ลาดกระบัง",
				ProvinceEn:     "Bangkok",
				StatQuarantine: 0,
			},
			{
				ConfirmDate:    "2022-06-16",
				No:             2,
				Age:            10,
				Gender:         "หญิง",
				GenderEn:       "Female",
				Nation:         "สหรัฐอเมริกา",
				NationEn:       "United States",
				Province:       "Bangkok",
				ProvinceId:     2,
				District:       "ลาดกระบัง",
				ProvinceEn:     "Bangkok",
				StatQuarantine: 1,
			},
			{
				ConfirmDate:    "2022-06-16",
				No:             2,
				Age:            77,
				Gender:         "หญิง",
				GenderEn:       "Female",
				Nation:         "สหรัฐอเมริกา",
				NationEn:       "United States",
				Province:       "Nakhon Pathom",
				ProvinceId:     2,
				District:       "ลาดกระบัง",
				ProvinceEn:     "Nakhon Pathom",
				StatQuarantine: 1,
			},
		},
	}
	return &data, nil
}
