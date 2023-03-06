package assignment01_test

import (
	"fmt"
	groupcompanies "tesfayprep/assignment_01/groupCompanies"
	"testing"
	"time"
)

func TestGroupCompaniesByWeek(t *testing.T) {
	now := time.Now().UTC()
	companies := []groupcompanies.Company{
		{Department: "D2", CreatedAt: now.AddDate(0, 7, 0)},
		{Department: "D3", CreatedAt: now.AddDate(0, 7, 0)},
		{Department: "D4", CreatedAt: now.AddDate(0, 4, 0)},
		{Department: "D5", CreatedAt: now.AddDate(0, 2, 0)},
		{Department: "D6", CreatedAt: now.AddDate(0, 9, 0)},
		{Department: "D7", CreatedAt: now.AddDate(0, 10, 0)},
	}
	result := groupcompanies.GroupCompaniesByMonth(companies)
	fmt.Println(result)
	expectedResult := map[string][]groupcompanies.Company{

		"September": {
			{Department: "D2", CreatedAt: now.AddDate(0, 7, 0)},
			{Department: "D3", CreatedAt: now.AddDate(0, 7, 0)}},
		"June": {
			{Department: "D4", CreatedAt: now.AddDate(0, 4, 0)},
		},
		"April": {
			{Department: "D5", CreatedAt: now.AddDate(0, 2, 0)},
		},
		"November": {
			{Department: "D6", CreatedAt: now.AddDate(0, 9, 0)},
		},
		"December": {
			{Department: "D7", CreatedAt: now.AddDate(0, 10, 0)},
		},
	}
	if len(result) != len(expectedResult) {
		t.Fatalf("Expected result length to be %d but got %d", len(expectedResult), len(result))
	}
	for week, companies := range result {
		expectedCompanies, ok := expectedResult[week]
		if !ok {
			t.Fatalf("Unexpected week %v found in result", week)
		}
		if len(companies) != len(expectedCompanies) {
			t.Fatalf("Expected company list length to be %v but got %v for week %v", len(expectedCompanies), len(companies), week)
		}

	}
}
