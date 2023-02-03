package companiesName

// That is the struct for the data returned by Bing.

type Data struct {
	Draw            int `json:"draw"`
	RecordsFiltered int `json:"recordsFiltered"`
	RecordsTotal    int `json:"recordsTotal"`
	Data            []struct {
		Count        int    `json:"count"`
		EmployerName string `json:"employer_name"`
		EmployerUrl  string `json:"employer_url"`
	} `json:"data"`
}
