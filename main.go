package main

import "findCompaniesEmailForH1bVisa/sendEmailToFoundCompanies"

func main() {
	//first getting the companiesName name from h1b grader website for software developer position
	//companiesName.GetCompaniesName()
	//then find the companiesName website according to their name
	//companiesWebsites.FindWebsites()
	//time to find companies email if it is existed and save them to use it.
	//companiesEmail.FindEmail()
	sendEmailToFoundCompanies.SendEmail()
}
