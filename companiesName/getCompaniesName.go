package companiesName

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

// GetCompaniesName is getting the companiesName name from https://h1bgrader.com/ website for software developer position
func GetCompaniesName() {

	url := "https://h1bgrader.com/api/listby"
	method := "POST"

	payload := strings.NewReader(`draw=5&columns%5B0%5D%5Bdata%5D=employer_name&columns%5B0%5D%5Bname%5D=&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=false&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&start=0&length=4225&search%5Bvalue%5D=&search%5Bregex%5D=false&job=Software+Developer`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "h1bgrader.com")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("accept-language", "en-US,en;q=0.9,tr;q=0.8")
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("cookie", "_gid=GA1.2.677227712.1675026469; _clck=qz4lq8|1|f8o|0; session=eyJhbGciOiJSUzI1NiIsImtpZCI6InRCME0yQSJ9.eyJpc3MiOiJodHRwczovL3Nlc3Npb24uZmlyZWJhc2UuZ29vZ2xlLmNvbS9oMWJncmFkZXIxIiwibmFtZSI6IkFsaSBBa2dvbCIsImF1ZCI6ImgxYmdyYWRlcjEiLCJhdXRoX3RpbWUiOjE2NzUwMjY5NDYsInVzZXJfaWQiOiJkZVh0VFZScXF2Z0ZsR2lyRlhTakxQWkFKT0YyIiwic3ViIjoiZGVYdFRWUnFxdmdGbEdpckZYU2pMUFpBSk9GMiIsImlhdCI6MTY3NTAyNjk0NywiZXhwIjoxNjc1NDU4OTQ3LCJlbWFpbCI6ImFrZ29sOTdfQGhvdG1haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbImFrZ29sOTdfQGhvdG1haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.dKHeKEk8soqNQS4NARxibNu_eo1hAIFSHfaQ2u9UUpMepI_85YsarRFSIRLYThzsjwx7tukzZgU7D2v1MT01R8dQja5iOGdyQVigm-spXgGiIrNBVCJ6MxXcV14ZEbmkg7k_y9oZClECJ-sBu88Krjuw858SHom0GqBGF7FruSDzZyIZ3M6x9Vz4PwrMkhv6tx3f_jpMufX7DYcrS0ySbaJLmiX7kBxYniCfJzFNkg02CdT82qLSYBDquYTH-M0_5VxlI7Ho1H7YFqOAISM12Fm2KBH2f-CK8T4p4mLWK_jFnXMv6P6Kxn4TNwA_umtG2FigaFGzvoCH2pFmp-dQWA; _ga_FQCGQZP8BF=GS1.1.1675026469.5.1.1675028955.0.0.0; _ga=GA1.2.2137793615.1670283876; mp_79dd0f285a9c72f4bcdab6cf2346014e_mixpanel=%7B%22distinct_id%22%3A%20%22akgol97_%40hotmail.com%22%2C%22%24device_id%22%3A%20%22184e4ad5604d9-0b08428fd88726-7d5d5476-1fa400-184e4ad56051019%22%2C%22%24search_engine%22%3A%20%22google%22%2C%22%24initial_referrer%22%3A%20%22https%3A%2F%2Fwww.google.com%2F%22%2C%22%24initial_referring_domain%22%3A%20%22www.google.com%22%2C%22%24user_id%22%3A%20%22akgol97_%40hotmail.com%22%7D; _clsk=7v9x2t|1675028956106|18|1|h.clarity.ms/collect")
	req.Header.Add("origin", "https://h1bgrader.com")
	req.Header.Add("referer", "https://h1bgrader.com/job-titles/software-developer-5g2rng360l")
	req.Header.Add("sec-ch-ua", "\"Not_A Brand\";v=\"99\", \"Microsoft Edge\";v=\"109\", \"Chromium\";v=\"109\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.70")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("something went wrong")
		}
	}(res.Body)
	body := &bytes.Buffer{}
	_, err = io.Copy(body, res.Body)
	if err != nil {
		fmt.Println("copy")
	}

	var data Data

	err = json.Unmarshal(body.Bytes(), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	f := excelize.NewFile()

	for i, item := range data.Data {
		cellA := "A" + strconv.Itoa(i+1)
		cellB := "B" + strconv.Itoa(i+1)
		cellC := "C" + strconv.Itoa(i+1)
		err := f.SetCellValue("Sheet1", cellA, i+1)
		if err != nil {
			fmt.Println("cella", err.Error())
			return
		}
		err = f.SetCellValue("Sheet1", cellB, item.EmployerName)
		if err != nil {
			fmt.Println("cellb")

			return
		}
		err = f.SetCellValue("Sheet1", cellC, item.EmployerUrl)
		if err != nil {
			fmt.Println("cellc")

			return
		}
	}
	if err := f.SaveAs("employerH1B.xlsx"); err != nil {
		log.Fatal(err)
	}

}
