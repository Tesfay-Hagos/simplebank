package company

type CompanyProfit struct {
	ID         string
	Profit     float64
	Department string
}

type DepartmentProfit struct {
	Department       string
	DepartmentProfit float64
}
type CompanyTotalProfit struct {
	CompanyID   string
	TotalProfit float64
	Department  []DepartmentProfit
}

// INSTRUCTION:
//
//	calculate total profit of a given company and their department
//	Given that you have a list of companies and their profit at a given time ("CompanyProfit"),
//	create a function that calculates the total profit of each companies department and sum the departments profit as the companies total profit.
//	note :- you could have multiple profit for one department of a given company
//
// Ex: Given this
//
//	{
//		ID:         "1",
//		Profit:     10,
//		Department: "A",
//	},
//
//	{
//		ID:         "1",
//		Profit:     5,
//		Department: "A",
//	},
//
//	{
//		ID:         "2",
//		Profit:     5,
//		Department: "A",
//	}
//
// return
//
//	{
//		CompanyID:   "1",
//		TotalProfit: 20,
//		Department: []company.DepartmentProfit{
//			{
//				Department:       "A",
//				DepartmentProfit: 15,
//			},
//			{
//				Department:       "B",
//				DepartmentProfit: 5,
//			},
//		},
//	},
func CalculateProfits(profits []CompanyProfit) []CompanyTotalProfit {
	// TODO: add your code here
	return []CompanyTotalProfit{}
}
