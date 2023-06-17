package service_test

// import (
// 	"testing"
// )

// func TestSummarizeCovidData(t *testing.T) {
// 	// Test case data
// 	covidData := []Case{
// 		{Province: "Samut Sakhon", Age: 25},
// 		{Province: "Bangkok", Age: 40},
// 		{Province: "Samut Sakhon", Age: 70},
// 		{Province: "Bangkok", Age: 20},
// 		{Province: "Chiang Mai", Age: 45},
// 		{Province: "Chiang Mai", Age: 55},
// 		{Province: "Samut Sakhon", Age: -1}, // Age less than 0 should count as N/A
// 		{Province: "Bangkok", Age: 80},
// 	}

// 	expectedSummary := Summary{
// 		ProvinceCount: map[string]int{
// 			"Samut Sakhon": 3,
// 			"Bangkok":      3,
// 			"Chiang Mai":   2,
// 		},
// 		AgeGroupCount: map[string]int{
// 			"0-30":  2,
// 			"31-60": 3,
// 			"61+":   2,
// 			"N/A":   1,
// 		},
// 	}

// 	// Call the function being tested
// 	summary := summarizeCovidData(covidData)

// 	// Compare the actual and expected results
// 	if len(summary.ProvinceCount) != len(expectedSummary.ProvinceCount) {
// 		t.Errorf("Unexpected province count. Expected: %d, Got: %d", len(expectedSummary.ProvinceCount), len(summary.ProvinceCount))
// 	}

// 	if len(summary.AgeGroupCount) != len(expectedSummary.AgeGroupCount) {
// 		t.Errorf("Unexpected age group count. Expected: %d, Got: %d", len(expectedSummary.AgeGroupCount), len(summary.AgeGroupCount))
// 	}

// 	for province, count := range expectedSummary.ProvinceCount {
// 		if summary.ProvinceCount[province] != count {
// 			t.Errorf("Unexpected province count for %s. Expected: %d, Got: %d", province, count, summary.ProvinceCount[province])
// 		}
// 	}

// 	for ageGroup, count := range expectedSummary.AgeGroupCount {
// 		if summary.AgeGroupCount[ageGroup] != count {
// 			t.Errorf("Unexpected age group count for %s. Expected: %d, Got: %d", ageGroup, count, summary.AgeGroupCount[ageGroup])
// 		}
// 	}
// }
