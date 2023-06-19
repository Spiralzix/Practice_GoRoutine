package service

import (
	"testing"

	"github.com/Spiralzix/LinemanAssignment/internal/core/service/mockdata"
)

func TestSummarizeCovidData(t *testing.T) {
	// Data preparation
	expectedSummary := CovidReport{
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
	mockCovidHistoricalData := &mockdata.MockCovidHistoricalData{}
	newCOVIDService := NewCOVIDService(mockCovidHistoricalData)

	// act
	result, _ := newCOVIDService.GetReport()

	// Compare the actual and expected results
	// condition 1: Check length of Actual Number of Province compare to Expected result
	if len(result.Province) != len(expectedSummary.Province) {
		t.Errorf("Unexpected province count. Expected: %d, Got: %d", len(expectedSummary.Province), len(result.Province))
	}

	// condition 2: Check length of Actual Number of AgeGroup compare to Expected result
	if len(result.AgeGroup) != len(expectedSummary.AgeGroup) {
		t.Errorf("Unexpected AgeGroup count. Expected: %d, Got: %d", len(expectedSummary.AgeGroup), len(result.AgeGroup))
	}

	// condition 3: Check the Result of each province and compare to Expected result
	for province, count := range expectedSummary.Province {
		if result.Province[province] != count {
			t.Errorf("Unexpected province count for %s. Expected: %d, Got: %d", province, count, result.Province[province])
		}
	}

	// condition 4: Check the Result of each AgeGroup and compare to Expected result
	for AgeGroup, count := range expectedSummary.AgeGroup {
		if result.AgeGroup[AgeGroup] != count {
			t.Errorf("Unexpected AgeGroup count for %s. Expected: %d, Got: %d", AgeGroup, count, result.AgeGroup[AgeGroup])
		}
	}
}
