package companiesWebsites

import (
	"fmt"
	"github.com/EdmundMartin/gosearcher"
	"github.com/xuri/excelize/v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func FindWebsites() {
	f, err := excelize.OpenFile("employerH1B.xlsx")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.GetRows("Sheet1")

	for index, row := range rows {

		companyName := row[1]
		companyUrl := bingSearch(companyName)
		cellD := "D" + strconv.Itoa(index+1)
		fmt.Println(companyUrl, index+1)
		err = f.SetCellValue("Sheet1", cellD, companyUrl)
		if err != nil {
			fmt.Println("cella", err.Error())
			return
		}
		err = f.SaveAs("employerH1B.xlsx")
		if err != nil {
			fmt.Println("saving error", err.Error())
			return
		}
	}

}

func findCompanyWebsiteForBingUrl(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("url can't be found")
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(body)
}

func avoidLinkedInUrls(companyName string) string {
	res, err := gosearcher.BingScrape(companyName, "com", nil, 1, 1, 1)
	count := 1
	var companyUrl string
	if err == nil {
		for _, res := range res {
			count = count + 1
			companyUrl = res.ResultURL
			print("check", companyUrl, "\n")
		}
	} else {
		fmt.Println("try", err.Error())
		return ""
	}

	//sometimes the first search can be LinkedIn.We need to make sure it is not linked in because we can't get in linked website without auth
	//making sure it is not empty or LinkedIn website with recursive function

	if count == 10 {
		return companyUrl
	} else if urlStartsWithLinkedIn(companyUrl) {
		return avoidLinkedInUrls(companyName)
	} else if companyUrl == "" {
		return avoidLinkedInUrls(companyName)
	} else {
		return companyUrl
	}

}

func bingSearch(companyName string) string {

	companyUrl := avoidLinkedInUrls(companyName)
	// if url starts with bing that means we need to find the actual website. Good thing is actual website can be found easily with function urlStartsWithBing
	if urlStartsWithBing(companyUrl) {

		htmlBody := findCompanyWebsiteForBingUrl(companyUrl)
		if htmlBody != "" {
			re := regexp.MustCompile(`var u = "(.+?)"`)
			matches := re.FindStringSubmatch(htmlBody)
			if len(matches) > 1 {
				companyUrl = matches[1]
				return companyUrl
			} else {
				fmt.Println("No match found")
			}
		}
		return ""
	}
	return companyUrl

}

// if url starts with bing that means we are not going to get the html body when we try to find email of the company. So we should check it
func urlStartsWithBing(url string) bool {

	if strings.HasPrefix(url, "https://www.bing.com/") {
		// fmt.Println("The URL starts with https://www.bing.com/")
		return true
	} else {
		// fmt.Println("The URL does not start with https://www.bing.com/")
		return false
	}

}

func urlStartsWithLinkedIn(url string) bool {
	if strings.HasPrefix(url, "https://in.linkedin.com/") || strings.HasPrefix(url, "https://www.linkedin.com") {
		// fmt.Println("The URL starts with https://www.bing.com/")
		return true
	} else {
		// fmt.Println("The URL does not start with https://www.bing.com/")
		return false
	}
}
