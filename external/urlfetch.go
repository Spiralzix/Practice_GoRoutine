package external

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Spiralzix/LinemanAssignment/config"
)

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

type CovidRecord struct {
	web *config.Config
}

func NewCOVIDRecord(url *config.Config) ICovidRecord {
	return &CovidRecord{url}
}

type ICovidRecord interface {
	FetchData() (*CovidHistoricalData, error)
}

func (r *CovidRecord) FetchData() (*CovidHistoricalData, error) {
	resp, err := http.Get(r.web.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data CovidHistoricalData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
