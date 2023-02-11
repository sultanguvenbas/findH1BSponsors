package main

import "findCompaniesEmailForH1bVisa/sendEmailToFoundCompanies"

func main() {
	///** h1b grader .com version
	//first getting the companiesName name from h1b grader website for software developer position
	//companiesName.GetCompaniesName()
	//then find the companiesName website according to their name
	//companiesWebsites.FindWebsites()
	//time to find companies email if it is existed and save them to use it.
	//companiesEmail.FindEmail()

	// **** Google Map Search Version
	//get all the companies information
	//googleMapsSearch.getCompaniesNameByGoogleSearch()

	//find all the companies website by their name
	//googleMapsSearch.FindWebsites()

	//find companies email from their website
	//googleMapsSearch.FindEmailGoogleMap()

	//send emails
	sendEmailToFoundCompanies.SendEmail()

	//sendEmailToFoundCompanies.EmailTemplate("Sultan", "sultanguvenbas@gmail.com")

}
