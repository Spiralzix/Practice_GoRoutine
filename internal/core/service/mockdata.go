package service

// type CovidHistoricalData struct {
// 	Data []MockCovidData `json:"Data"`
// }

// type MockCovidData struct {
// 	ConfirmDate    string `json:"ConfirmDate"`
// 	No             int    `json:"No"`
// 	Age            int    `json:"Age"`
// 	Gender         string `json:"Gender"`
// 	GenderEn       string `json:"GenderEn"`
// 	Nation         string `json:"Nation"`
// 	NationEn       string `json:"NationEn"`
// 	Province       string `json:"Province"`
// 	ProvinceId     int    `json:"ProvinceId"`
// 	District       string `json:"District"`
// 	ProvinceEn     string `json:"ProvinceEn"`
// 	StatQuarantine int    `json:"StatQuarantine"`
// }

// type covidRepo struct{}

// type ICovidRepo interface {
// 	FetchData() (*CovidHistoricalData, error)
// }

// func (r *covidRepo) FetchData() (*CovidHistoricalData, error) {
// 	data := CovidHistoricalData{
// 		Data: []MockCovidData{
// 			{
// 				ConfirmDate:    "2022-06-15",
// 				No:             1,
// 				Age:            32,
// 				Gender:         "ชาย",
// 				GenderEn:       "Male",
// 				Nation:         "ไทย",
// 				NationEn:       "Thai",
// 				Province:       "กรุงเทพมหานคร",
// 				ProvinceId:     1,
// 				District:       "ลาดกระบัง",
// 				ProvinceEn:     "Bangkok",
// 				StatQuarantine: 0,
// 			},
// 			{
// 				ConfirmDate:    "2022-06-16",
// 				No:             2,
// 				Age:            45,
// 				Gender:         "หญิง",
// 				GenderEn:       "Female",
// 				Nation:         "สหรัฐอเมริกา",
// 				NationEn:       "United States",
// 				Province:       "สมุทรสาคร",
// 				ProvinceId:     2,
// 				District:       "บางคนที",
// 				ProvinceEn:     "Samut Sakhon",
// 				StatQuarantine: 1,
// 			},
// 		},
// 	}
// 	return &data, nil
// }
