package googleMapsSearch

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func FindEmailGoogleMap() {
	f, err := excelize.OpenFile("employerForMapSearch.xlsx")

	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.GetRows("Sheet1")
	for index, row := range rows {

		if len(row) == 2 {
			companyUrl := row[1]
			cellC := "C" + strconv.Itoa(index+1)

			htmlBody := getHtmlBody(companyUrl)

			/*re := regexp.MustCompile("mailto:[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}")
			emails := re.FindAllString(htmlBody, -1)*/

			re := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)
			emails := re.FindAllString(htmlBody, -1)

			uniqueEmailList := getUniqueEmails(emails)
			fmt.Println(index+1, uniqueEmailList)
			err = f.SetCellValue("Sheet1", cellC, uniqueEmailList)
			err = f.SaveAs("employerForMapSearch.xlsx")
			if err != nil {
				fmt.Println("saving error", err.Error())
				return
			}
		}

	}

}

func getHtmlBody(companyUrl string) string {
	//skipping SSL certificate verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: time.Second * 25}
	res, err := client.Get(companyUrl)
	if err != nil {
		fmt.Println("error occurred while making the request:", err.Error())
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	var bodyBytes bytes.Buffer
	_, err = io.Copy(&bodyBytes, res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return bodyBytes.String()

}

func getUniqueEmails(emails []string) []string {
	emailMap := make(map[string]bool)
	for _, email := range emails {
		emailLower := strings.ToLower(email)
		emailMap[emailLower] = true
	}

	var uniqueEmails []string
	for email := range emailMap {
		if strings.Contains(email, ".png") || strings.Contains(email, "wixpress.com") || strings.Contains(email, ".jpg") || strings.Contains(email, ".jpeg") || strings.Contains(email, "sentry.io") {
			fmt.Println(email)
		} else {
			uniqueEmails = append(uniqueEmails, email)

		}

	}

	return uniqueEmails
}

// if you want to make sure the e-mail is definitely right you can use that
/*if checkEmailIsValid(email){
uniqueEmails = append(uniqueEmails, email)
}*/
//***Important also you need to sign up and get your api key from isitarealemail.com
/*func checkEmailIsValid(email string) bool {
	apiKey := "passyourapiKey"

	sentUrl := "https://isitarealemail.com/api/email/validate?email=" + url.QueryEscape(email)

	req, _ := http.NewRequest("GET", sentUrl, nil)
	req.Header.Set("Authorization", "bearer "+apiKey)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error %v", err)
		return false
	}
	if res.StatusCode != 200 {
		fmt.Printf("unexpected result, check your api key. %v", res.Status)
		return false
	}

	var myJson RealEmailResponse
	err = json.Unmarshal(body, &myJson)
	if err != nil {
		return false
	}
	if myJson.Status == "invalid" {
		return false
	} else {
		return true
	}
}
*/
