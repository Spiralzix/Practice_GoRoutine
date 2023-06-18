package service

import (
	"testing"

	"github.com/Spiralzix/LinemanAssignment/internal/core/repositorys"
)

func TestSummarizeCovidData(t *testing.T) {
	data := repositorys.CovidHistoricalData{
		Data: []repositorys.CovidData{
			{
				ConfirmDate:    "2022-06-15",
				No:             1,
				Age:            32,
				Gender:         "ชาย",
				GenderEn:       "Male",
				Nation:         "ไทย",
				NationEn:       "Thai",
				Province:       "กรุงเทพมหานคร",
				ProvinceId:     1,
				District:       "ลาดกระบัง",
				ProvinceEn:     "Bangkok",
				StatQuarantine: 0,
			},
			{
				ConfirmDate:    "2022-06-16",
				No:             2,
				Age:            45,
				Gender:         "หญิง",
				GenderEn:       "Female",
				Nation:         "สหรัฐอเมริกา",
				NationEn:       "United States",
				Province:       "สมุทรสาคร",
				ProvinceId:     2,
				District:       "บางคนที",
				ProvinceEn:     "Samut Sakhon",
				StatQuarantine: 1,
			},
		},
	}

	mockCovidHistoricalData := data
	newCOVIDService := NewCOVIDService(mockCovidHistoricalData)

	// expectedSummary := Response{
	// 	Province: map[string]int{
	// 		"Bangkok":       2,
	// 		"Nakhon Pathom": 2,
	// 		"Yala":          1,
	// 	},
	// 	AgeGroup: map[string]int{
	// 		"0-30":  2,
	// 		"31-60": 3,
	// 		"61+":   2,
	// 		"N/A":   1,
	// 	},
	// }
	// Call the function being tested
	// mockCovidHistoricalData := MockCovidHistoricalData{}
	// newCOVIDService := NewCOVIDService(mockCovidHistoricalData)

	// Arrange
	// result, _ := newCOVIDService.GetReport()

	// // Compare the actual and expected results
	// if len(result.Province) != len(expectedSummary.Province) {
	// 	t.Errorf("Unexpected province count. Expected: %d, Got: %d", len(expectedSummary.Province), len(result.Province))
	// }

	// if len(summary.AgeGroupCount) != len(expectedSummary.AgeGroupCount) {
	// 	t.Errorf("Unexpected age group count. Expected: %d, Got: %d", len(expectedSummary.AgeGroupCount), len(summary.AgeGroupCount))
	// }

	// for province, count := range expectedSummary.ProvinceCount {
	// 	if summary.ProvinceCount[province] != count {
	// 		t.Errorf("Unexpected province count for %s. Expected: %d, Got: %d", province, count, summary.ProvinceCount[province])
	// 	}
	// }

	// for ageGroup, count := range expectedSummary.AgeGroupCount {
	// 	if summary.AgeGroupCount[ageGroup] != count {
	// 		t.Errorf("Unexpected age group count for %s. Expected: %d, Got: %d", ageGroup, count, summary.AgeGroupCount[ageGroup])
	// 	}
	// }
}
