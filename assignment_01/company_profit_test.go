package assignment01_test

import (
	"tesfayprep/assignment_01/company"
	"testing"
)

func TestCompanyProfits(t *testing.T) {
	testCases := []struct {
		Name  string
		Input []company.CompanyProfit
		Want  []company.CompanyTotalProfit
	}{
		{
			Name: "profit of one company with one department",
			Input: []company.CompanyProfit{
				{
					ID:         "1",
					Profit:     10,
					Department: "A",
				},
				{
					ID:         "1",
					Profit:     5,
					Department: "A",
				},
				{
					ID:         "2",
					Profit:     5,
					Department: "A",
				},
			},
			Want: []company.CompanyTotalProfit{
				{
					CompanyID:   "1",
					TotalProfit: 20,
					Department: []company.DepartmentProfit{
						{
							Department:       "A",
							DepartmentProfit: 15,
						},
						{
							Department:       "B",
							DepartmentProfit: 5,
						},
					},
				},
			},
		},
		{
			Name: "profit of one company with one department",
			Input: []company.CompanyProfit{
				{
					ID:         "1",
					Profit:     10,
					Department: "A",
				},
			},
			Want: []company.CompanyTotalProfit{
				{
					CompanyID:   "1",
					TotalProfit: 10,
					Department: []company.DepartmentProfit{
						{
							Department:       "A",
							DepartmentProfit: 10,
						},
					},
				},
			},
		},
		{
			Name: "profit of multiple company's with multiple department",
			Input: []company.CompanyProfit{
				{
					ID:         "1",
					Profit:     10,
					Department: "A",
				},
				{
					ID:         "1",
					Profit:     15,
					Department: "B",
				},
				{
					ID:         "1",
					Profit:     15,
					Department: "C",
				},
				{
					ID:         "2",
					Profit:     15,
					Department: "A",
				},
				{
					ID:         "2",
					Profit:     20,
					Department: "B",
				},
				{
					ID:         "2",
					Profit:     10,
					Department: "B",
				},
				{
					ID:         "2",
					Profit:     15,
					Department: "C",
				},
				{
					ID:         "3",
					Profit:     10,
					Department: "A",
				},
				{
					ID:         "3",
					Profit:     10,
					Department: "B",
				},
				{
					ID:         "3",
					Profit:     15,
					Department: "C",
				},
			},
			Want: []company.CompanyTotalProfit{
				{
					CompanyID:   "1",
					TotalProfit: 40,
					Department: []company.DepartmentProfit{
						{
							Department:       "A",
							DepartmentProfit: 10,
						},
						{
							Department:       "B",
							DepartmentProfit: 15,
						},
						{
							Department:       "C",
							DepartmentProfit: 15,
						},
					},
				},
				{
					CompanyID:   "2",
					TotalProfit: 60,
					Department: []company.DepartmentProfit{
						{
							Department:       "A",
							DepartmentProfit: 15,
						},
						{
							Department:       "B",
							DepartmentProfit: 30,
						},
						{
							Department:       "C",
							DepartmentProfit: 15,
						},
					},
				},
				{
					CompanyID:   "3",
					TotalProfit: 35,
					Department: []company.DepartmentProfit{
						{
							Department:       "A",
							DepartmentProfit: 10,
						},
						{
							Department:       "B",
							DepartmentProfit: 10,
						},
						{
							Department:       "C",
							DepartmentProfit: 15,
						},
					},
				},
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			got := company.CalculateProfits(test.Input)
			var count int
			var countTotal int

			for i, g := range got {

				if g.TotalProfit == test.Want[i].TotalProfit {
					countTotal++
					for j, d := range g.Department {
						if d.DepartmentProfit == test.Want[i].Department[j].DepartmentProfit {
							count++
						}
					}
				}

			}
			if len(test.Want) != count {
				t.Errorf("got %v less result than expected %v", got, test.Want)
			}
		})

	}
}
