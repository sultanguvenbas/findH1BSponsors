package googleMapsSearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const APIKey = "yourapiKey"

var count = 0
var f = excelize.NewFile()

func findCompanies(query string) {

	url := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" + query + "&key=" + APIKey
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body := &bytes.Buffer{}

	_, err = io.Copy(body, res.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println("copy")
	}
	var data mapSearch

	err = json.Unmarshal(body.Bytes(), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range data.Results {
		count += 1
		cellA := "A" + strconv.Itoa(count)
		err := f.SetCellValue("Sheet1", cellA, item.Name)
		if err != nil {
			fmt.Println("cella", err.Error())
			return
		}
		fmt.Println(count, item.Name)
	}
	if err := f.SaveAs("employerForMapSearch.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func getCompaniesNameByGoogleSearch() {

	for state, cities := range cityAndState {
		//to be able to search in google replace white spaces by %20 which stands for space
		stateName := strings.Replace(state, " ", "%20", -1)
		for _, city := range cities {
			cityName := strings.Replace(city, " ", "%20", -1)
			findCompanies("IT%20Companies%20in%20" + cityName + "%20" + stateName)
		}

	}
}

// UniqueCompany I don't want to make another google search because it is going to charge me more :D So I took care of them right here
func UniqueCompany() {

	f, err := excelize.OpenFile("employerForMapSearch.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.GetRows("Sheet1")
	// Create a map with string keys and empty struct values
	companySet := make(map[string]struct{})

	// Add elements to the set

	for _, companyName := range rows {
		companySet[companyName[0]] = struct{}{}
	}
	//unique companies
	count := 0
	for company := range companySet {
		count += 1
		cellA := "A" + strconv.Itoa(count)
		err = f.SetCellValue("Sheet1", cellA, company)
		if err != nil {
			fmt.Println("cella", err.Error())
			return
		}
		fmt.Println(company)

	}
	fmt.Println(len(companySet))
	if err := f.SaveAs("employerForMapSearch.xlsx"); err != nil {
		log.Fatal(err)
	}

}
