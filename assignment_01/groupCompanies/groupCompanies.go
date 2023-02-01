package groupcompanies

import (
	"time"
)

type Company struct {
	Department string
	CreatedAt  time.Time
}

//	INSTRUCTION
//
// group the companies by the their respective month and store them in a map,
//
//	i.e companies created in the same month should belong together ex-
//
//	"October":[]Company{
//	 {
//	  Profit:     10,
//	  Department: "D1",
//	 }
//	}
//
//	"November" : []Company{
//	 {
//	  Profit:     10,
//	  Department: "D1",
//	 }
//	}
func GroupCompaniesByMonth(companies []Company) map[string][]Company {
	result := make(map[string][]Company)
	// TODO: Add your code here

	return result
}
