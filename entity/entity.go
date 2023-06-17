package entity

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

type Response struct {
	Response []struct {
		Province []string `json:"Province"`
		AgeGroup []string `json:"AgeGroup"`
	} `json:"response"`
}
