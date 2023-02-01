package assignment01_test

import (
	"assignment/assignment_01/group"
	"testing"
)

func TestGroupUsersByWeak(t *testing.T) {
	testCases := []struct {
		Name  string
		Input []group.User
		Want  []group.CreatedWeek
	}{
		{
			Name: "profit of one group with one department",
			Input: []group.User{
				{
					ID: "1",
				},
			},
			Want: []group.CreatedWeek{
				{},
			},
		},
		{
			Name: "profit of one group with multiple department",
			Input: []group.User{
				{
					ID: "1",
				},
				{
					ID: "1",
				},
				{
					ID: "1",
				},
			},
			Want: []group.CreatedWeek{
				{},
			},
		},
		{
			Name: "profit of multiple group's with multiple department",
			Input: []group.User{
				{
					ID: "1",
				},
				{
					ID: "1",
				},
				{
					ID: "1",
				},
				{
					ID: "2",
				},
				{
					ID: "2",
				},
				{
					ID: "2",
				},
				{
					ID: "3",
				},
				{
					ID: "3",
				},
				{
					ID: "3",
				},
			},
			Want: []group.CreatedWeek{
				{},
				{},
				{},
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			got := group.GroupUsersByWeek(test.Input)
			var count int
			for _, g := range got {
				for _, w := range test.Want {
					if g.WeekNumber == w.WeekNumber {
						count++
					}
				}

			}
			if len(test.Want) != count {
				t.Errorf("got %v less result than expected %v", got, test.Want)
			}
		})

	}
}
