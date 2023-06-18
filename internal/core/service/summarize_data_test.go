package service

import (
	"testing"
)

func TestSummarizeCovidData(t *testing.T) {

	// Data preparation
	expectedSummary := Response{
		Province: map[string]int{
			"Bangkok":       2,
			"Nakhon Pathom": 1,
		},
		AgeGroup: map[string]int{
			"0-30":  1,
			"31-60": 1,
			"61+":   1,
			"N/A":   0,
		},
	}

	// arrange
	mockCovidHistoricalData := &CovidHistoricalData{}
	newCOVIDService := NewCOVIDService(mockCovidHistoricalData)

	// act
	result, _ := newCOVIDService.GetReport()

	// Compare the actual and expected results
	if len(result.Province) != len(expectedSummary.Province) {
		t.Errorf("Unexpected province count. Expected: %d, Got: %d", len(expectedSummary.Province), len(result.Province))
	}
	if len(result.AgeGroup) != len(expectedSummary.AgeGroup) {
		t.Errorf("Unexpected AgeGroup count. Expected: %d, Got: %d", len(expectedSummary.AgeGroup), len(result.AgeGroup))
	}

	for province, count := range expectedSummary.Province {
		if result.Province[province] != count {
			t.Errorf("Unexpected province count for %s. Expected: %d, Got: %d", province, count, result.Province[province])
		}
	}
	for AgeGroup, count := range expectedSummary.AgeGroup {
		if result.AgeGroup[AgeGroup] != count {
			t.Errorf("Unexpected AgeGroup count for %s. Expected: %d, Got: %d", AgeGroup, count, result.AgeGroup[AgeGroup])
		}
	}
}
