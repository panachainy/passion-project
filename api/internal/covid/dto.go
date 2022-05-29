package covid

type CovidClientResponse []struct {
	CovidSerializer
}

type CovidSerializer struct {
	TxnDate                string `json:"txn_date"`
	NewCase                int    `json:"new_case"`
	TotalCase              int    `json:"total_case"`
	NewCaseExcludeabroad   int    `json:"new_case_excludeabroad"`
	TotalCaseExcludeabroad int    `json:"total_case_excludeabroad"`
	NewDeath               int    `json:"new_death"`
	TotalDeath             int    `json:"total_death"`
	NewRecovered           int    `json:"new_recovered"`
	TotalRecovered         int    `json:"total_recovered"`
	UpdateDate             string `json:"update_date"`
}
